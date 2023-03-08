package types

type UserLoginRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type UserRegisterRequest struct {
	UserEmail string `json:"useremail"`
	PassWord  string `json:"Password"`
	RegType   int8   `json:"regtype"`
	UserTel   string `json:"usertel"`
	Captcha   string `json:"captcha"`
}

type UserVipRequest struct {
	UserId uint `json:"userid"`
}

type UserVipResponse struct {
	UserNick string `json:"usernick"`
	UserFace string `json:"userface"`
	IsVip    uint   `json:"isvip"`
	BuyTime  string `json:"buytime"`
	BuyMoney string `json:"buymoney"`
}

type WechatCode struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

type TestRequest struct {
	Text string `form:"text" json:"text"`
}
