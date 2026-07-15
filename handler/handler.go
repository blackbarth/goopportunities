package handler

import "time"



func (r *CreateOpeningRequest) ToModel() *CreateOpeningResponse {
	return &CreateOpeningResponse{
		Role:     r.Role,
		Company:  r.Company,
		Location: r.Location,
		Remote:   *r.Remote,
		Link:     r.Link,
		Salary:   r.Salary,
	}
}

// CreateOpening

func (r *CreateOpeningResponse) ToRequest() *CreateOpeningRequest {
	return &CreateOpeningRequest{
		Role:     r.Role,
		Company:  r.Company,
		Location: r.Location,
		Remote:   &r.Remote,
		Link:     r.Link,
		Salary:   r.Salary,
	}
}



type CreateOpeningResponse struct {
	Id uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Role string `json:"role"`
	Company string `json:"company"`
	Location string `json:"location"`
	Remote bool `json:"remote"`
	Link string `json:"link"`
	Salary int64 `json:"salary"`
}