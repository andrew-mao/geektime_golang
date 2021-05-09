package middleware

import (
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	SigningKey []byte
}

//JWT载荷，添加部分自定义内容
//访问者身份、IP、地址、过期时间
type ClaimsMy struct {
	UserIdentity bool   `json:"user_identity"` //访问者身份
	IP           string `json:"user_id"`       //访问者IP
	Address      string `json:"user_address"`  //访问者地址
	//ExpireTime   time.Time `json:"expire_time"`   //过期时间
	jwt.StandardClaims //实现标准载荷
}

//生成token
func (j *JWT) CreateToken(clamis ClaimsMy) (string, error) {
	//token : jwt.NewWithClaims(jwt.SigningMethodHS384,clamis)
	return jwt.NewWithClaims(jwt.SigningMethodHS384, clamis).SignedString(j.SigningKey)
}

//解析token
//func (j *JWT)ParseToken(tokenString string) (*ClaimsMy,error) {
//	token, err :=jwt.ParseWithClaims(tokenString, &ClaimsMy{}, func(token *jwt.Token) (interface{}, error) {
//		return j.SigningKey,nil
//	})
//	if err != nil {
//		if ve,ok :=err.
//	}
//}

//更新token
