package main

import (
	"github.com/kataras/iris/v12"
	irisInfra "go-clean-arch-ddd/account/infra/in/iris"
)

/*
	We are using google wire to inject all dependencies
*/

func main() {
	application := iris.Default()

	usecase := newAccountUsecase()
	irisInfra.BindUseCasesWithWire(application, usecase)

	_ = application.Run(iris.Addr(":8080"))
}
