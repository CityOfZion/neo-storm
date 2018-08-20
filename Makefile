BUILD_DIR=bin
EXEC_NAME=neo-go-sc

install: deps
	@echo "installing neo-go-sc framework"
	@go build -o $(BUILD_DIR)/$(EXEC_NAME) ./cli
	@echo "done installing, happy coding!"

deps: 
	@echo "installing project dependencies"
	@dep ensure

clean:
	@echo "cleaning build artifacts"
	@rm -rf $(BUILD_DIR) 
