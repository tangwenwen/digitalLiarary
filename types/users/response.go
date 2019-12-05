package usersType



type LoginResp struct {
	Token string `json:"token"`
	UserType int `json:"user_type"`
}

type Users struct{
	ID int `xorm:"id"`
	Username string `xorm:"username"`
	Password string `xorm:"password"`
	UserType int `xorm:"user_type"`
}

type TokenValue struct {
	UserName string
	UserType string
}