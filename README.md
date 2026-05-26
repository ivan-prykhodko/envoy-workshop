# Envoy Mesh Workshop

## mTLS

### Product service certificate (SPIFFE-Compatible)

```bash
openssl genrsa -out ./docker/certs/product-spiffe.key 2048
openssl req -new -key ./docker/certs/product-spiffe.key -out ./docker/certs/product-spiffe.csr -config ./docker/certs/product-openssl.cnf
openssl x509 -req -in ./docker/certs/product-spiffe.csr -CA ./docker/certs/ca.crt -CAkey ./docker/certs/ca.key -CAcreateserial -out ./docker/certs/product-spiffe.crt -days 365 -sha256 -extensions v3_req -extfile ./docker/certs/product-openssl.cnf
```

### Currency exchange rate service certificate (SPIFFE-Compatible)

```bash
openssl genrsa -out ./docker/certs/currency-spiffe.key 2048
openssl req -new -key ./docker/certs/currency-spiffe.key -out ./docker/certs/currency-spiffe.csr -config ./docker/certs/currency-openssl.cnf
openssl x509 -req -in ./docker/certs/currency-spiffe.csr -CA ./docker/certs/ca.crt -CAkey ./docker/certs/ca.key -CAcreateserial -out ./docker/certs/currency-spiffe.crt -days 365 -sha256 -extensions v3_req -extfile ./docker/certs/currency-openssl.cnf
```

### Verify mTLS

```bash
openssl s_client -connect currency-envoy:9092 -cert /certs/product-spiffe.crt -key /certs/product-spiffe.key -CAfile /certs/ca.crt
```

## Protobuf

```bash
protoc -I/srv --go_opt=module=github.com/ivan-prykhodko/envoy-workshop --go_out=/srv --go-grpc_opt=module=github.com/ivan-prykhodko/envoy-workshop --go-grpc_out=/srv /srv/protobuf/*/*.proto
```
