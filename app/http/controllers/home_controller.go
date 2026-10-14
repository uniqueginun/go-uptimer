package controllers

import (
	"fmt"
	"uptimer/app/facades"
	"uptimer/app/models"

	"github.com/goravel/framework/contracts/http"

	inertiafacades "github.com/goravel/inertia/facades"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (r *HomeController) Index(ctx http.Context) http.Response {
	inertiafacades.Inertia().Defer(ctx, "stats", func() any {
		var (
			targetsCount int64
			resultsCount int64
		)

		targets, err := facades.Orm().Query().Table("targets").Count()

		if err != nil {
			ctx.Request().Session().Flash("error", fmt.Sprintf("counting targets failed: %s", err.Error()))
			targetsCount = 0
		} else {
			targetsCount = targets
		}

		results, err := facades.Orm().Query().Table("results").Count()

		if err != nil {
			ctx.Request().Session().Flash("error", fmt.Sprintf("counting results failed: %s", err.Error()))
			resultsCount = 0
		} else {
			resultsCount = results
		}

		return map[string]any{
			"targets": targetsCount,
			"results": resultsCount,
		}
	})

	inertiafacades.Inertia().Defer(ctx, "targets", func() any {
		var targs []models.Target

		facades.Orm().Query().Get(&targs)

		return targs
	})

	return inertiafacades.Inertia().Render(ctx, "Home", map[string]any{
		"message": "Monitoring Targets In Realtime",
	})
}
