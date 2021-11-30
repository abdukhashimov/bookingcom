// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Category struct {
	ID          string  `json:"id"`
	ParentID    *string `json:"parent_id"`
	Name        string  `json:"name"`
	Image       string  `json:"image"`
	Active      bool    `json:"active"`
	Slug        string  `json:"slug"`
	Lang        string  `json:"lang"`
	Information *string `json:"information"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type CreateCategory struct {
	Name        string  `json:"name"`
	ParentID    *string `json:"parent_id"`
	Image       string  `json:"image"`
	Active      bool    `json:"active"`
	Information *string `json:"information"`
}

type CreateFaq struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Active   bool   `json:"active"`
}

type Faq struct {
	ID        string  `json:"id"`
	Question  string  `json:"question"`
	Answer    string  `json:"answer"`
	Active    bool    `json:"active"`
	Slug      string  `json:"slug"`
	Lang      string  `json:"lang"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type GetAllCategory struct {
	Categories []*Category `json:"categories"`
	Count      *int        `json:"count"`
}

type GetAllResp struct {
	Faqs  []*Faq `json:"faqs"`
	Count *int   `json:"count"`
}

type LoginParams struct {
	PhoneNumber   string `json:"phone_number"`
	OtpCode       string `json:"otp_code"`
	PasscodeToken string `json:"passcode_token"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type NewUser struct {
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	PhoneNumber string   `json:"phone_number"`
	Long        *float64 `json:"long"`
	Lat         *float64 `json:"lat"`
}

type UpdateCategory struct {
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Lang        string  `json:"lang"`
	ParentID    *string `json:"parent_id"`
	Image       string  `json:"image"`
	Active      bool    `json:"active"`
	Information *string `json:"information"`
}

type UpdateFaq struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Active   bool   `json:"active"`
	Slug     string `json:"slug"`
	Lang     string `json:"lang"`
}

type UpdateResponse struct {
	ID string `json:"id"`
}

type UpdateUser struct {
	FirstName *string  `json:"first_name"`
	LastName  *string  `json:"last_name"`
	Long      *float64 `json:"long"`
	Lat       *float64 `json:"lat"`
}

type User struct {
	ID          string   `json:"id"`
	CreatedAt   *string  `json:"created_at"`
	UpdatedAt   *string  `json:"updated_at"`
	FirstName   *string  `json:"first_name"`
	LastName    *string  `json:"last_name"`
	PhoneNumber string   `json:"phone_number"`
	Long        *float64 `json:"long"`
	Lat         *float64 `json:"lat"`
	UserType    *int     `json:"user_type"`
}
