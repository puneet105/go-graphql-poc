FROM golang:1.18 AS builder
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod tidy && go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server server.go


FROM alpine:latest
WORKDIR /root/
ENV PORT=8084
COPY --from=builder /usr/src/app/server ./
EXPOSE ${PORT}
CMD ["./server"]
