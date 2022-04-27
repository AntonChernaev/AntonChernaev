FROM golang:1.16-alpine

WORKDIR /app

#COPY Nexo.go ./
COPY Nexo-web-server.go ./
COPY go.mod ./

RUN go mod download github.com/lib/pq

#RUN go build Nexo.go
RUN go build Nexo-web-server.go

EXPOSE 8080
CMD [ "sh", "-c", "./Nexo-web-server $DBHOST" ]

#docker build --tag nexo-app . - C:\Users\achernaev\Desktop\Nexo_zadacha\AntonChernaev\GolangApp