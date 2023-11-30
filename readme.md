# grpc authentication and  autharization service

protoc
    -I                                  working dir
    empty arg                           input file path
    --go_out                            output files path
    --go_opt=paths=source_relative      output files will use the save package as the input file
    -go-grpc-out                        grpc output files path

protoc -I protos/proto protos/proto/sso/sso.proto --go_out=./protos/gen/go --go_opt=paths=source_relative  --go-grpc_out=./protos/gen/go --go-grpc_opt=paths=source_relative 