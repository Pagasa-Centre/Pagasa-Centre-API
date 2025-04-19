package dto

type GetAllOutreachesResponse struct {
	Message    string      `json:"message"`
	Outreaches []*Outreach `json:"outreaches,omitempty"`
}

type Outreach struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	AddressLine1 string    `json:"address_line_1"`
	AddressLine2 string    `json:"address_line_2"`
	Postcode     string    `json:"postcode"`
	City         string    `json:"city"`
	Country      string    `json:"country"`
	VenueName    string    `json:"venue_name"`
	Region       string    `json:"region"`
	ThumbnailURL string    `json:"thumbnail_url"`
	Services     []Service `json:"services,omitempty"`
}

type Service struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Day       string `json:"day"`
}
