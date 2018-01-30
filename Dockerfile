FROM golang:1.9.3-alpine3.7
WORKDIR /app
COPY . .
RUN go build -o hello .
EXPOSE 8080
ENTRYPOINT ["./hello"]
