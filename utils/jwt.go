package utils

// import (
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/google/uuid"
// 	"github.com/sunnysingha911/user-service/config"
// )

// // GenerateJWT creates a JWT token string for the given user ID
// func GenerateJWT(userID uuid.UUID) (string, error) {
// 	claims := jwt.MapClaims{
// 		"user_id": userID.String(),
// 		"exp":     time.Now().Add(time.Hour * 72).Unix(),
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(config.JWTSecret))
// }
