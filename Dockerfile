FROM golang:1.20-alpine
WORKDIR /opt/app
COPY . .
RUN go mod tidy && go build cmd/generator/main.go
CMD ["./main"]
EXPOSE 8080 8002