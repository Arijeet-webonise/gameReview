#/bin/bash

goose -dir=./db/migrations postgres "user=local host=127.0.0.1 dbname=gamereview sslmode=disable password=toor" up

xo pgsql://local:toor@127.0.0.1/gamereview -o app/models --suffix=.go --template-path templates/
