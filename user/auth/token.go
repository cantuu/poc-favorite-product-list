package auth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"user/config"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(user_id uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Duration(config.ENV.TokenExpirationTime * int(time.Minute))).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.ENV.ApiSecret))

}

func ExtractTokenID(context *gin.Context) (uint32, error) {

	tokenString, err := ExtractToken(context)
	if err != nil {
		return 0, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.ENV.ApiSecret), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}
	return 0, nil
}

func ExtractToken(context *gin.Context) (string, error) {
	auth := context.Request.Header.Get("Authorization")

	if auth == "" {
		return "", errors.New("Authorization Header não passado")
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	if token == auth {
		return "", errors.New("Bearer Token não foi encontrado no Header")
	}

	return token, nil
}

// func TokenValid(r *http.Request) error {
// 	tokenString := ExtractToken(r)
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("API_SECRET")), nil
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		Pretty(claims)
// 	}
// 	return nil
// }

// //Pretty display the claims licely in the terminal
// func Pretty(data interface{}) {
// 	b, err := json.MarshalIndent(data, "", " ")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	fmt.Println(string(b))
// }
