package main

import (
	"go-index/packages/gin_series/01_hello_gin/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
