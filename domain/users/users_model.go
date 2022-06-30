package users

import "time"

type userRegisterReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required,min=8"`
}

type userRegisterResp struct {
	CreatedUserId  string    `json:"created_user_id"`
	Token          string    `json:"token"`
	TokenExpiredAt time.Time `json:"token_expired_at"`
}

type userLoginReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required,min=8"`
}

type userLoginResp struct {
	LoggedInUserId string    `json:"logged_in_user_id"`
	Token          string    `json:"token"`
	TokenExpiredAt time.Time `json:"token_expired_at"`
}

type userCheckAuthResp struct {
	Id        string `json:"_id"`
	Username  string `json:"username"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
