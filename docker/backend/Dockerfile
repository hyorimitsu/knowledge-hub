FROM golang:1.24.0-alpine3.21

ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV AIR_VERSION=v1.61.7
ENV DELVE_VERSION=v1.24.0
ENV MIGRATE_VERSION=v4.17.0
ENV SWAG_VERSION=v1.16.3

RUN apk add --no-cache curl

# Install Air for hot reload
RUN go install github.com/air-verse/air@"${AIR_VERSION}"

# Install Delve for debugging
RUN go install github.com/go-delve/delve/cmd/dlv@"${DELVE_VERSION}"

# Install golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/migrate && \
    chmod +x /usr/local/bin/migrate

# Install swag for OpenAPI documentation
RUN go install github.com/swaggo/swag/cmd/swag@"${SWAG_VERSION}"