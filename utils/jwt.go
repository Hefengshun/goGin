package utils

import (
	"ginDemo/models/system"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 定义秘钥
var jwtKey = []byte("123456")
var NoVerify = []interface{}{"/user/login"}

type Claims struct {
	UserId   uint
	UserName string
	jwt.StandardClaims
}

// 登录成功之后发放token
func ReleaseToken(user system.SysUser) (string, error) {
	expirationTime := time.Now().Add(3 * time.Minute) //token的有效期是3分钟
	claims := &Claims{
		UserId:   user.ID,
		UserName: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //token的有效期
			IssuedAt:  time.Now().Unix(),     //token发放的时间
			Issuer:    "chengqiang",          //作者
			Subject:   "user token",          //主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey) //根据前面自定义的jwt秘钥生成token

	if err != nil {
		//返回生成的错误
		return "", err
	}
	//返回成功生成的字符换
	return tokenString, nil
}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjYsImV4cCI6MTYzYW5nIiwic3ViIjoidXNlciB0b2tlbiJ9.nt8K7vxrAT4XXzh0RbtFveQCyt7J4r1XZnVgDNSVjkQ
//token由三部分组成
//1.加密协议、2.荷载（程序信息Claims）、前面两部分+自定义密匙组成的一个哈希值 3.密钥
//使用base64解密保存的信息(分三段进行解密) ：  echo eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9 | base64 -D

// 解析从前端获取到的token值
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
