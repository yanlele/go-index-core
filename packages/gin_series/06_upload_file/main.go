package main

import (
	"go-index/packages/gin_series/06_upload_file/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
