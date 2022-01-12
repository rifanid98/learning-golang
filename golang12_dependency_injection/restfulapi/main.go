package main

import (
	_ "github.com/go-sql-driver/mysql"
	error2 "golang12_dependency_injection/restfulapi/common/error"
	"golang12_dependency_injection/restfulapi/di"
)

func main() {
	server := di.InitializedServer()

	err := server.ListenAndServe()
	error2.PanicIfError(err)
}
