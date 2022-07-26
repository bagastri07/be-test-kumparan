run:
	go run server.go
db-new:
	dbmate new $(name)
db-migrate:
	dbmate -e DB_URL migrate
db-rollback:
	dbmate -e DB_URL rollback
mock:
	mockery --dir=./services --case=underscore --all --disable-version-string
test:
	go test -v -cover -coverprofile=cover.out `go list ./... | grep -v /mock | grep -v /constants | grep -v /qtest`
cover:
	go test -cover -coverprofile=cover.out `go list ./... | grep -v /mock | grep -v /constants | grep -v /qtest` && go tool cover -html=cover.out