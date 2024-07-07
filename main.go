package main

// Fiber คือ library ที่ได้แรงบันดาลใจมาจาก Express (ของฝั่ง node.js) ที่ build อยู่บน Fasthttp
// - Fasthttp คือ fastest HTTP engine for Go
// - เน้นไปที่ความไว และความสามารถในการจัดการ "zero memory allocation" ได้

// # run แบบปกติ go run  (file name)
// go run  *.go (api ต้องรันทุกไฟล์)
// go run .

// # run แบบ build
// go build

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	_ "github.com/toomtam/go-example/docs" // load generated docs
)

// Book struct to hold book data
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book // Slice to store books

// @title Book API
// @description This is a sample server for a book API.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Initialize standard Go html template engine
    engine := html.New("./views", ".html")

    // Pass the engine to Fiber
    app := fiber.New(fiber.Config{
        Views: engine,
    })

    // Enable CORS
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*", // You can specify specific origins instead of "*"
        AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
    }))

    app.Get("/swagger/*", swagger.HandlerDefault) // default
    

    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

     // JWT Secret Key
    secretKey := "secret"

    // Login route
    app.Post("/login", login(secretKey))

    // JWT Middleware
    app.Use(jwtware.New(jwtware.Config{
        SigningKey: []byte(secretKey),
    }))

    // Middleware to extract user data from JWT
    app.Use(extractUserFromJWT)
  
    // Group routes under /book
    bookGroup := app.Group("/book")

    // Apply the isAdmin middleware only to the /book routes
    bookGroup.Use(isAdmin)

    


	// Initialize in-memory data
	books = append(books, Book{ID: 1, Title: "1984", Author: "George Orwell"})
	books = append(books, Book{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"})


    

    // Setup route
    app.Get("/", renderTemplate)
    app.Get("/api/config", getConfig)
    app.Post("/upload", uploadImage)
  
    // Now, only authenticated admins can access these routes
    bookGroup.Get("/", getBooks)
    bookGroup.Get("/:id", getBook)
    bookGroup.Post("/", createBook)
    bookGroup.Put("/:id", updateBook)
    bookGroup.Delete("/:id", deleteBook)

	// app.Listen(":8080")
    
    // Use the environment variable for the port, default to 8080 if not set
    // port := getEnv("PORT", "8080")

    // Use the environment variable for the port
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if not specified
    }

    app.Listen(":" + port)
    
}

// Dummy user for example
var user = struct {
    Email    string
    Password string
}{
    Email:    "user@example.com",
    Password: "password123",
}

// login handles the user login and JWT token generation.
func login(secretKey string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Define the structure for the login request payload.
        type LoginRequest struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }

        // Create an instance to hold the parsed login request.
        var request LoginRequest

        // Parse the JSON body of the request into the login request struct.
        if err := c.BodyParser(&request); err != nil {
            // If there's an error in parsing the body, return the error.
            return err
        }

        // Check credentials - In a real-world application, you should verify against a database.
        if request.Email != user.Email || request.Password != user.Password {
            // If credentials are invalid, return an unauthorized error.
            return fiber.ErrUnauthorized
        }

        // Create a new JWT token with the HS256 signing method.
        token := jwt.New(jwt.SigningMethodHS256)

        // Set the claims for the token.
        claims := token.Claims.(jwt.MapClaims)
        claims["email"] = user.Email
        claims["role"] = "admin" // example role
        claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Expiration time (72 hours from now)

        // Generate the encoded token string.
        t, err := token.SignedString([]byte(secretKey))
        if err != nil {
            // If there's an error in generating the token, return a 500 Internal Server Error status.
            return c.SendStatus(fiber.StatusInternalServerError)
        }

        // Return the generated token as a JSON response.
        return c.JSON(fiber.Map{"token": t})
    }
}

// UserData represents the user data extracted from the JWT token
type UserData struct {
    Email string
    Role  string
}
  
  // userContextKey is the key used to store user data in the Fiber context
  const userContextKey = "user"
  
// extractUserFromJWT is a middleware that extracts user data from the JWT token
func extractUserFromJWT(c *fiber.Ctx) error {
    user := &UserData{}

    // Extract the token from the Fiber context (inserted by the JWT middleware)
    token, ok := c.Locals("user").(*jwt.Token)
    if !ok {
        return fiber.ErrUnauthorized
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return fiber.ErrUnauthorized
    }

    // Safely extract email and role claims
    if email, ok := claims["email"].(string); ok {
        user.Email = email
    } else {
        return fiber.ErrUnauthorized
    }

    if role, ok := claims["role"].(string); ok {
        user.Role = role
    } else {
        return fiber.ErrUnauthorized
    }

    // Store the user data in the Fiber context
    c.Locals(userContextKey, user)

    return c.Next()
}

// isAdmin checks if the user is an admin
func isAdmin(c *fiber.Ctx) error {
    user := c.Locals(userContextKey).(*UserData)

    if user.Role != "admin" {
        return fiber.ErrUnauthorized
    }

    return c.Next()
}

func uploadImage(c *fiber.Ctx) error {
    // Read the file from the request.
    file, err := c.FormFile("image")
    if err != nil {
        // If there's an error in reading the file, return a 400 Bad Request status with the error message.
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    // Save the file to the server.
    err = c.SaveFile(file, "./uploads/" + file.Filename)
    if err != nil {
        // If there's an error in saving the file, return a 500 Internal Server Error status with the error message.
        return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
    }

    // Return a success message with the filename.
    return c.SendString("File uploaded successfully: " + file.Filename)
}
