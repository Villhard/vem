APP=vem

format:
	go fmt ./...

build:
	go build -o $(APP) main.go

install: build
	sudo mv $(APP) /usr/local/bin/

uninstall:
	sudo rm /usr/local/bin/$(APP)