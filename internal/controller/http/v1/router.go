package v1

import (
	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/gofiber/fiber/v3"
)

// NewTranslationRoutes -.
func NewTranslationRoutes(apiV1Group fiber.Router, t usecase.Translation, l logger.Interface) {
	translationHandler := NewTranslationHandler(t, l)

	translationGroup := apiV1Group.Group("/translation")

	{
		translationGroup.Get("/history", translationHandler.history)
		translationGroup.Post("/do-translate", translationHandler.doTranslate)
	}
}
