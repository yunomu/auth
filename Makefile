.PHONY: build clean proto proto-clean check

ELM_DIR=console/src

PROTO_TARGETS=proto/api/api.pb.go
PROTO_ELM_TARGETS=$(ELM_DIR)/Proto/Api.elm

build: proto
	CGO_ENABLED=0 sam build --debug

proto: $(PROTO_TARGETS) $(PROTO_ELM_TARGETS)

check:
	go test ./...
	sam validate --lint

clean:
	rm -rf .aws-sam


proto/api/api.pb.go: proto/api.proto
	protoc --go_out=. $<

$(ELM_DIR)/Proto/Api.elm: proto/api.proto
	protoc --elm_out=$(ELM_DIR) $< 2> /dev/null
