# === Build Stage ===

FROM golang:1.19 as builder

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app ./cmd/my_app

# === Run Stage ===

FROM alpine:3.14

# enable this application to make secure network calls.
RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin/app /go/bin/app

ENTRYPOINT ["/go/bin/app"]
