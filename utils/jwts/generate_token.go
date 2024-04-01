package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
	"gvd_server/global"
	"time"
)

// GenToken 生成token
func GenToken(user JwyPayLoad) (string, error) {
	claims := CustomClaims{
		JwyPayLoad: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Duration(global.Config.Jwt.Expires) * time.Hour)), // 过期时间
			Issuer:    global.Config.Jwt.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.Config.Jwt.Secret))
}
