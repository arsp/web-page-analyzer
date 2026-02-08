FROM golang:1.25.6-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o webanalyzer ./cmd/webanalyzer

EXPOSE 8080
CMD ["./webanalyzer"]
