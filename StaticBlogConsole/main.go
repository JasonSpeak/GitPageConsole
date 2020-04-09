package main

import "sbc/server"

func main() {
	r := server.NewRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
