FROM golang:1.15 AS builder
WORKDIR $GOPATH/src/github.com/QuoineFinancial/liquid-chain-explorer-api
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /liquid-chain-explorer-api .

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /liquid-chain-explorer-api ./
ENTRYPOINT ["./liquid-chain-explorer-api"]
