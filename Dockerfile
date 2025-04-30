FROM golang:1.24.2 as builder
LABEL authors="tomiwa"
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o coffeeshop .

FROM scratch
COPY --from=builder /app ./

CMD ["./coffeeshop"]
