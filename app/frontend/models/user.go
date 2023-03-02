package models

type Users struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserName    string `json:"username" gorm:"size:255;index"`
	PassWord    string `json:"password" gorm:"size:255;index"`
	NickName    string `json:"nickname" gorm:"size:255;"`
	UserFace    string `json:"userface" gorm:"size:255;"`
	UserSex     uint   `json:"usersex" gorm:"size:1;"`
	UserEmail   string `json:"useremail" gorm:"size:255;"`
	UserTel     string `json:"usertel" gorm:"size:255;"`
	OpenId      string `json:"openid" gorm:"size:255;index"`
	IsVip       uint   `json:"isvip" gorm:"size:1;index"`
	UseQuantity uint64 `json:"usequantity" gorm:"size:10;index"`
	BuyTime     string `json:"buytime" gorm:"size:255;"`
	VipTime     string `json:"viptime" gorm:"size:255;"`
	CreateTime  string `json:"createtime" gorm:"size:255;"`
	UpdateTime  string `json:"updatetime" gorm:"size:255;"`
}

func (*Users) TableName() string {
	return "users"
}

type UserVipLogs struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	UserId     uint    `json:"userid" gorm:"size:255;index"`
	BuyTime    string  `json:"buytime" gorm:"size:255;"`
	LastTime   string  `json:"lasttime" gorm:"size:255;"`
	BuyMoney   float64 `json:"buymoney" gorm:"size:255;"`
	CreateTime string  `json:"createtime" gorm:"size:255;"`
}
