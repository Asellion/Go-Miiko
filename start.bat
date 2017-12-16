@echo off
:x
go get -fix -u -v github.com/NatoBoram/Go-Miiko
cls
go run main.go
goto x