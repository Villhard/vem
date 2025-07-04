APP=vem

format:
	go fmt ./...

build:
	go build -o $(APP) main.go

install: build
	sudo mv $(APP) /usr/local/bin/
	mkdir -p ~/.config/vem
	cp config.yaml.example ~/.config/vem/config.yaml

update:
	@if [ ! -f /usr/local/bin/$(APP) ]; then \
		echo "Ошибка: $(APP) не установлен. Используйте 'make install' для установки."; \
		exit 1; \
	fi
	@echo "Обновление $(APP)..."
	git pull
	$(MAKE) build
	sudo mv $(APP) /usr/local/bin/
	@echo "$(APP) успешно обновлен!"

uninstall:
	sudo rm -f /usr/local/bin/$(APP)
	rm -rf ~/.config/vem