package main

import (
	"go-index/packages/gin_series/05_db/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
