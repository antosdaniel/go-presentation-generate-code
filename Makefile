BUF_VERSION := "1.12.0"
BUF_BIN := "/usr/local/bin"

.PHONY: dev
dev:
	@docker compose down
	@make -s deps
	@docker compose up

.PHONY: install
install:
	@printf "Installing buf...\n"
	# This works without `sudo` in Docker. If you want to run it locally, you probably will have to prefix each line with `sudo`.
	@rm -rf ${BUF_BIN}/buf && \
		curl -sSL "https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$$(uname -s)-$$(uname -m)" -o "${BUF_BIN}/buf" && \
	  	chmod +x "${BUF_BIN}/buf"

	@printf "\nInstalling linter...\n"
	@go install -mod=readonly github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2

	@printf "\nInstalling compile daemon...\n"
	@go install -mod=readonly github.com/githubnemo/CompileDaemon@v1.4.0

	@printf "\nInstalling sqlboiler...\n"
	@go install -mod=readonly github.com/volatiletech/sqlboiler/v4@v4.14.2
	@go install -mod=readonly github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.14.2

	@printf "\nInstalling gowrap...\n"
	@go install -mod=readonly github.com/hexdigest/gowrap/cmd/gowrap@v1.3.2

	@printf "\nInstalling mockgen...\n"
	@go install -mod=readonly github.com/golang/mock/mockgen@v1.7.0-rc.1

.PHONY: lint
lint:
	@printf "Linting protos...\n"
	@buf lint --config internal/grpc/buf.yaml

	@printf "Linting Go...\n"
	@golangci-lint run

.PHONY: deps
deps:
	@printf "Refreshing Go modules...\n"
	@go mod tidy && go mod vendor

.PHONY: generate
generate:
	@printf "Generating protos...\n"
	@buf generate --template internal/grpc/buf.gen.yaml

	@printf "Generating db models...\n"
	@sqlboiler --config db/sqlboiler.toml psql

	@printf "go generate...\n"
	@go generate ./...

	@make -s deps

	@printf "Formatting files...\n"
	@golangci-lint run --fix

GRPC_BINARY := "bin/grpc"

.PHONY: start-grpc
start-grpc:
	@CompileDaemon \
		-build="go build -o ${GRPC_BINARY} cmd/grpc/main.go" \
		-command="./${GRPC_BINARY}" \
		-log-prefix=false \
		-graceful-kill=true \
		-color=true \
		-polling
