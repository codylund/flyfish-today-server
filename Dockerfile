FROM golang:1.22-alpine AS builder

# Set destination for COPY
WORKDIR /app

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY ./src ./

# Build the binary
RUN go mod download
RUN go build -o /go/bin/server

# Serve the 'build' directory on port 4200 using 'serve'
CMD ["go", "run", "."]

FROM scratch

# Copy the binary.
COPY --from=builder /go/bin/server /go/bin/server

# Run the hello binary.
ENTRYPOINT ["/go/bin/server"]
