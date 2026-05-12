# Build / test
.PHONY: test vet generate verify
test:
	go test ./...
vet:
	go vet ./...
generate:
	bash scripts/sdk-generate-go.sh
verify:
	bash scripts/sdk-verify-go.sh
