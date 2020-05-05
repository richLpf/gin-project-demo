package main

import "myapp/router"

func main() {
	apiRouter := router.Router()
	apiRouter.Run(":9000")
}
