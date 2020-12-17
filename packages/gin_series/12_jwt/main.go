package main

import "go-index/packages/gin_series/12_jwt/initRouter"

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
