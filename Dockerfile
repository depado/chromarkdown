# Build Step
FROM golang:1.24.0-alpine@sha256:3d74d23af285af08b6a2c89a15c437b9bc2854f63948fb8fd703823528820230 as builder

# Source
WORKDIR $GOPATH/src/github.com/depado/capybara
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify
COPY . .

# Build
RUN go build -o /tmp/chromarkdown

# Final Step
FROM gcr.io/distroless/static@sha256:3f2b64ef97bd285e36132c684e6b2ae8f2723293d09aae046196cca64251acac
COPY --from=builder /tmp/chromarkdown /go/bin/chromarkdown
ENTRYPOINT ["/go/bin/chromarkdown"]
