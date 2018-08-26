.DEFAULT_GOAL := help
BUILD_DIR=bin
BIN=neo-storm
INSTALL_PATH=/usr/local/bin

help:          ## Show available options with this Makefile
	@grep -F -h "##" $(MAKEFILE_LIST) | grep -v grep | awk 'BEGIN { FS = ":.*?##" }; { printf "%-18s  %s\n", $$1,$$2 }'

install: deps ## Build and install neo-storm cli application
	@echo "installing neo-storm framework"
	@go build -o $(BUILD_DIR)/$(BIN) ./cli
	@cp $(BUILD_DIR)/$(BIN) $(INSTALL_PATH)
	@echo "done installing, happy coding!"

deps:   ## Build all the dependencies.
	@echo "installing project dependencies"
	@dep ensure

clean:  ## Clean the build-directory
	@echo "cleaning build artifacts"
	@go clean -i ./...
	@rm -rf $(BUILD_DIR)

uninstall: ## Uninstall the install binary
	@echo "uninstalling neo-storm framework"
	@rm -rf $(INSTALL_PATH)/$(BIN)

test:   ## Execute all the tests and generate cover reports
	@echo "running tests"
	@go test ./... -cover
