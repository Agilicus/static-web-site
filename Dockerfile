FROM golang:1.16 as build
ENV GO111MODULE=on
ENV CGO_ENABLED=0

COPY . /src
WORKDIR /src

RUN go get github.com/rakyll/statik \
 && /go/bin/statik -src=assets \
 && CGO_ENABLED=0 go build \
 && adduser --disabled-password --gecos '' web

FROM scratch

COPY --from=build /etc/passwd /etc/group /etc/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/static-web-site /static-web-site

USER web

ENTRYPOINT ["/static-web-site"]
