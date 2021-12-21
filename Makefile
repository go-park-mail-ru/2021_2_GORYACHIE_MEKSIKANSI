.PHONY: test
test:
	go test -count=1 -coverprofile=cover_wo_mock -v -race -timeout 300s ./... -coverpkg=./... ./...
	grep -F -v "mocks" cover_wo_mock | grep -F -v "proto"  | grep -F -v "easyjson" > cover
	go tool cover -func cover > result
	go tool cover -func cover
	rm cover_wo_mock

.PHONY: save_db
save_db:
	pg_dump hot_mexicans_db -f ./build/postgresql/fill --inserts
	grep -v "\-\-['\n'|' ']" ./build/postgresql/fill > ./build/postgresql/fill.sql
	rm ./build/postgresql/fill

