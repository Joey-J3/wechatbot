FROM golang:latest

WORKDIR /app/wechatgpt
COPY . .

RUN go build main.go

ENTRYPOINT ["./main"]