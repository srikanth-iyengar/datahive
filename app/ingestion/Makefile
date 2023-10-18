.PHONY: proto

PROTO_PKG="io/datahive/ingestion/proto"

proto-compile:
	@ mvn compile

proto: proto-compile
	@ echo "Moving the proto files from target/ to src/" && mkdir -p ./src/main/java/${PROTO_PKG} && \
	  cp ./target/generated-sources/protobuf/java/${PROTO_PKG}/* ./src/main/java/${PROTO_PKG} && \
	  cp ./target/generated-sources/protobuf/grpc-java/${PROTO_PKG}/* ./src/main/java/${PROTO_PKG}

proto-clean:
	@ rm -rf ./src/main/java/io/datahive/ingestion/proto
