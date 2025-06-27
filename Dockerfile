FROM golang:latest

WORKDIR /cmd

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN cd cmd && go build -o ../main .

EXPOSE 8080

CMD ["./main"]