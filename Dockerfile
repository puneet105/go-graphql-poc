FROM golang:1.18 AS builder
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod tidy && go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server server.go
#CGO_ENABLED=0 . This means that the resulting binary will not be linked to any C libraries.
#$GOOS and $GOARCH The name of the target operating system and compilation architecture.
#-installsuffix cgo is a suffix to use in the name of the package installation directory,in order to keep output separate from default builds.


FROM alpine:latest
WORKDIR /root/
ENV PORT=8084
COPY --from=builder /usr/src/app/server ./
EXPOSE ${PORT}
CMD ["./server"]
