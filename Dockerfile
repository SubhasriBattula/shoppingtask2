FROM golang:alpine
# create a working directory
WORKDIR /go/src/app

COPY go.mod./ go.mod./

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD["./main"]
