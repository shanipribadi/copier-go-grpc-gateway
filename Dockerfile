# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.22 as builder
# Create and change to the app directory.
WORKDIR /src
# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
COPY go.* ./
RUN go mod download
# Copy local code to the container image.
COPY . ./
# Build the binary.
# -mod=readonly ensures immutable go.mod and go.sum in container builds.
RUN CGO_ENABLED=1 GOOS=linux go build -mod=readonly -v -o go-cookiecutter

# Build the runtime container image from scratch, copying what is needed from the previous stage.  
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM golang:1.22
# Copy the binary to the production image from the builder stage.
COPY --from=builder /src/go-cookiecutter /app/go-cookiecutter
# Run the web service on container startup.
ENTRYPOINT ["/app/go-cookiecutter"]
