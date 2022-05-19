//go:build postgres
// +build postgres

package postgres

import (
	_ "github.com/kyleconroy/sqlc/cmd/sqlc"
)

//go:generate go run github.com/kyleconroy/sqlc/cmd/sqlc generate
