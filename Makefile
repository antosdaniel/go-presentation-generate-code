BUF_VERSION := "1.12.0"
BUF_BIN := "/usr/local/bin"

install-system-deps:
	@printf "Installing buf...\n"
	rm -rf ${BUF_BIN}/buf && \
		curl -sSL "https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$$(uname -s)-$$(uname -m)" -o "${BUF_BIN}/buf" && \
	  	chmod +x "${BUF_BIN}/buf"

lint:
	@printf "Linting protos...\n"
	@buf lint

generate:
	@printf "Generating protos...\n"
	@buf generate
