package main_mock

//go:generate mockgen -source ../../contract/database.go -package main_mock  -destination ./main_mock/main_mock.go database/sql/driver Conn,Driver
