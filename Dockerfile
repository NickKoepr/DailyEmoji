FROM golang:1.24rc3-alpine

WORKDIR /app

COPY . ./

RUN go build -o /dailyemoji

EXPOSE 8080

ENTRYPOINT ["/dailyemoji"]