FROM golang:alpine AS builder

# Set destination for COPY
WORKDIR /app

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY ./src ./

# Build the binary
RUN go mod download
RUN go build -o /go/bin/server

# Download latest CA certs
RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates

FROM scratch

# Copy the binary.
COPY --from=builder /go/bin/server /go/bin/server

# Copy CA certs.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Run the hello binary.
ENTRYPOINT ["/go/bin/server"]
