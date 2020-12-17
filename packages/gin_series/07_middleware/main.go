package main

import (
	"go-index/packages/gin_series/07_middleware/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
