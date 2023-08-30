FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main .

CMD ["./main"]

## docker build -t elevator-simulation .
## docker run -it elevator-simulation

