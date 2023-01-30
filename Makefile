gen-gql:
	cd cmd/back && go run github.com/99designs/gqlgen generate

gen-proto:
	protoc --proto_path=pkg/proto/media/v1 --go_out=pkg/pb/media/v1 --go_opt=paths=source_relative --go-grpc_out=pkg/pb/media/v1 --go-grpc_opt=paths=source_relative pkg/proto/media/v1/*.proto
