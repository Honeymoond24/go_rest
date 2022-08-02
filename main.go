package main

import (
	// "log"

	user "go_rest/handlers/user_handler.go"

	// "github.com/gin-gonic/gin"
)

func main() {
    fmt.Println("hello world")
    fmt.Println(user.a)
}
// func main() {
// 	r := gin.Default()
// 	r.POST("/login", auth.Login)
// 	r.POST("/register", auth.Register)
// 	log.Fatal(r.Run(":8080"))
// }