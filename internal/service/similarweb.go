package service

import (
	"context"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/matisiekpl/similarcli/internal/repository"
	"github.com/olekukonko/tablewriter"
)

type SimilarWebService interface {
	Print(domain string) error
}

type similarWebService struct {
	similarWebRepository repository.SimilarWebRepository
}

func NewSimilarWebService() SimilarWebService {
	return &similarWebService{
		similarWebRepository: repository.NewSimilarWebRepository(),
	}
}

func formatDate(dateStr string) string {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return dateStr
	}
	return t.Format("January 2006")
}

func formatNumber(num int) string {
	switch {
	case num >= 1000000:
		return fmt.Sprintf("%.1fM", float64(num)/1000000)
	case num >= 1000:
		return fmt.Sprintf("%.1fK", float64(num)/1000)
	default:
		return fmt.Sprintf("%d", num)
	}
}

func formatPercentage(value float64) string {
	return fmt.Sprintf("%.1f%%", value*100)
}

func (s *similarWebService) Print(domain string) error {
	website, err := s.similarWebRepository.Load(context.Background(), domain)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if website.Title == "" {
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
	basicTable.Render()

	visitsTable := tablewriter.NewWriter(os.Stdout)
	fmt.Println()

	var months []string
	for month := range website.EstimatedMonthlyVisits {
		months = append(months, month)
	}
	sort.Strings(months)

	headers := []string{"Date"}
	visits := []string{"Visits"}

	for _, month := range months {
		count := website.EstimatedMonthlyVisits[month]
		headers = append(headers, formatDate(month))
		visits = append(visits, formatNumber(count))
	}

	visitsTable.SetHeader(headers)
	visitsTable.SetAutoWrapText(false)
	visitsTable.SetAutoFormatHeaders(true)
	visitsTable.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	visitsTable.SetAlignment(tablewriter.ALIGN_LEFT)
	visitsTable.SetCenterSeparator("")
	visitsTable.SetColumnSeparator("")
	visitsTable.SetRowSeparator("")
	visitsTable.SetHeaderLine(false)
	visitsTable.SetBorder(false)
	visitsTable.SetTablePadding("\t")
	visitsTable.SetNoWhiteSpace(true)

	visitsTable.Append(visits)
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
		{"Search", website.TrafficSources.Search},
		{"Social", website.TrafficSources.Social},
		{"Referrals", website.TrafficSources.Referrals},
		{"Mail", website.TrafficSources.Mail},
		{"Paid Referrals", website.TrafficSources.PaidReferrals},
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
