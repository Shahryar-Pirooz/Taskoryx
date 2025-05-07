package handlers

type LoginReq struct {
	Email    string `query:"email"`
	Password string `query:"password"`
}
