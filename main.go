package main

import (
	_ "myapp/dbs"
	"myapp/router"
)

func main() {
	apiRouter := router.Router()
	apiRouter.Run(":9000")
}
