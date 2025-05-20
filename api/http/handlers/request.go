package handlers

type LoginReq struct {
	Email    string `query:"email"`
	Password string `query:"password"`
}

type RefreshTokenReq struct {
	AccessToken string `query:"access_token"`
	UserID      string `query:"user_id"`
}
