FROM golang:1.22.6-alpine

WORKDIR /app
COPY . .

RUN apk update \
  && apk add --no-cache \
  git \
  make \
  && rm -rf /var/cache/apt/* \
  && make setup


CMD ["task", "hotload"]
