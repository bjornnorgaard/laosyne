//go:build graphql
// +build graphql

package graphql

import (
	_ "github.com/99designs/gqlgen"
)

//go:generate go mod tidy
//go:generate go run github.com/99designs/gqlgen generate
