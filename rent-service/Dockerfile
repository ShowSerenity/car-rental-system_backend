FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /rent-service

EXPOSE 8082

CMD [ "/rent-service" ]
