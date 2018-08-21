BUILD_DIR=bin
BIN=neo-go-sc
INSTALL_PATH=/usr/local/bin

install: deps
	@echo "installing neo-go-sc framework"
	@go build -o $(BUILD_DIR)/$(BIN) ./cli
	@cp $(BUILD_DIR)/$(BIN) $(INSTALL_PATH) 
	@echo "done installing, happy coding!"

deps: 
	@echo "installing project dependencies"
	@dep ensure

clean:
	@echo "cleaning build artifacts"
	@rm -rf $(BUILD_DIR) 

uninstall:
	@echo "uninstalling neo-go-sc framework"
	@rm -rf $(INSTALL_PATH)/$(BIN)

test:
	@echo "running tests"
	@go test ./... -cover
