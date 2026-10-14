package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"uptimer/app/facades"
)

type M20261014070435CreateResultsTable struct{}

// Signature The unique signature for the migration.
func (r *M20261014070435CreateResultsTable) Signature() string {
	return "20261014070435_create_results_table"
}

// Up Run the migrations.
func (r *M20261014070435CreateResultsTable) Up() error {
	if !facades.Schema().HasTable("results") {
		return facades.Schema().Create("results", func(table schema.Blueprint) {
			table.TimestampTz("created_at").Nullable()
			table.TimestampTz("updated_at").Nullable()
			table.BigIncrements("id")
			table.UnsignedBigInteger("target_id")
			table.BigInteger("http_code")
			table.BigInteger("latency")
			table.Text("err")
			table.TimestampTz("checked_at")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20261014070435CreateResultsTable) Down() error {
	return facades.Schema().DropIfExists("results")
}
