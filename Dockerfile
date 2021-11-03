# Use golang image as builder
FROM golang:1.17-alpine
# Make project directory inside the GOPATH
WORKDIR /go/src/github.com/hossainemruz/linked-issues
# Copy the source code
COPY . .

# Add certificates
RUN set -x \
    && apk update \
    && apk add ca-certificates

# Build static binary
RUN CGO_ENABLED=0 go install -installsuffix "static" .

# Now, build the final image on scratch base
FROM scratch
COPY --from=0 /go/bin/linked-issues /linked-issues
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT [ "/linked-issues" ]
