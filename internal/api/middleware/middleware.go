package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// Authorization middleware — для будущей реализации авторизации
func Authorization(token string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// здесь будет логика проверки токена
		// пока пропускаем запрос дальше
		return ctx.Next()
	}
}
