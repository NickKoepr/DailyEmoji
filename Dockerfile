FROM golang:1.20-alpine

WORKDIR /app

COPY . ./

RUN go build -o /dailyemoji

EXPOSE 8080

ENTRYPOINT ["/dailyemoji"]