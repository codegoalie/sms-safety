package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gobuffalo/envy"
	"github.com/gofiber/fiber"
	"github.com/sfreiberg/gotwilio"
)

func main() {

	from, err := envy.MustGet("TWILIO_NUMBER")
	if err != nil {
		err = fmt.Errorf("failed to load TWILIO_NUMBER: %w", err)
		log.Fatal(err)
	}
	to, err := envy.MustGet("RECEIVER_NUMBER")
	if err != nil {
		err = fmt.Errorf("failed to load RECEIVER_NUMBER: %w", err)
		log.Fatal(err)
	}

	accountSID, err := envy.MustGet("TWILIO_ACCOUNT_SID")
	if err != nil {
		err = fmt.Errorf("failed to load TWILIO_ACCOUNT_SID: %w", err)
		log.Fatal(err)
	}
	authToken, err := envy.MustGet("TWILIO_AUTH_TOKEN")
	if err != nil {
		err = fmt.Errorf("failed to load TWILIO_AUTH_TOKEN: %w", err)
		log.Fatal(err)
	}
	twilioClient := gotwilio.NewTwilioClient(accountSID, authToken)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, world!")
	})

	app.Post("/sms", func(c *fiber.Ctx) {
		origFrom := c.Body("From")
		origBody := c.Body("Body")

		message := fmt.Sprintf("SMSafety msg from %s:\n %s", origFrom, origBody)
		twilioClient.SendSMS(from, to, message, "", "")

		c.Status(http.StatusOK)
	})

	app.Listen(envy.Get("PORT", "8080"))
}
