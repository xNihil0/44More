package main

import "mtdn.io/44More/routers"

func main() {
	router := routers.InitRouter()
	router.Run(":8080")
}
