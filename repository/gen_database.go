//go:build postgres
// +build postgres

package repository

import (
	_ "github.com/kyleconroy/sqlc/cmd/sqlc"
)

//go:generate go run github.com/kyleconroy/sqlc/cmd/sqlc generate
