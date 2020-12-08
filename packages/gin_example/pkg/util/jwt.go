package util

import (
	"gin-example/models"
	"gin-example/pkg/logging"
	"gin-example/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"github.com/tidwall/gjson"
	"strings"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id       int    `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(username, password string, id int) (string, error) {
	// 这个地方可以考虑通过密码动态授权
	var jwtSecret = []byte(setting.AppSetting.JwtSecret + password)

	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	payload := strings.Split(token, ".")

	// 获取token 的中间段信息
	bytes, e := jwt.DecodeSegment(payload[1])

	if e != nil {
		println(e.Error())
	}
	content := ""
	for i := 0; i < len(bytes); i++ {
		content += string(bytes[i])
	}

	id := gjson.Get(content, "id").Int()

	logging.Info("id", id)
	user := models.GetAuthById(id)
	logging.Info("user", user)

	// 通过密码动态授权
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(setting.AppSetting.JwtSecret + user.Password), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
