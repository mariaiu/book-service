FROM golang:1.17.1-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o book ./cmd/main.go

EXPOSE 9080
EXPOSE 3306
ENTRYPOINT ["./book"]