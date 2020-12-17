package main

import "go-index/packages/gin_series/09_restful/initRouter"

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
