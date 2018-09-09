FROM golang:1.11 as builder
COPY . /go/src/github.com/cagataygurturk/hellogo/
WORKDIR /go/src/github.com/cagataygurturk/hellogo/
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5/dep-linux-amd64 && \
    chmod +x /usr/local/bin/dep && \
    dep ensure -vendor-only
COPY app.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/cagataygurturk/hellogo/app .
ENV GIN_MODE release
ENV PORT 8080
CMD ["./app"]
EXPOSE 8080