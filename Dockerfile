FROM golang:1.22-bullseye as builder

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc-aarch64-linux-gnu \
    crossbuild-essential-arm64 \
    librdkafka-dev

COPY . .

RUN apt-get update && apt-get install -y musl-tools

# Устанавливаем musl-gcc
RUN ln -s /usr/bin/musl-gcc /usr/local/bin/musl-gcc

RUN apt-get clean && rm -rf /var/lib/apt/lists/*

RUN go mod tidy

RUN CC=/usr/local/bin/musl-gcc go build --ldflags '-linkmode external -extldflags "-static"' -tags musl -o /app-run /app/cmd/ims/main.go

CMD ["/app-run"]
