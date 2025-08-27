package models

type Incident struct {
	ID            uint32    `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	IncidentType  string    `json:"incident_type"`
	Location      string    `json:"location"`
	Image         string    `json:"image"`
}
