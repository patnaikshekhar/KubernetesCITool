FROM golang:1.12-alpine3.9 as build-env
ENV GO111MODULE=on
ENV CGO_ENABLED=0

RUN apk update && \
    apk add --no-cache make ca-certificates git && \
    update-ca-certificates

WORKDIR $GOPATH/src/github.com/patnaikshekhar/kubernetescitool

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make build-cp && mv ./kci-server /kci-server

FROM scratch

COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /kci-server /kci-server

ENTRYPOINT [ "/kci-server" ]