FROM golang:alpine AS build-env
LABEL maintainer "Alexander Makhinov <monaxgt@gmail.com>" \
      repository="https://github.com/MonaxGT/gosddl"

COPY . /go/src/github.com/MonaxGT/gosddl

RUN apk add --no-cache git mercurial \
    && cd /go/src/github.com/MonaxGT/gosddl/service/gosddl \
    && go get -t . \
    && CGO_ENABLED=0 go build -ldflags="-s -w" \
                              -a \
                              -installsuffix static \
                              -o /gosddl
RUN adduser -D app

FROM scratch

COPY --from=build-env /gosddl /app/gosddl
COPY --from=build-env /etc/passwd /etc/passwd

USER app

WORKDIR /app

ENTRYPOINT ["./gosddl"]