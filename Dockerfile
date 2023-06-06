FROM golang:1.20

WORKDIR /app

COPY . .
RUN go mod tidy

RUN go build -o main
RUN chmod +x main

EXPOSE 9091

ENTRYPOINT ["./main"]