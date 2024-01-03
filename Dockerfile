# バイナリを作成
FROM golang:1.21.5-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# -------------------------------------------------------

# デプロイ用
FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# -------------------------------------------------------

# ホットリロード
FROM golang:1.21.5 as dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]
