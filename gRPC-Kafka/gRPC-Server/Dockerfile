FROM golang:1.18-buster

WORKDIR /app/src
ENV GOPATH=/app
COPY . /app/src
RUN go mod download
RUN go build -o Server ./src
EXPOSE 50051

CMD ["./Server"]