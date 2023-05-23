FROM golang:1.20-alpine@sha256:ee2f23f1a612da71b8a4cd78fec827f1e67b0a8546a98d257cca441a4ddbebcb


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go test -v ./...

RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o main .

CMD ["./main"]