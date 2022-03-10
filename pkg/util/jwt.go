package util

import (
	"Bus/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var jwtSecret = []byte(setting.JWTSecret)

type Claims struct {
	ID   primitive.ObjectID `bson:"_id" json:"id"`
	Role int                `bson:"role" json:"role"`
	jwt.StandardClaims
}

func GenerateToken(id primitive.ObjectID, role int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)

	claims := Claims{
		id,
		role,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Japs0n",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
