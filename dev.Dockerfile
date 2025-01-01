# Build stage
FROM golang:1.22.1-alpine AS build

WORKDIR /app/src

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN go build -o bin/guestbook ./cmd/guestbook


# Final image
FROM alpine:3

WORKDIR /bin

COPY --from=build /app/src/bin/guestbook .

ENTRYPOINT ["/bin/guestbook"]