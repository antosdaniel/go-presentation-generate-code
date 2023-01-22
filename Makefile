BUF_VERSION := "1.12.0"
BUF_BIN := "/usr/local/bin"

install:
	@printf "Installing buf...\n"
	# You might need to run it with sudo
	rm -rf ${BUF_BIN}/buf && \
		curl -sSL "https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$$(uname -s)-$$(uname -m)" -o "${BUF_BIN}/buf" && \
	  	chmod +x "${BUF_BIN}/buf"

	@printf "\nInstalling linter...\n"
	@go install -mod=mod github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1

	@printf "\nInstalling compile daemon...\n"
	@go install -mod=readonly github.com/githubnemo/CompileDaemon@v1.4.0

lint:
	@printf "Linting protos...\n"
	@buf lint

	@printf "Linting Go...\n"
	@golangci-lint run

generate:
	@printf "Generating protos...\n"
	@buf generate

GRPC_BINARY := "bin/grpc"

start-grpc:
	@CompileDaemon \
		-build="go build -o ${GRPC_BINARY} cmd/grpc/main.go" \
		-command="./${GRPC_BINARY}" \
		-log-prefix=false \
		-graceful-kill=true \
		-color=true \
		-polling
