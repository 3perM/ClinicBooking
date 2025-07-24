FROM golang:1.22

WORKDIR /app
COPY go.mod ./ 
RUN go mod tidy
COPY . .
RUN go build -o server ./cmd/server
EXPOSE 50051
CMD ["./server"]
