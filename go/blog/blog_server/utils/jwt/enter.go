package jwt

import (
  "blog_server/global"
  "errors"
  "github.com/gin-gonic/gin"
  "strings"
  "time"

  "github.com/dgrijalva/jwt-go"
)

type Claims struct {
  UserID   uint   `json:"userID"`
  UserName string `json:"username"`
  Role     int8   `json:"role"`
}

type MyClaims struct {
  Claims
  jwt.StandardClaims
}

// get token
func GetToken(claims Claims) (string, error) {
  cla := MyClaims{
    Claims: claims,
    StandardClaims: jwt.StandardClaims{
      ExpiresAt: time.Now().Add(time.Duration(global.Config.Jwt.Expire) * time.Hour).Unix(), // 过期时间
      Issuer:    global.Config.Jwt.Issuer,                                                     // 签发人
    },
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
  return token.SignedString([]byte(global.Config.Jwt.Secret)) // 进行签名生成对应的token
}

// parse token
func ParseToken(tokenString string) (*MyClaims, error) {
  if tokenString == "" {
    return nil, errors.New("请登录")
  }
  token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
    return []byte(global.Config.Jwt.Secret), nil
  })
  if err != nil {
    if strings.Contains(err.Error(), "token is expired") {
      return nil, errors.New("token过期")
    }
    if strings.Contains(err.Error(), "signature is invalid") {
      return nil, errors.New("token无效")
    }
    if strings.Contains(err.Error(), "token contains an invalid") {
      return nil, errors.New("token非法")
    }
    return nil, err
  }
  if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
    return claims, nil
  }
  return nil, errors.New("invalid token")
}

func ParseTokenByGin(c *gin.Context) (*MyClaims, error) {
  token := c.GetHeader("token")
  if token == "" {
    token = c.Query("token")
  }
  return ParseToken(token)
} 
