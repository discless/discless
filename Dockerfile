FROM golang:1.16

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o /discless main.go

EXPOSE 8080

CMD ["/discless"]