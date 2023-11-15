# Getting all files and installing deps
FROM golang:1.20 as base

WORKDIR /src/
COPY . .

# Dev image for running locally
FROM golang:1.20 as dev

WORKDIR /src/
COPY --from=base /src/ .

RUN --mount=type=cache,target=${GOPATH}/pkg/mod/\
    make install

CMD ["make", "start-grpc"]

# Build production binary
FROM golang:1.20 as builder

WORKDIR /src/
COPY --from=base /src/ .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.Version=${version}" -mod=readonly -v -o ./bin/scheduler ./cmd/scheduler

# Production image
FROM gcr.io/distroless/base-debian11:nonroot
COPY --from=builder /src/bin/grpc /bin/grpc

CMD ["/bin/grpc"]
