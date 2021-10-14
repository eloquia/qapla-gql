package model

import "time"

type User struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	GoesBy      string `json:"goes_by"`
	MiddleName  string `json:"middle_name"`
	Email       string `json:"email"`
	Password    string
	Gender      string    `json:"gender"`
	Ethnicity   string    `json:"ethnicity"`
	Position    string    `json:"position"`
	Institution string    `json:"institution"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUser struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	GoesBy      string    `json:"goes_by"`
	MiddleName  string    `json:"middle_name"`
	Email       string    `json:"email"`
	Gender      string    `json:"gender"`
	Ethnicity   string    `json:"ethnicity"`
	Position    string    `json:"position"`
	Institution string    `json:"institution"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserDetails struct {
	ID               string     `json:"id"`
	FirstName        string     `json:"firstName"`
	LastName         string     `json:"lastName"`
	Email            string     `json:"email"`
	AssignedProjects []*Project `json:"assignedProjects"`
}
