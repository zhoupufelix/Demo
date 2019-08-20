package library

import (
	"math/rand"
	"encoding/hex"
	"crypto/md5"
	"github.com/dgrijalva/jwt-go"
	"Demo/config"
)

//jwt密钥
var jwtSecret = []byte(config.JWTSECRET)

type Claims struct{
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}


func GenerateUUID()chan int{
	ch := make(chan int,10)
	go func() {
		ch <-rand.Int()
	}()
	return ch
}

func MakeMD5(str string)string{
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

/**
 * 验签
 */
func ParseToken(token string)(*Claims, error){
	tokenClaims,err := jwt.ParseWithClaims(token,&Claims{},func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil})

	if tokenClaims != nil {
		if claims,ok := tokenClaims.Claims.(*Claims);ok && tokenClaims.Valid {
			return claims,nil
		}
	}
	return nil,err
}