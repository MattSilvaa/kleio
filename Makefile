.PHONY: start-frontend install-backend build-backend start-backend start-all lint-backend clean-backend

start-all:
	@$(MAKE) start-backend & \
	$(MAKE) start-frontend & \
	wait

# Frontend tasks
start-frontend:
	bun build-client && bun start

# Backend tasks
BACKEND_DIR := server

build-backend:
	cd ${BACKEND_DIR} && go build -o dist

clean-backend:
	rm -f $(BACKEND_DIR)/dist

format-backend:
	@echo "Formatting code..."
	cd $(BACKEND_DIR) && goimports -w .

install-backend:
	cd $(BACKEND_DIR) && go mod tidy

lint-backend:
	cd $(BACKEND_DIR) && golangci-lint run ./...

start-backend:
	cd ${BACKEND_DIR} && go run main.go
