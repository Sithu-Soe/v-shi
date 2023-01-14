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
