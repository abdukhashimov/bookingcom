// Code generated by sqlc. DO NOT EDIT.

package sqlc

import (
	"time"

	"abdukhashimov/mybron.uz/storage/custom"
)

type BookObject struct {
	ID              string      `json:"id"`
	Category        string      `json:"category"`
	Title           string      `json:"title"`
	Location        string      `json:"location"`
	Long            float64     `json:"long"`
	Lat             float64     `json:"lat"`
	About           string      `json:"about"`
	Discount        *int32      `json:"discount"`
	DiscountExpires custom.Time `json:"discount_expires"`
	Status          *int32      `json:"status"`
	OpensAt         string      `json:"opens_at"`
	ClosesAt        string      `json:"closes_at"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

type Category struct {
	ID          string    `json:"id"`
	ParentID    *string   `json:"parent_id"`
	Name        string    `json:"name"`
	Image       *string   `json:"image"`
	Active      *bool     `json:"active"`
	Slug        string    `json:"slug"`
	Lang        string    `json:"lang"`
	Information *string   `json:"information"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Faq struct {
	ID        string    `json:"id"`
	Question  *string   `json:"question"`
	Answer    *string   `json:"answer"`
	Slug      string    `json:"slug"`
	Lang      string    `json:"lang"`
	Active    *bool     `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Status struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
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
