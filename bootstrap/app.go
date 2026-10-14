package bootstrap

import (
	"uptimer/app/facades"
	"uptimer/app/models"

	"github.com/goravel/framework/contracts/database/schema"
	contractsfoundation "github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/foundation"

	"uptimer/config"
	"uptimer/routes"
)

func Boot() contractsfoundation.Application {
	return foundation.Setup().
		WithMigrations(Migrations).
		WithRouting(func() {
			routes.Web()
			routes.Grpc()
		}).
		WithCallback(func() {
			facades.Schema().Extend(schema.Extension{
				Models: []any{
					models.Target{}, // Already there by default
					models.Result{}, // Add your new model struct here
				},
			})
		}).
		WithProviders(Providers).
		WithConfig(config.Boot).
		Create()
}
