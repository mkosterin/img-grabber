.PHONY:
.SILENT:

build:
	go build -o ./.bin/img-grabber cmd/img-grabber.go

run: build
	./.bin/img-grabber