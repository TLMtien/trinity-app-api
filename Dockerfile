FROM golang:1.22

WORKDIR /app

ADD . /app/

# Download and install golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate.linux-amd64 /usr/local/bin/migrate && \
    chmod +x /usr/local/bin/migrate

RUN go install github.com/swaggo/swag/cmd/swag@latest 