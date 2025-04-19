package dto

type MinistryApplicationResponse struct {
	Message string `json:"message"`
}

type GetAllMinistriesResponse struct {
	Ministries []*Ministry `json:"ministries,omitempty"`
	Message    string      `json:"message,"`
}

type Ministry struct {
	ID              string  `json:"id"`
	OutreachID      string  `json:"outreach_id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Day             string  `json:"day"`
	StartTime       *string `json:"start_time,omitempty"`
	EndTime         *string `json:"end_time,omitempty"`
	MeetingLocation string  `json:"meeting_location,omitempty"`
}
