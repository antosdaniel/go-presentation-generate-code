BUF_VERSION := "1.12.0"
BUF_BIN := "/usr/local/bin"

install:
	@printf "Installing buf...\n"
	sudo rm -rf ${BUF_BIN}/buf && \
		sudo curl -sSL "https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$$(uname -s)-$$(uname -m)" -o "${BUF_BIN}/buf" && \
	  	sudo chmod +x "${BUF_BIN}/buf"

	@printf "\nInstalling linter...\n"
	@go install -mod=mod github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1

lint:
	@printf "Linting protos...\n"
	@buf lint

	@printf "Linting Go...\n"
	@golangci-lint run

generate:
	@printf "Generating protos...\n"
	@buf generate
