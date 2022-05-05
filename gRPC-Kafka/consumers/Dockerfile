FROM golang:1.18-buster

WORKDIR /app/src
ENV GOPATH=/app
COPY . /app/src
RUN go mod download
RUN go build -o Consumer ./src

CMD ["./Consumer"]