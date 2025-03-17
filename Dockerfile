FROM golang:1.24.1-alpine

WORKDIR /app

COPY . ./

RUN go build -o /dailyemoji

EXPOSE 8080

ENTRYPOINT ["/dailyemoji"]