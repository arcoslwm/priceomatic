FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./database ./database
COPY ./products ./products
COPY *.go ./

RUN go build -o /priceomatic

EXPOSE 8080

CMD [ "/priceomatic" ]
