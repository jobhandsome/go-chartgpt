package utils

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    gin.H  `json:"data"`
}

func Success(ctx *gin.Context, message string, data gin.H) {
	ctx.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
	return
}

func Fail(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, Response{
		Code:    http.StatusBadRequest,
		Message: message,
		Data:    nil,
	})
	return
}

func ServerError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, Response{
		Code:    http.StatusInternalServerError,
		Message: message,
		Data:    nil,
	})
	return
}

func RandString(len int) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, len)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(62)]
	}
	return string(b)
}

func VerifyEmail(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func VerifyMobile(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func EncryptPassword(pwd string) string {
	password, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MaxCost)
	if err != nil {
		return ""
	}
	return string(password)
}

func DecryptPassword(EnPassword string, loginPassWord string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(EnPassword), []byte(loginPassWord))
	if err == nil {
		return true
	}
	return false
}
