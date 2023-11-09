//go:generate go run ./../tools/generate.go
package graph

import "github.com/0x726f6f6b6965/mono-service/services/graph-service/internal/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	authService service.UserService
}
