FROM golang:1.27.1-trixie AS builder

WORKDIR /app

COPY go.mod go.sum ./
# COPY . .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /app/server /app/cmd/main.go

FROM alpine:3.24

WORKDIR /srv

COPY --from=builder /app/server ./server

RUN mkdir -p ./public/img ./public/doc

CMD [ "/srv/server" ]