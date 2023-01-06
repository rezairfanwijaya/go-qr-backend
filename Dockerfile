FROM golang:1.20-rc-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o go-qr

ENTRYPOINT [ "./go-qr" ]