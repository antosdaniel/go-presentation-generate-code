FROM postgres:15.2-alpine as builder

ENV VERSION v4.15.2
WORKDIR /db

RUN apk --no-cache add curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/$VERSION/migrate.linux-amd64.tar.gz \
    | tar xvz

COPY ./db/migrations ./migrations
COPY ./db/seeds ./seeds
COPY ./db/run-dev-migrations ./run-dev-migrations

CMD ["./run-dev-migrations"]

FROM alpine:3.17

WORKDIR /db

COPY --from=builder /db/migrate ./migrate
COPY ./db/migrations ./migrations

ENTRYPOINT ./migrate -path migrations -database "$DB_URL" up
