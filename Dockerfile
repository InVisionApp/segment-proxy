# This is the dockerfile we use to run the project locally as well as
# compile the code for a slim production image
FROM golang:1.10-alpine3.7 as builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/github.com/InVisionApp/segment-proxy

# Add rest of source code
COPY . ./

RUN CGO_ENABLED=$CGO_ENABLED \
  GOOS=$GOOS \
  GOARCH=$GOARCH \
  go build \
    -a -tags netgo -ldflags '-w -extldflags "-static"' \
    -o segment-proxy .

FROM alpine:3.7

RUN apk add --update --no-cache ca-certificates

COPY --from=builder /go/src/github.com/InVisionApp/segment-proxy/segment-proxy /segment-proxy
COPY --from=builder /go/src/github.com/InVisionApp/segment-proxy/static /static

ENV PORT 80
EXPOSE 80
WORKDIR /
ENTRYPOINT ["/segment-proxy"]
