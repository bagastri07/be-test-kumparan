run:
	go run server.go
db-new:
	dbmate new $(name)
db-migrate:
	dbmate -e DB_URL migrate
db-rollback:
	dbmate -e DB_URL rollback