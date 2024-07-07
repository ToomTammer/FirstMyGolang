package main

// Fiber คือ library ที่ได้แรงบันดาลใจมาจาก Express (ของฝั่ง node.js) ที่ build อยู่บน Fasthttp
// - Fasthttp คือ fastest HTTP engine for Go
// - เน้นไปที่ความไว และความสามารถในการจัดการ "zero memory allocation" ได้

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Handlers
///"c *fiber.Ctx" is
// - c: A variable that holds the context of the current HTTP request.
// - *fiber.Ctx: A pointer to the fiber.Ctx type, which provides an interface to the request and response.

// Handler functions
// getBooks godoc
// @Summary Get all books
// @Description Get details of all books
// @Tags books
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} Book
// @Router /book [get]
func getBooks(c *fiber.Ctx) error {
	// Retrieve user data from the context
	user := c.Locals(userContextKey).(*UserData)

	// Use the user data (e.g., for authorization, custom responses, etc.)
	fmt.Printf("User Email: %s, Role: %s\n", user.Email, user.Role)

	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	// Convert the id parameter from the URL to an integer.
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// for index, book := range books {} if there are index
	for _, book := range books {
		if book.ID == id {
			return c.JSON(book)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

// createBook handles the creation of a new book record.
func createBook(c *fiber.Ctx) error {
    // Create a new instance of the Book struct.
    book := new(Book)

    // Parse the JSON body of the request into the book struct.
    if err := c.BodyParser(book); err != nil {
        // If there's an error in parsing the body, return a 400 Bad Request status with the error message.
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    // Assign an ID to the new book. The ID is one more than the current number of books.
    book.ID = len(books) + 1

    // Append the new book to the books slice.
    books = append(books, *book)

    // Return the newly created book as a JSON response.
    return c.JSON(book)
}

// updateBook handles the updating of an existing book record.
func updateBook(c *fiber.Ctx) error {
    // Convert the id parameter from the URL to an integer.
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        // If there's an error in converting the id, return a 400 Bad Request status.
        return c.SendStatus(fiber.StatusBadRequest)
    }

    // Create a new instance of the Book struct to hold the updated data.
    bookUpdate := new(Book)

    // Parse the JSON body of the request into the bookUpdate struct.
    if err := c.BodyParser(bookUpdate); err != nil {
        // If there's an error in parsing the body, return a 400 Bad Request status with the error message.
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    // Iterate through the books slice to find the book with the matching ID.
    for i, book := range books {
        if book.ID == id {
            // Update the title and author of the found book with the new data.
            book.Title = bookUpdate.Title
            book.Author = bookUpdate.Author
            // Save the updated book back into the books slice.
            books[i] = book
            // Return the updated book as a JSON response.
            return c.JSON(book)
        }
    }

    // If the book with the specified ID is not found, return a 404 Not Found status.
    return c.SendStatus(fiber.StatusNotFound)
}


// deleteBook handles the deletion of an existing book record.
func deleteBook(c *fiber.Ctx) error {
    // Convert the id parameter from the URL to an integer.
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        // If there's an error in converting the id, return a 400 Bad Request status.
        return c.SendStatus(fiber.StatusBadRequest)
    }

    // Iterate through the books slice to find the book with the matching ID.
    for i, book := range books {
        if book.ID == id {
            // Remove the book from the slice by creating a new slice that excludes the book. books[(start):(end)]
            //example [1,2,3,4,5]
            ///[1,2] + [4,5] = [1,2,4,5]
            books = append(books[:i], books[i+1:]...)
            // Return a 204 No Content status indicating successful deletion.
            return c.SendStatus(fiber.StatusNoContent)
        }
    }

    // If the book with the specified ID is not found, return a 404 Not Found status.
    return c.SendStatus(fiber.StatusNotFound)
}
