# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go
FROM golang

# Copy the local package files to the container's workspace
ADD . /go/src/jsonparse

# Build the outyet command inside the continer.
# (You may fetch or manage depdencies here,
# either manually or with a tool like "godep".)
RUN go install jsonparse

# Run the jsonparse command by default when the container starts.
ENTRYPOINT /go/bin/jsonparse

# Document that the service listens on port 8080
#EXPOSE 8080
