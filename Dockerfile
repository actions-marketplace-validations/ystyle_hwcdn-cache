FROM golang:alpine AS build-env
ADD . /code
WORKDIR /code
RUN go build -v -o hwcdn main.go

FROM alpine
COPY --from=build-env /code/hwcdn /bin/hwcdn
ENTRYPOINT /bin/hwcdn