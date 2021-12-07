FROM golang:1.17 as build
WORKDIR /go/src/app
COPY ./library .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

FROM alpine:latest
WORKDIR /root/
COPY --from=build /go/src/app/link.yaml ./
COPY --from=build /go/src/app/app ./
ENTRYPOINT ["./app"]