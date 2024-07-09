package project

//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc@latest generate
//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest  --target internal/server/openapi -package openapi --clean docs/api.yaml  --config cfg.yaml
