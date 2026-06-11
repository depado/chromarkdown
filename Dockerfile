# Build Step
FROM golang:1.26.4-alpine@sha256:7a3e50096189ad57c9f9f865e7e4aa8585ed1585248513dc5cda498e2f41812c as builder

# Source
WORKDIR $GOPATH/src/github.com/depado/capybara
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify
COPY . .

# Build
RUN go build -o /tmp/chromarkdown

# Final Step
FROM gcr.io/distroless/static@sha256:3592aa8171c77482f62bbc4164e6a2d141c6122554ace66e5cc910cadb961ff0
COPY --from=builder /tmp/chromarkdown /go/bin/chromarkdown
ENTRYPOINT ["/go/bin/chromarkdown"]
