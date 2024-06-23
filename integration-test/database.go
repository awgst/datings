package integrationtest

import (
	"fmt"

	"gorm.io/gorm"
)

func TruncateDatabase(db *gorm.DB) {
	if err := db.Exec("SET FOREIGN_KEY_CHECKS = 0").Error; err != nil {
		panic(err)
	}

	type Table struct {
		Name string `gorm:"column:TABLE_NAME"`
	}
	var tables []Table
	if err := db.Raw(`
		SELECT 
			TABLE_NAME
		FROM
			information_schema.tables
		WHERE
			TABLE_SCHEMA='datings'
		AND
			TABLE_NAME != 'schema_migrations'
	`).Scan(&tables).Error; err != nil {
		panic(err)
	}

	for _, table := range tables {
		if err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table.Name)).Error; err != nil {
			panic(err)
		}
	}
}
