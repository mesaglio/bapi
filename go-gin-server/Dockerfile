FROM golang:alpine

WORKDIR /app
ADD go.* /app/

RUN go mod download

COPY . /app/
RUN go build -o api src/main.go

EXPOSE 8080
ENV GIN_MODE=release
CMD ["/app/api"]