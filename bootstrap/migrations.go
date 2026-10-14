package bootstrap

import (
	"github.com/goravel/framework/contracts/database/schema"

	"uptimer/database/migrations"
)

func Migrations() []schema.Migration {
	return []schema.Migration{
		&migrations.M20210101000001CreateJobsTable{},
		&migrations.M20261014065756CreateTargetsTable{},
		&migrations.M20261014070435CreateResultsTable{},
		&migrations.M20261014070450CreateTargetsTable{},
	}
}
