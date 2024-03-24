FROM golang:1.22-alpine as buildbase

RUN apk add build-base git

ARG CI_JOB_TOKEN

WORKDIR /go/src/github.com/rarimo/rarimo-core

ENV GO111MODULE="on"
ENV CGO_ENABLED=1
ENV GOOS="linux"

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go mod vendor

RUN go build -o /usr/local/bin/rarimo-core github.com/rarimo/rarimo-core/cmd/rarimo-cored



###

FROM alpine:3.9

RUN apk add --no-cache ca-certificates

COPY --from=buildbase /usr/local/bin/rarimo-core /usr/local/bin/rarimo-core

ENTRYPOINT ["rarimo-core"]
