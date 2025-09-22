package v1

import (
	"net/http"

	"github.com/evrone/go-clean-template/internal/controller/http/v1/request"
	"github.com/evrone/go-clean-template/internal/entity"
	"github.com/evrone/go-clean-template/internal/usecase"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type TranslationHandler struct {
	t usecase.Translation
	l logger.Interface
	v *validator.Validate
}

func NewTranslationHandler(t usecase.Translation, l logger.Interface) *TranslationHandler {
	return &TranslationHandler{
		t: t,
		l: l,
		v: validator.New(validator.WithRequiredStructEnabled()),
	}
}

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} entity.TranslationHistory
// @Failure     500 {object} response.Error
// @Router      /translation/history [get]
func (r *TranslationHandler) history(ctx fiber.Ctx) error {
	translationHistory, err := r.t.History(ctx)
	if err != nil {
		r.l.Error().Err(err).Msg("http - v1 - doTranslate")

		return errorResponse(ctx, http.StatusInternalServerError, "database problems")
	}

	return ctx.Status(http.StatusOK).JSON(translationHistory)
}

// @Summary     Translate
// @Description Translate a text
// @ID          do-translate
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Param       request body request.Translate true "Set up translation"
// @Success     200 {object} entity.Translation
// @Failure     400 {object} response.Error
// @Failure     500 {object} response.Error
// @Router      /translation/do-translate [post]
func (r *TranslationHandler) doTranslate(ctx fiber.Ctx) error {
	var body request.Translate

	if err := ctx.Bind().Body(&body); err != nil {
		r.l.Error().Err(err).Msg("http - v1 - doTranslate")

		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	if err := r.v.Struct(body); err != nil {
		r.l.Error().Err(err).Msg("http - v1 - doTranslate")

		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	translation, err := r.t.Translate(
		ctx,
		entity.Translation{
			Source:      body.Source,
			Destination: body.Destination,
			Original:    body.Original,
		},
	)
	if err != nil {
		r.l.Error().Err(err).Msg("http - v1 - doTranslate")

		return errorResponse(ctx, http.StatusInternalServerError, "translation service problems")
	}

	return ctx.Status(http.StatusOK).JSON(translation)
}
