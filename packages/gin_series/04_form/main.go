package main

import (
	"go-index/packages/gin_series/04_form/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
