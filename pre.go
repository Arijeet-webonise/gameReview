package todo

//go:generate goose up

//go:generate xo pgsql://local:local@localhost/todo?sslmode=disable -o internal/models --suffix=.go --template-path templates/
