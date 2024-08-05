package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	fiberApp := fiber.New()
	setupControllers(fiberApp)

	// Run the server at port 3000
	fiberApp.Listen(":3000")
}

func setupControllers(app *fiber.App) {

	app.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("POST")
	})

	// GET http://127.0.0.1:3000/hello [Hello World.]
	// GET http://127.0.0.1:3000/hello/there [Hello there.]
	// GET http://127.0.0.1:3000/hello/there?msg=!!! [Hello there!!!]
	app.Get("/hello/:name?", func(ctx *fiber.Ctx) error {
		// Extract path param
		name := ctx.Params("name")
		if name == "" {
			name = " World"
		} else {
			name = " " + name // Add a space in front if not empty
		}

		// Extract query param (if none is found, default value is used)
		message := ctx.Query("msg", ".")

		return ctx.SendString("Hello" + name + message)
	})

	// POST http://127.0.0.1:3000/jsonResponse
	// Receive json request body and return that json back to user
	app.Post("/jsonResponse", func(ctx *fiber.Ctx) error {
		var jsonData fiber.Map

		if err := ctx.BodyParser(&jsonData); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(
				fiber.Map{
					"error": "Cannot parse JSON",
				},
			)
		}
		return ctx.JSON(jsonData)
	})

	// POST http://127.0.0.1:3000/uploadFile
	app.Post("/uploadFile", func(ctx *fiber.Ctx) error {
		// Parse the multipart form:
		form, err := ctx.MultipartForm()
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse multipart form",
			})
		}

		// Extract the file
		files := form.File["file"]
		if len(files) == 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "No file uploaded",
			})
		}

		// Extract file name
		file := files[0]
		fileName := file.Filename

		return ctx.SendString("Uploaded fileName: " + fileName)
	})

	// POST http://127.0.0.1:3000/sendEmail
	app.Post("/sendEmail", func(ctx *fiber.Ctx) error {
		// Some other libraries to send email with more configuration (e.g. TLS, Bcc, Importance level, HTML template)
		// a) GoMail - https://github.com/wneessen/go-mail (https://dev.to/wneessen/sending-mails-in-go-the-easy-way-1lm7)

		// sendEmail()    // Send email synchronously
		// go sendEmail() // Send email asynchronously with Goroutines

		return nil // Equal to httpStatusCode 200
	})
}

/* func sendEmail() {
	// Configure smtp server config
	// 	smtpServer   = "smtp.gmail.com"
	//	smtpPort     = "587"
	//	smtpUsername = [YOUR_EMAIL_ADDRESS]
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)

	// Compose the email
	emailBody := "To: " + userEmail + "\r\n" +
		"Subject: " + emailSubject + "\r\n\r\n" +
		emailBody

	// Send email
	sendErr := smtp.SendMail(smtpServer+":"+smtpPort, auth, smtpUsername, userEmail, []byte(emailBody))
	if sendErr != nil {
		fmt.Println("Failed to send email.")
	}
	fmt.Println("Email sent successfully to: " + userEmail)
} */
