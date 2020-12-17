package main

import "go-index/packages/gin_series/11_gorm/initRouter"

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
