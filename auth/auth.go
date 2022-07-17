package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"shape-api/model"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"
	"github.com/twinj/uuid"
)

//AuthController ...
type Auth struct{}

//CreateToken ...
func (m Auth) CreateToken(userName string) (*model.TokenDetails, error) {

	td := &model.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.NewV4().String()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_name"] = userName
	atClaims["exp"] = td.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_name"] = userName
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

//TokenValid ...
func (m Auth) TokenValid(c *gin.Context) {

	tokenAuth, err := m.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please login first"})
		return
	}

	if tokenAuth.AccessUUID == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Please login first"})
		return
	}
	return
}

//ExtractTokenMetadata ...
func (m Auth) ExtractTokenMetadata(r *http.Request) (*model.AccessDetails, error) {
	token, err := m.VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userName, ok := claims["user_name"].(string)
		if !ok {
			return nil, err
		}
		return &model.AccessDetails{
			AccessUUID: accessUUID,
			UserName:   userName,
		}, nil
	}
	return nil, err
}

//VerifyToken ...
func (m Auth) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := m.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//ExtractToken ...
func (m Auth) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
