FROM golang:1.16-alpine

WORKDIR /app

COPY Nexo.go ./

RUN go build Nexo.go

EXPOSE 8080

CMD [ "sh", "-c", "./Nexo" ]