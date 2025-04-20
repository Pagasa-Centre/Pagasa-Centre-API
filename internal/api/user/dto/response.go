package dto

type UpdateUserDetailsResponse struct {
	Message string       `json:"message"`
	User    *UserDetails `json:"user,omitempty"`
}

type UserDetails struct {
	FirstName    string   `json:"first_name,omitempty"`
	LastName     string   `json:"last_name,omitempty"`
	Email        string   `json:"email,omitempty" validate:"omitempty,email"`
	Birthday     string   `json:"birthday,omitempty"`
	PhoneNumber  string   `json:"phone_number,omitempty"`
	CellLeaderID *string  `json:"cell_leader_id,omitempty"`
	OutreachID   string   `json:"outreach_id,omitempty"`
	Roles        []string `json:"roles,omitempty"`
}

type RegisterResponse struct {
	Token   *string      `json:"token,omitempty"`
	User    *UserDetails `json:"user,omitempty"`
	Message string       `json:"message"`
}

type LoginResponse struct {
	Message string       `json:"message"`
	Token   *string      `json:"token,omitempty"`
	User    *UserDetails `json:"user,omitempty"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}
