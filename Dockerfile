FROM golang:1.19-alpine

WORKDIR /app

COPY ./go.mod ./go.sum ./*.go ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./server

EXPOSE 8080

CMD [ "./server" ]