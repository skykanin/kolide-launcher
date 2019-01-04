package table

import (
	"context"

	"github.com/boltdb/bolt"
	"github.com/kolide/launcher/pkg/osquery"
	"github.com/kolide/osquery-go/plugin/table"
)

func LauncherConfigTable(db *bolt.DB) *table.Plugin {
	columns := []table.ColumnDefinition{
		table.TextColumn("config"),
	}
	return table.NewPlugin("kolide_launcher_config", columns, generateLauncherConfig(db))
}

func generateLauncherConfig(db *bolt.DB) table.GenerateFunc {
	return func(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
		config, err := osquery.ConfigFromDB(db)
		if err != nil {
			return nil, err
		}
		results := []map[string]string{
			map[string]string{
				"config": config,
			},
		}

		return results, nil
	}
}
