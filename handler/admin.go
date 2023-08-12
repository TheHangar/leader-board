package handler

import (
	"log"

	custometype "github.com/TheHangar/leader-board/custome_type"
	"github.com/TheHangar/leader-board/store"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AdminHandler struct {
    store *store.Storage
}

type ErrorMessage struct {
    Message string
}

func NewAdminHandler(db *store.Storage) *AdminHandler {
    return &AdminHandler{ store: db }
}

func (h *AdminHandler) HandlePostLogin(c *fiber.Ctx) error {
    username := c.FormValue("username")
    password := c.FormValue("password")

    userFound, err := h.store.User.GetAdminUserByUsername(username)

    if err != nil {
        log.Println(err)
        data := &custometype.ErrorMessage{ Message: "Database error." }
        return c.Render("components/error", data)
    }

    if userFound.Username == "" {
        data := &custometype.ErrorMessage{ Message: "User not found." }
        return c.Render("components/error", data)
    }

    if isPasswordEqual := verifyPassword(userFound.Password, password); !isPasswordEqual {
        data := &custometype.ErrorMessage{ Message: "Incorrect password." }
        return c.Render("components/error", data)
    }

    return c.Next()
}

func (h *AdminHandler) HandleCreateAdmin(c *fiber.Ctx) error {
    username := c.FormValue("username")
    password := c.FormValue("password")

    userFound, err := h.store.User.GetAdminUserByUsername(username)

    if err != nil {
        data := &custometype.ErrorMessage{ Message: "Database error." }
        return c.Render("components/error", data)
    }

    if userFound.Username != "" {
        data := &custometype.ErrorMessage{ Message: "User already exist." }
        return c.Render("components/error", data)
    }

    hashPwd, err := createHashFromPassword(password)

    if err != nil {
        data := &custometype.ErrorMessage{ Message: "Database error." }
        return c.Render("components/error", data)
    }

    newUser := &custometype.User{ Username: username, Password: hashPwd }
    _, err = h.store.User.AddAdminUser(newUser)

    if err != nil {
        data := &custometype.ErrorMessage{ Message: "Database error." }
        return c.Render("components/error", data)
    }

    return c.Next()
}

func (h *AdminHandler) HandlePostGame(c *fiber.Ctx) error {
    name := c.FormValue("game-name")

    newGame := &custometype.Game{ UUID: uuid.New().String(), Name: name }

    _, err := h.store.Game.AddGame(newGame)

    if err != nil {
        log.Println(err)
        data := &custometype.ErrorMessage{ Message: "Database error." }
        return c.Render("components/error", data)
    }

    log.Println("New game added:", newGame)
    return c.Render("components/newGame", newGame)
}

func (h *AdminHandler) HandleGetGames(c *fiber.Ctx) error {
    games, err := h.store.Game.GetGames()

    if err != nil {
        data := &custometype.ErrorMessage{ Message: "Database error." }
        return c.Render("components/error", data)
    }

    return c.Render("components/game", games)
}

func createHashFromPassword(pwd string) (string, error) {
    hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

    if err != nil {
        return "", err
    }

    return string(hashPwd), nil
}

func verifyPassword(pwdHash, pwd string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(pwd))
    return err == nil
}
