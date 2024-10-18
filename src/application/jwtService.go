package application

import (
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewJwtService(privateKeyPath, publicKeyPath string) *JWTService {
	jwtService := &JWTService{}
	jwtService.loadPrivateKey(privateKeyPath)
	jwtService.loadPublicKey(publicKeyPath)
	return jwtService
}

func (jwtService *JWTService) loadPrivateKey(privateKeyPath string) {
	privKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		panic(err)
	}
	jwtService.privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privKeyBytes)
	if err != nil {
		panic(err)
	}
}

func (jwtService *JWTService) loadPublicKey(publicKeyPath string) {
	publicKeyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		panic(err)
	}
	jwtService.publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		panic(err)
	}
}

func (jwtService *JWTService) GenerateJWT(username string) string {
	claims := jwt.MapClaims{
		"iss": "test",
		"sub": username,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(jwtService.privateKey)
	if err != nil {
		panic(err)
	}
	return tokenString
}

func (jwtService *JWTService) VerifyToken(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		//     return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		// }
		return jwtService.publicKey, nil
	})
	if err != nil {
		panic(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	}
	panic(fmt.Errorf("invalid token"))
}
