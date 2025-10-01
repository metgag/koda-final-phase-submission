package models

type RegisterBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Uname    string `json:"username" binding:"required"`
}

type LoginBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginScan struct {
	UID     int
	HashPwd string
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
