FROM golang:1.16-alpine

WORKDIR /app

## Add the wait script to the image
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./database ./database
COPY ./products ./products
COPY *.go ./

RUN go build -o /priceomatic

EXPOSE 8080

## Launch the wait tool and then your application
CMD /wait && /priceomatic
#CMD [ "/priceomatic" ]
