FROM golang:latest as builder
WORKDIR /go/src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o vecosy-server

FROM alpine:latest
RUN mkdir /
RUN mkdir /config
EXPOSE 8080
EXPOSE 8081
COPY --from=builder /go/src/vecosy-server /vecosy-server
CMD /vecosy-server --config /config/vecosy.yml