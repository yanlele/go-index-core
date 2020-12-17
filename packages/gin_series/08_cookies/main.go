package main

import (
	"go-index/packages/gin_series/08_cookies/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
