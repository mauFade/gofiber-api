package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ResponseInterface struct {
	ID      string
	Message string
	User    string
}

func NewMessage(m string, usr string) *ResponseInterface {
	return &ResponseInterface{
		ID:      uuid.NewString(),
		Message: m,
		User:    usr,
	}
}

func main() {
	app := fiber.New()

	app.Get("/", func(context *fiber.Ctx) error {
		response := ResponseInterface{
			ID:      uuid.NewString(),
			Message: "Testing",
			User:    "Mauricio",
		}

		user, err := json.Marshal(response)

		if err != nil {
			panic(err)
		}

		return context.SendString(string(user))
	})

	app.Listen(":9090")
}
