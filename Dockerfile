# builder image
FROM golang:1.22-alpine AS builder
RUN mkdir /build
WORKDIR /build
COPY go.mod go.sum ./
# you may use `GOPROXY` to speed it up in Mainland China.
# RUN  GOPROXY=https://goproxy.cn,direct go mod download
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o 0g-scan .

# final target image for multi-stage builds
FROM alpine:3.18
RUN apk --no-cache add ca-certificates
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=builder /build/0g-scan .
COPY ./config/config.yml ./config.yml
ENTRYPOINT [ "./0g-scan" ]
CMD [ "--help" ]