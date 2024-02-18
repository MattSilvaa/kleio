# Frontend tasks
start-frontend:
	bun build-client && bun start

# Backend tasks
BACKEND_DIR := server

build-backend:
	cd ${BACKEND_DIR} && go build -o dist

clean:
	rm -f $(BACKEND_DIR)/dist

install-backend:
	cd $(BACKEND_DIR) && go mod tidy

lint-backend:
	cd $(BACKEND_DIR) && golangci-lint run ./...

start-backend:
	cd ${BACKEND_DIR} && go run main.go


start:
	cd $(BACKEND_DIR) && go run main.go &
	bun build-client && bun start

.PHONY: build-frontend start-frontend install-backend build-backend start-backend start test-backend lint-backend clean