// Code generated by sqlc. DO NOT EDIT.

package sqlc

import (
	"database/sql"
	"time"
)

type Faq struct {
	ID        string       `json:"id"`
	Question  *string      `json:"question"`
	Answer    *string      `json:"answer"`
	Slug      *string      `json:"slug"`
	Lang      *string      `json:"lang"`
	Active    sql.NullBool `json:"active"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type User struct {
	ID          string    `json:"id"`
	FirstName   *string   `json:"first_name"`
	LastName    *string   `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	IsVerified  *bool     `json:"is_verified"`
	Long        *float64  `json:"long"`
	Lat         *float64  `json:"lat"`
	UserType    int32     `json:"user_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserType struct {
	ID   int32   `json:"id"`
	Name *string `json:"name"`
}
