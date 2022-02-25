FROM golang:1.18rc1-alpine AS builder

RUN apk update

RUN apk add --no-cache git

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

FROM scratch

COPY --from=builder /dist/main /

EXPOSE 8008

ENTRYPOINT ["/main"]