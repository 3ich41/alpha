############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
COPY . $GOPATH/src/mypackage/myapp/
WORKDIR $GOPATH/src/mypackage/myapp/
# Fetch dependencies.
# Using go get.
RUN git config --global http.sslVerify false
RUN go get -insecure -d -v
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o /go/bin/coa
#RUN GOOS=linux GOARCH=amd64 go build -ldflags '-linkmode external -extldflags -static' -o /go/bin/hello
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/coa /go/bin/coa

EXPOSE 50081

# Run the hello binary.
ENTRYPOINT ["/go/bin/coa"]
