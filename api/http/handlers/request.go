package handlers

type CreateUserReq struct {
	Name     string `json:"name" xml:"name" form:"name"`
	Email    string `json:"email" xml:"email" form:"email"`
	Password string `json:"password" xml:"password" form:"password"`
	Role     uint8  `json:"role" xml:"role" form:"role"`
}
