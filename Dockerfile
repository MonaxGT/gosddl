FROM golang:alpine AS build-env
LABEL maintainer "Alexander Makhinov <monaxgt@gmail.com>"

ENV CGO_ENABLED 0

COPY gosddl.go /go/src/gosddl/gosddl.go

WORKDIR /go/src/gosddl

RUN apk add --no-cache git mercurial \
    && go get github.com/gorilla/mux/... \
    && go build -o gosddl

FROM alpine:edge

RUN adduser -D app

COPY --from=build-env /go/src/gosddl/gosddl /app/gosddl

RUN chmod +x /app/gosddl \
  && chown -R app /app

USER app

WORKDIR /app

EXPOSE 8000

ENTRYPOINT ["/app/gosddl"]