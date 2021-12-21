.PHONY: test
test:
	go test -count=1 -coverprofile=cover_wo_mock -v -race -timeout 300s ./... -coverpkg=./... ./...
	grep -F -v "mocks" cover_wo_mock | grep -F -v "proto"  | grep -F -v "easyjson" > cover
	go tool cover -func cover > result
	go tool cover -func cover
	rm cover_wo_mock

fill:
	grep -v "\-\-['\n'|' ']" ./build/postgresql/fill > ./build/postgresql/fill.sql
	rm fill

save_db: fill
	pg_dump hot_mexicans_db -f fill --inserts

