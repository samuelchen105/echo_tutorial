package api

type recieveUser struct {
	Name  string `json:"name" form:"name" validate:"required"`
	Email string `json:"email" form:"email" validate:"required,email"`
}

type user struct {
	Name    string `json:"name" xml:"name"`
	Email   string `json:"email" xml:"email"`
	IsAdmin bool   `json:"is_admin" xml:"is_admin"`
}
