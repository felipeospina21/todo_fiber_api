package utils

import "github.com/gofiber/fiber/v2"

func ThrowNotFoundError(err error) *fiber.Error {
	return fiber.NewError(fiber.StatusNotFound, "Not Found: "+err.Error())
}

func ThrowBodyParserError(err error) *fiber.Error {
	return fiber.NewError(fiber.StatusBadRequest, "Body Parser Err: "+err.Error())
}
