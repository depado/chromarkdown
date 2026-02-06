# Build Step
FROM golang:1.25.7-alpine@sha256:f6751d823c26342f9506c03797d2527668d095b0a15f1862cddb4d927a7a4ced as builder

# Source
WORKDIR $GOPATH/src/github.com/depado/capybara
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify
COPY . .

# Build
RUN go build -o /tmp/chromarkdown

# Final Step
FROM gcr.io/distroless/static@sha256:972618ca78034aaddc55864342014a96b85108c607372f7cbd0dbd1361f1d841
COPY --from=builder /tmp/chromarkdown /go/bin/chromarkdown
ENTRYPOINT ["/go/bin/chromarkdown"]
