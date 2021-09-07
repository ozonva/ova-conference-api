FROM golang:1.17-alpine  AS builder

RUN apk add --update make

WORKDIR /ova-conference-api/

COPY . /ova-conference-api/

RUN make deps && make build

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /ova-conference-api/bin/ova-conference-api .

RUN chown root:root ova-conference-api

EXPOSE 8080
EXPOSE 9090
CMD ["./ova-conference-api"]