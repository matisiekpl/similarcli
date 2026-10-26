package dto

import "time"

type RdapDomain struct {
	Events []RdapEvent `json:"events"`
}

type RdapEvent struct {
	EventAction string    `json:"eventAction"`
	EventDate   time.Time `json:"eventDate"`
}
