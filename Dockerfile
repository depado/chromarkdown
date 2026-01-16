# Build Step
FROM golang:1.25.6-alpine@sha256:bc2596742c7a01aa8c520a075515c7fee21024b05bfaa18bd674fe82c100a05d as builder

# Source
WORKDIR $GOPATH/src/github.com/depado/capybara
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify
COPY . .

# Build
RUN go build -o /tmp/chromarkdown

# Final Step
FROM gcr.io/distroless/static@sha256:cd64bec9cec257044ce3a8dd3620cf83b387920100332f2b041f19c4d2febf93
COPY --from=builder /tmp/chromarkdown /go/bin/chromarkdown
ENTRYPOINT ["/go/bin/chromarkdown"]
