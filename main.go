package main

import (
	"log"
	"startupfunding/handler"
	"startupfunding/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
|--------------------------------------------------------------------------
| Concept
|--------------------------------------------------------------------------
|
| input
| handler -> mapping input
| services -> mapping input to struct
| repository -> mapping to struct user
| db
|
*/

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/go_startupfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/world/email-checkers", userHandler.EmailAvailability)
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/user/fetch", userHandler.User)
	api.POST("user/upload-avatar", userHandler.UploadAvatar)

	router.Run()
}

// func handler(c *gin.Context) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/go_startupfunding?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	fmt.Println("Connecting to database is successful!")

// 	var users []user.User

// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)

// }
