# Payment Gateway Golang

## Generate Private, Public Key Pair

generate key paris

- `openssl genpkey -algorithm RSA -out rsa_private.pem -pkeyopt rsa_keygen_bits:2048`

- `openssl rsa -in rsa_private.pem -pubout -out rsa_public.pem`

## TO generate graphql Code

```shell
go run github.com/99designs/gqlgen init
go run github.com/99designs/gqlgen generate
```

## MINIO

### With Local

```shell
minio server ~/minio --console-address :9090
```

### With Docker (FOR the first time)

```shell
mkdir -p ~/minio/data

docker run \
   -p 9000:9000 \
   -p 9090:9090 \
   -v ~/minio/data:/data \
   -e "MINIO_ROOT_USER=minioadmin" \
   -e "MINIO_ROOT_PASSWORD=minioadmin" \
   quay.io/minio/minio server /data --console-address ":9090"
```
