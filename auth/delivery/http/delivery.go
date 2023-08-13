package http

import (
	"context"
	"fmt"
	"nicetry/auth/models"
	"nicetry/auth/usecase"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthDelivery interface {
	Login(c *fiber.Ctx) error
	Me(c *fiber.Ctx) error
}

func NewAuthDelivery(jwtPriv string, uc usecase.AuthUsecase) AuthDelivery {
	return &auth{
		JWTPriv: jwtPriv,
		usecase: uc,
	}
}

type auth struct {
	JWTPriv string
	usecase usecase.AuthUsecase
}

func (a *auth) Login(c *fiber.Ctx) error {
	dataReq := new(models.LoginAttempt)
	if err := c.BodyParser(dataReq); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}
	user, err := a.usecase.Login(context.Background(), dataReq.Username, dataReq.Password)
	if err != nil {
		fmt.Println(err)
		c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"success": false,
			"message": "username or password invalid",
		})
		return err
	}
	expired := time.Now().Add(60 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":        time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local).Unix(),
		"authorized": true,
		"user":       user,
		"exp":        expired.Unix(),
	})

	tokenString, err := token.SignedString([]byte(a.JWTPriv))
	if err != nil {
		fmt.Println(err)
		c.Status(fiber.StatusInternalServerError).JSON(err)
		return err
	}
	c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"user":    user,
		"authorisation": &fiber.Map{
			"token":   tokenString,
			"expired": expired,
		},
	})
	return nil
}

func (a *auth) Me(c *fiber.Ctx) error {
	tokenString := strings.Replace(c.Get("Authorization"), "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.JWTPriv), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "selamat datang",
			"profile": claims["user"],
		})
	} else {
		fmt.Println(err)
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "unauthorized access",
		})
	}
	return nil
}
