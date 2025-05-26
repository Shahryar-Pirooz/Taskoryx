package handlers

import "time"

type LoginReq struct {
	Email    string `query:"email"`
	Password string `query:"password"`
}

type RegisterReq struct {
	Name     string `query:"name"`
	Email    string `query:"email"`
	Password string `query:"password"`
}

type RefreshTokenReq struct {
	AccessToken string `query:"access_token"`
	UserID      string `query:"user_id"`
}

type UserReq struct {
	ID        string    `query:"id"`
	Name      string    `query:"name"`
	Email     string    `query:"email"`
	Password  string    `query:"password"`
	Role      int8      `query:"role"`
	CreatedAt time.Time `query:"created_at"`
	UpdatedAt time.Time `query:"updated_at"`
}
