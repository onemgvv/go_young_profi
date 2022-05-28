include .env
.PHONY:

build_w:
	go build -o ./.bin/webhook cmd/webhook/main.go
build_p:
	go build -o ./.bin/polling cmd/polling/main.go

run_w: build_w
	./.bin/webhook
run_p: build_p
	./.bin/polling