FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go .
COPY database/*.go ./database/

RUN go build -o nome-que-escolhi

EXPOSE 8080

CMD [ "./nome-que-escolhi" ]