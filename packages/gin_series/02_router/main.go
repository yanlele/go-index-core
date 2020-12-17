package main

import "go-index/packages/gin_series/02_router/initRouter"

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
