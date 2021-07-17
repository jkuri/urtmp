# stage 1 - user interface
FROM node:14-alpine as ui

COPY ./web/urtmp ./app/ui

WORKDIR /app/ui

RUN npm i npm -g && npm install && npm run build

# stage 2 - build
FROM golang:1.16-alpine as build

WORKDIR /app

RUN apk --no-cache add make git ca-certificates alpine-sdk

COPY . /app/

COPY --from=ui /app/ui/dist /app/web/urtmp/dist

RUN go install github.com/jkuri/statik/...@latest

RUN make

# stage 3 - image
FROM scratch

LABEL maintainer="Jan Kuri <jkuri88@gmail.com>" \
  org.label-schema.schema-version="1.0" \
  org.label-schema.name="urtmp" \
  org.label-schema.description="uRTMP is a simple RTMP server with convenent UI to watch live streams" \
  org.label-schema.vcs-url="https://github.com/jkuri/urtmp"

COPY --from=build /app/build/urtmp /usr/bin/urtmp

ENTRYPOINT [ "/usr/bin/urtmp" ]

EXPOSE 1935 8080
