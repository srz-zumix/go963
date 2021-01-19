FROM golang:1.15-alpine3.12 as builder
RUN apk add --no-cache make
WORKDIR /go/src/go963
COPY . .
RUN CGO_ENABLED=0 go build -v -o ./dist/go963 -tags=prod

FROM alpine:3.12
# https://circleci.com/docs/2.0/custom-images/#required-tools-for-primary-containers
RUN apk add --no-cache git openssh tar gzip ca-certificates
COPY --from=builder /go/src/go963/dist/* /usr/bin/
ENTRYPOINT ["go963"]
CMD ["help"]
