FROM golang:latest
WORKDIR /app
COPY go.mod index.html ./
RUN go mod download
COPY *.go ./
RUN go build -o /gohtmx
EXPOSE 8080
CMD ["/gohtmx"]
