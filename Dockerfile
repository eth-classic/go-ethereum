# Build Geth in a stock Go builder container
FROM golang:1.12-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /go-ethereum

# Set up github credential for pulling private go module dependencies (remove section once public repo)
ARG github_token

RUN git config --global url."https://$github_token:x-oauth-basic@github.com/".insteadOf "https://github.com/"
# End credential configuration section

RUN cd /go-ethereum && make cmd/geth
# RUN cd /go-ethereum && go build -o ./bin/geth -tags="netgo" ./cmd/geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-ethereum/bin/geth /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["geth"]