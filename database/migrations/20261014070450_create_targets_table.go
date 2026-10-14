package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"uptimer/app/facades"
)

type M20261014070450CreateTargetsTable struct{}

// Signature The unique signature for the migration.
func (r *M20261014070450CreateTargetsTable) Signature() string {
	return "20261014070450_create_targets_table"
}

// Up Run the migrations.
func (r *M20261014070450CreateTargetsTable) Up() error {
	if !facades.Schema().HasTable("targets") {
		return facades.Schema().Create("targets", func(table schema.Blueprint) {
			table.TimestampTz("created_at").Nullable()
			table.TimestampTz("updated_at").Nullable()
			table.BigIncrements("id")
			table.String("name")
			table.String("url")
			table.BigInteger("timeout")
			table.BigInteger("interval")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20261014070450CreateTargetsTable) Down() error {
	return facades.Schema().DropIfExists("targets")
}
