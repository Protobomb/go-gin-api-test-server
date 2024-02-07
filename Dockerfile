FROM golang:latest as builder
WORKDIR /go/src/app
COPY . .

# get go.mod dependencies
RUN go get -d -v ./...
RUN go install -v ./...

# run unit tests
RUN go test -v ./...

# build binary
RUN CGO_ENABLED=0 GOOS=linux  go build -ldflags="-w -s" -o /go/bin/main
RUN chmod +x /go/bin/main

# # Build smallest possible container
FROM scratch

COPY ./templates /templates
COPY --from=builder /go/bin/main /main

CMD ["/main"]
