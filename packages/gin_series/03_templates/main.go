package main

import (
	"go-index/packages/gin_series/03_templates/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
