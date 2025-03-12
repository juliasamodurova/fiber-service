package dto

import (
	"github.com/gofiber/fiber/v2"
)

// константы для кодов ошибок
const (
	FieldBadFormat = "FIELD_BADFORMAT" // будем использовать потом, при создании авторизации, например
	FieldIncorrect = "FIELD_INCORRECT" // можно тоже использовать в будущем, когда данные
	// не соответстуют ожиданиям
	ServiceUnavailable = "SERVICE_UNAVAILABLE"
	InternalError      = "SERVICE_UNAVAILABLE"
)

// Response представляет общую структуру ответа API
type Response struct {
	Status string `json:"status"`
	Error  *Error `json:"error,omitempty"`
	Data   any    `json:"data,omitempty"`
}

// Error представляет ошибку с кодом и описанием
type Error struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

// BadResponseError возвращает ошибку с кодом 400 (Bad Request)
func BadResponseError(ctx *fiber.Ctx, code, desc string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(Response{
		Status: "error",
		Error: &Error{
			Code: code,
			Desc: desc,
		},
	})
}

// InternalServerError возвращает ошибку с кодом 500 (Internal Server Error)
func InternalServerError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(Response{
		Status: "error",
		Error: &Error{
			Code: ServiceUnavailable,
			Desc: InternalError,
		},
	})
}

// NotFoundError возвращает ошибку с кодом 404 (Not Found)
func NotFoundError(ctx *fiber.Ctx, desc string) error {
	return ctx.Status(fiber.StatusNotFound).JSON(Response{
		Status: "error",
		Error: &Error{
			Code: "RESOURCE_NOT_FOUND",
			Desc: desc,
		},
	})
}

// UnauthorizedError возвращает ошибку с кодом 401 (Unauthorized) // тож можно использовать для авторизации потом
func UnauthorizedError(ctx *fiber.Ctx, desc string) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(Response{
		Status: "error",
		Error: &Error{
			Code: "UNAUTHORIZED",
			Desc: desc,
		},
	})
}

// ConflictError возвращает ошибку с кодом 409 (Conflict)
func ConflictError(ctx *fiber.Ctx, desc string) error {
	return ctx.Status(fiber.StatusConflict).JSON(Response{
		Status: "error",
		Error: &Error{
			Code: "CONFLICT",
			Desc: desc,
		},
	})
}
