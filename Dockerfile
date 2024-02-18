FROM golang:alpine AS builder

WORKDIR /app
RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates
COPY ./src ./
RUN go mod download
RUN go build -o /go/bin/server

FROM scratch
COPY --from=builder /go/bin/server /go/bin/server
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/go/bin/server"]
