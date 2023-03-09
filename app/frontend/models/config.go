package models

type OpenaiConfig struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	ApiKey     string `json:"username" gorm:"size:255;index"`
	IsUse      uint   `json:"isuse" gorm:"size:1;index;default:0"`
	CreateTime string `json:"createtime" gorm:"size:255;"`
	UpdateTime string `json:"updatetime" gorm:"size:255;"`
}

func (*OpenaiConfig) TableName() string {
	return "openai_configs"
}
