FROM golang:1.15.3 as builder

ENV GO111MODULE=on
ARG ARCH=amd64

ADD / /go/src/github.com/stavangler/venue
WORKDIR /go/src/github.com/stavangler/venue

RUN go get -d -v ./...

RUN go install -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o main .

FROM scratch
COPY --from=builder /go/src/github.com/stavangler/venue/main /app/

WORKDIR /app
CMD ["./main"]
EXPOSE 8080
