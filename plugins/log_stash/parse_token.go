package log_stash

import "github.com/dgrijalva/jwt-go/v4"

type JwyPayLoad struct {
	NickName string `json:"nickName"`
	RoleID   uint   `json:"roleID"`
	UserID   uint   `json:"userID"`
	UserName string `json:"userName"`
}

type CustomClaims struct {
	JwyPayLoad
	jwt.StandardClaims
}

func parseToken(token string) (jwtPayload *JwyPayLoad) {
	Token, _ := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	if Token.Claims == nil {
		return
	}
	claims, ok := Token.Claims.(*CustomClaims)
	if !ok {
		return nil
	}
	return &claims.JwyPayLoad
}
