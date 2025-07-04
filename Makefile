APP=vem

format:
	@go fmt ./...

build:
	@go build -o $(APP) main.go

install:
	@echo "Install $(APP)..."
	@$(MAKE) build
	@sudo mv $(APP) /usr/local/bin/
	@mkdir -p ~/.config/vem
	@cp config.yaml.example ~/.config/vem/config.yaml
	@echo "Installed!"

update:
	@if [ ! -f /usr/local/bin/$(APP) ]; then \
		echo "Error: $(APP) don't install. Use 'make install' for install"; \
		exit 1; \
	fi
	@echo "Update $(APP)..."
	@git pull
	@$(MAKE) build
	@sudo mv $(APP) /usr/local/bin/
	@echo "Updated!"

uninstall:
	@echo "Uninstall $(APP)..."
	@sudo rm -f /usr/local/bin/$(APP)
	@rm -rf ~/.config/vem
	@echo "Uninstalled!"