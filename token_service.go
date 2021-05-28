package AutoMigration

import (
	pb "github.com/Buzzology/shippy-service-user/proto/user"
	"github.com/dgrijalva/jwt-go"
)

var (
	// This is the salt - recommendation is a randomly generated md5 hash
	key = []byte("mySuperSecretKey")
)

type CustomClaims struct {
	User *pb.user
	jwt.StandardClaims
}

type TokenService struct {
	repo Repository
}

// Decode a token string into a token object
func (srv *TokenService) Decode(token string) (*CustomClaims, error) {

	// Parse token
	tokenType, err := jwt.ParseWithClaims(string(key), &CustomClaims{}, func(token *jwt.Token) (interface{}, error)
	{
		return key, nil
	})

	// Checks if the claims value is of type CustomClaims
	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		return claims, nil
	}
	else {
		return nil, err
	}
}


// Encode a claim into a JWT
func (srv *TokenService) Encode(user *pb.User) (string, error) {

	// Create the claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims {
			ExpiresAt: 15000,
			Issuer: "shippy.service.user"
		}
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token and return
	return token.SignedString(key)
}