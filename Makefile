mock:
	mockgen -source internal/contract/database.go -package main_mock  -destination internal/db/mock/main_mock/main_mock.go database/sql/driver Conn,Driver