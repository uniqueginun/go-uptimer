package controllers

import (
	"fmt"
	"time"
	"uptimer/app/facades"
	"uptimer/app/models"

	"github.com/goravel/framework/contracts/http"

	inertiafacades "github.com/goravel/inertia/facades"
)

type TargetsController struct{}

func NewTargetsController() *TargetsController {
	return &TargetsController{}
}

func (r *TargetsController) Create(ctx http.Context) http.Response {
	var targets []models.Target
	if err := facades.Orm().Query().Order("id desc").Find(&targets); err != nil {
		facades.Log().Errorf("failed to load targets: %v", err)
	}

	return inertiafacades.Inertia().Render(ctx, "TargetsCreate", map[string]any{
		"targets": targets,
	})
}

func (r *TargetsController) Store(ctx http.Context) http.Response {
	validator, err := ctx.Request().Validate(map[string]any{
		"name":     "string|required|max:255",
		"url":      "required|url",
		"timeout":  "required|integer|min:1|max:30",
		"interval": "required|integer|min:10",
	})

	if err != nil || validator.Fails() {
		inertiafacades.Inertia().FlashErrors(ctx, validator.Errors())
		return inertiafacades.Inertia().Redirect(ctx, "/targets")
	}

	target := models.Target{
		Name:     ctx.Request().Input("name"),
		Url:      ctx.Request().Input("url"),
		Interval: time.Second * time.Duration(ctx.Request().InputInt("interval")),
		Timeout:  time.Second * time.Duration(ctx.Request().InputInt("timeout")),
	}

	if err := facades.Orm().Query().Create(&target); err != nil {
		ctx.Request().Session().Flash("error", fmt.Sprintf("Target create failed: %s", err.Error()))
		return inertiafacades.Inertia().Redirect(ctx, "/targets")
	}

	ctx.Request().Session().Flash("success", "Target successfully created!")

	return inertiafacades.Inertia().Redirect(ctx, "/targets")
}

func (r *TargetsController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().RouteInt("id")

	result, err := facades.Orm().Query().Where("id", id).Delete(&models.Target{})

	if err != nil {
		ctx.Request().Session().Flash("error", fmt.Sprintf("Target deletion failed: %s", err.Error()))
	} else if result.RowsAffected == 0 {
		ctx.Request().Session().Flash("error", "Target does not exist!")
	} else {
		ctx.Request().Session().Flash("success", "Target was deleted successfully!")
	}

	return inertiafacades.Inertia().Redirect(ctx, "/targets")
}
