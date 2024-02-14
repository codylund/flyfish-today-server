FROM amd64/golang

# Set destination for COPY
WORKDIR /app

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY ./src ./

RUN go mod download

# Serve the 'build' directory on port 4200 using 'serve'
CMD ["go", "run", "."]
