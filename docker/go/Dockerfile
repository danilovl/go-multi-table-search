FROM golang:1.20.0-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o main -v ./cmd

WORKDIR /dist
RUN cp /build/.env .
RUN cp /build/.env.local .
RUN cp /build/main .


FROM scratch
COPY --from=builder /dist /

ENTRYPOINT ["/main"]
