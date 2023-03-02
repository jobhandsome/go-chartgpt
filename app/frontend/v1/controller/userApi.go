package controller

import (
	"fmt"
	"github.com/fastwego/offiaccount/apis/oauth"
	"github.com/gin-gonic/gin"
	"go-chatgpt/app/frontend/models"
	"go-chatgpt/app/frontend/types"
	"go-chatgpt/app/frontend/v1/config"
	"go-chatgpt/pkg/orm"
	"go-chatgpt/pkg/utils"
	"time"
)

func UserLogin(ctx *gin.Context) {
	// 绑定参数
	var params types.UserLoginRequest
	if err := ctx.ShouldBind(&params); err != nil {
		utils.ServerError(ctx, err.Error())
		return
	}

	fmt.Println(params)
	return
}

func UserRegister(ctx *gin.Context) {
	var params types.UserRegisterRequest
	if err := ctx.ShouldBind(&params); err != nil {
		utils.ServerError(ctx, err.Error())
		return
	}

	var user models.Users

	var findCount int64

	orm.Orm.Table(user.TableName()).Where("useremail = ?", params.UserEmail).Count(&findCount)

	if findCount > 0 {
		utils.Fail(ctx, "邮箱已注册，请前往登录...")
		return
	}

	username := "openai_" + utils.RandString(20)

	fmt.Println(username)

	switch params.RegType {
	case 0:
		if ok := utils.VerifyEmail(params.UserEmail); !ok {
			utils.Fail(ctx, "邮箱格式错误")
			return
		}
		break
	case 1:
		if ok := utils.VerifyMobile(params.UserTel); !ok {
			utils.Fail(ctx, "电话格式错误")
			return
		}

		if len(params.Captcha) == 0 {
			utils.Fail(ctx, "请输入验证码")
			return
		}
		break
	}

	fmt.Println(params)

	result := orm.Orm.Table(user.TableName()).Create(&models.Users{
		UserName:   username,
		PassWord:   utils.EncryptPassword(params.PassWord),
		NickName:   username,
		UserSex:    0,
		UserEmail:  params.UserEmail,
		UserTel:    params.UserTel,
		CreateTime: time.Now().Format(time.DateTime),
	})

	if result.Error != nil {
		utils.ServerError(ctx, result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		utils.ServerError(ctx, "注册失败")
		return
	}

	utils.Success(ctx, "注册成功，请前往登录...", nil)
	return
}

func WechatCode(ctx *gin.Context) {
	var params types.WechatCode
	if err := ctx.ShouldBind(&params); err != nil {
		utils.ServerError(ctx, err.Error())
		return
	}
	OauthAccessToken, err := oauth.GetAccessToken(config.Config.Wechat.Appid, config.Config.Wechat.Secret, params.Code)
	if err != nil {
		utils.ServerError(ctx, err.Error())
		return
	}

	utils.Success(ctx, "ok", gin.H{
		"AccessToken":  OauthAccessToken.AccessToken,
		"ExpiresIn":    OauthAccessToken.ExpiresIn,
		"RefreshToken": OauthAccessToken.RefreshToken,
		"Openid":       OauthAccessToken.Openid,
		"Scope":        OauthAccessToken.Scope,
	})
	return
}
