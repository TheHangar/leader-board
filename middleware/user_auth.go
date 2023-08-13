package middleware

import (
	"fmt"
	"os"
	"time"

	custometype "github.com/TheHangar/leader-board/custome_type"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var secret = os.Getenv("JWT_SECRET")

type AuthUserHandler struct {}

type unauthorizedResponse struct {
    PageTitle string
}

func NewAuthUserHandler() *AuthUserHandler {
    return &AuthUserHandler{}
}

func (a *AuthUserHandler) CreateToken(c *fiber.Ctx) error {
    // create token and pass in http only cookie
    tkn := createJWT()

    if tkn == "" {
        data := &custometype.ErrorMessage{ Message: "Internal server error." }
        return c.Render("components/error", data)
    }

    cookie := new(fiber.Cookie)
    cookie.Name = "X-Api-Token"
    cookie.Value = tkn
    cookie.HTTPOnly = true
    cookie.Expires = time.Now().Add(2 * time.Hour)

    c.Cookie(cookie)
    c.Response().Header.Add("HX-Push", "/")
    c.Response().Header.Add("HX-Redirect", "/")

    return c.Render("components/error", nil)
}

func (a *AuthUserHandler) VerifyToken(c *fiber.Ctx) error {
    tkn := c.Cookies("X-Api-Token")

    if isTokenValid := validJWT(tkn); !isTokenValid {
        c.Response().Header.Add("HX-Push", "/login")
        c.Response().Header.Add("HX-Redirect", "/login")
        data := &unauthorizedResponse{ PageTitle: "The hangar - unauthorized" }
        return c.Status(401).Render("unauthorized", data, "layouts/main")
    }

    return c.Next()
}

func validJWT(tkn string) bool {
    if tkn == "" {
        return false
    }

    token, _ := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unauthorized")
        }

        return []byte(secret), nil
    })

    if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return true
    } else {
        return false
    }
}

func createJWT() string {
    timestamp := time.Now()

    claims := jwt.MapClaims{
        "iat": timestamp.Unix(),
        "exp": timestamp.Add(time.Hour * 2).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
    tokenString, err := token.SignedString([]byte(secret))

    if err != nil {
        return ""
    }

    return tokenString
}
