FROM golang:alpine as build-env

ARG PACKAGE_NAME=es-curator

COPY ./vendor /go/src/
COPY ./main.go /go/src/${PACKAGE_NAME}/
RUN go build -o /go/src/${PACKAGE_NAME}/es-curator /go/src/${PACKAGE_NAME}/main.go

FROM alpine
RUN apk add --update --no-cache ca-certificates
COPY --from=build-env /go/src/${PACKAGE_NAME}/es-curator /usr/local/bin
COPY ./curator-config.yaml /usr/local/bin
COPY ./clean-up-indices.yaml /usr/local/bin
RUN apk upgrade --no-cache \
  && apk add --no-cache \
    python \
    py-pip \
  && adduser -D -g '' curator \
  && pip install --upgrade pip \
  && pip install elasticsearch-curator

ENTRYPOINT [ "es-curator" ]