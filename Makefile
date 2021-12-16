.PHONY: test
test:
	go test -count=1 -coverprofile=cover -v -race -timeout 300s ./... -coverpkg=./... ./...
	grep -F -v "mock" cover > cover_wo_mock
	go tool cover -func cover_wo_mock
