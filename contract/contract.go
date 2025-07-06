package contract

import (
	"gorm.io/gen"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE email = @email{{if name !=""}} AND name = @name{{end}}
	FilterWithNameAndEmail(name, email string) ([]gen.T, error)
}
