package service

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/matisiekpl/similarcli/internal/repository"
	"github.com/olekukonko/tablewriter"
)

type SimilarWebService interface {
	Print(domain string) error
}

type similarWebService struct {
	similarWebRepository repository.SimilarWebRepository
	rdapRepository       repository.RdapRepository
}

func NewSimilarWebService() SimilarWebService {
	return &similarWebService{
		similarWebRepository: repository.NewSimilarWebRepository(),
		rdapRepository:       repository.NewRdapRepository(),
	}
}

func formatDate(dateStr string) string {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return dateStr
	}
	return t.Format("January 2006")
}

func formatNumber(number int) string {
	digits := strconv.Itoa(number)
	for index := len(digits) - 3; index > 0; index -= 3 {
		digits = digits[:index] + "," + digits[index:]
	}
	return digits
}

func formatAge(date time.Time) string {
	days := int(time.Since(date).Hours() / 24)
	value, unit := days, "day"
	switch {
	case days >= 365:
		value, unit = days/365, "year"
	case days >= 30:
		value, unit = days/30, "month"
	}
	if value != 1 {
		unit += "s"
	}
	return fmt.Sprintf("%d %s ago", value, unit)
}

func formatPercentage(value float64) string {
	return fmt.Sprintf("%.1f%%", value*100)
}

func (s *similarWebService) Print(domain string) error {
	website, err := s.similarWebRepository.Load(context.Background(), domain)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if website.SiteName == "" {
		return fmt.Errorf("site %s not found", domain)
	}

	basicTable := tablewriter.NewWriter(os.Stdout)
	basicTable.SetAutoWrapText(false)
	basicTable.SetAutoFormatHeaders(true)
	basicTable.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	basicTable.SetAlignment(tablewriter.ALIGN_LEFT)
	basicTable.SetCenterSeparator("")
	basicTable.SetColumnSeparator("")
	basicTable.SetRowSeparator("")
	basicTable.SetHeaderLine(false)
	basicTable.SetBorder(false)
	basicTable.SetTablePadding("\t")
	basicTable.SetNoWhiteSpace(true)

	basicTable.Append([]string{"Site Name", website.SiteName})
	basicTable.Append([]string{"Title", website.Title})
	basicTable.Append([]string{"Description", website.Description})
	basicTable.Append([]string{"Category", website.Category})
	rdapDomain, err := s.rdapRepository.Load(context.Background(), domain)
	if err == nil {
		for _, event := range rdapDomain.Events {
			if event.EventAction == "registration" {
				basicTable.Append([]string{"Registered", fmt.Sprintf("%s (%s)", event.EventDate.Format("2 January 2006"), formatAge(event.EventDate))})
			}
		}
	}
	basicTable.Render()

	visitsTable := tablewriter.NewWriter(os.Stdout)
	fmt.Println()

	var months []string
	for month := range website.EstimatedMonthlyVisits {
		months = append(months, month)
	}
	sort.Strings(months)

	visitsTable.SetHeader([]string{"Date", "Visits"})
	visitsTable.SetAutoWrapText(false)
	visitsTable.SetAutoFormatHeaders(true)
	visitsTable.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	visitsTable.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_RIGHT})
	visitsTable.SetCenterSeparator("")
	visitsTable.SetColumnSeparator("")
	visitsTable.SetRowSeparator("")
	visitsTable.SetHeaderLine(false)
	visitsTable.SetBorder(false)
	visitsTable.SetTablePadding("\t")
	visitsTable.SetNoWhiteSpace(true)

	for _, month := range months {
		visitsTable.Append([]string{formatDate(month), formatNumber(website.EstimatedMonthlyVisits[month])})
	}
	visitsTable.Render()

	trafficTable := tablewriter.NewWriter(os.Stdout)
	fmt.Println()

	trafficHeaders := []string{"Source"}
	trafficValues := []string{"Percentage"}

	sources := []struct {
		name  string
		value float64
	}{
		{"Direct", website.TrafficSources.Direct},
		{"Organic Search", website.TrafficSources.SearchOrganic},
		{"Paid Search", website.TrafficSources.SearchPaid},
		{"Organic Social", website.TrafficSources.SocialOrganic},
		{"Paid Social", website.TrafficSources.SocialPaid},
		{"Referrals", website.TrafficSources.Referrals},
		{"Mail", website.TrafficSources.Mail},
		{"GenAI", website.TrafficSources.GenAi},
		{"Affiliate", website.TrafficSources.Affiliate},
		{"Display Ads", website.TrafficSources.DisplayAds},
	}

	sort.Slice(sources, func(i, j int) bool {
		return sources[i].value > sources[j].value
	})

	for _, source := range sources {
		trafficHeaders = append(trafficHeaders, source.name)
		trafficValues = append(trafficValues, formatPercentage(source.value))
	}

	trafficTable.SetHeader(trafficHeaders)
	trafficTable.SetAutoWrapText(false)
	trafficTable.SetAutoFormatHeaders(true)
	trafficTable.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	trafficTable.SetAlignment(tablewriter.ALIGN_LEFT)
	trafficTable.SetCenterSeparator("")
	trafficTable.SetColumnSeparator("")
	trafficTable.SetRowSeparator("")
	trafficTable.SetHeaderLine(false)
	trafficTable.SetBorder(false)
	trafficTable.SetTablePadding("\t")
	trafficTable.SetNoWhiteSpace(true)

	trafficTable.Append(trafficValues)
	trafficTable.Render()

	return nil
}
