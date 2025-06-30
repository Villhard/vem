APP=vem

format:
	go fmt ./...

build:
	go build -o $(APP) main.go

install: build
	sudo mv $(APP) /usr/local/bin/
	mkdir -p ~/.config/vem
	cp config.yaml.example ~/.config/vem/config.yaml

uninstall:
	sudo rm -f /usr/local/bin/$(APP)
	rm -rf ~/.config/vem