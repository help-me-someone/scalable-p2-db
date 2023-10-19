// The database service is responsible for the CRUD operations.

// TODO: Environment variables!

package main

import (
	"log"
	"net/http"

	"github.com/help-me-someone/scalable-p2-db/handlers"
	"github.com/help-me-someone/scalable-p2-db/models/user"
	"github.com/help-me-someone/scalable-p2-db/models/video"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:password@tcp(mysql:3306)/toktik-db?charset=utf8mb4&parseTime=True&loc=Local"
	toktik_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	} else {
		// We only initialize the table if it doesn't already exist.
		// This should not really impact performance since it only checks
		// during start up.
		if !toktik_db.Migrator().HasTable(&user.User{}) && !toktik_db.Migrator().HasTable(&video.Video{}) {
			InitTables(toktik_db)
			log.Println("Database initialized!")
		}
	}

	cmw := handlers.ConnectionMiddleware{
		Connection: toktik_db,
	}

	router := httprouter.New()

	router.GET("/user/:username", cmw.Attach(handlers.GetUser))
	router.GET("/user/:username/videos", cmw.Attach(handlers.GetUserVideos))
	router.GET("/user/:username/videos/:video", cmw.Attach(handlers.GetUserVideo))

	// Returns the most popular videos.
	router.GET("/popular/:amount/:page", cmw.Attach(handlers.GetTopPopularVideos))

	router.GET("/video/:key", cmw.Attach(handlers.GetVideo))

	// Expects: "username", "hashed_password" in JSON.
	router.POST("/user", cmw.Attach(handlers.CreateUser))

	// Expects: "key", "name", "owner_name" in JSON.
	router.POST("/video/", cmw.Attach(handlers.CreateVideo))

	// Add CORS support (Cross Origin Resource Sharing)
	handler := cors.Default().Handler(router)
	err = http.ListenAndServe(":8083", handler)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Serving on port 8083")
	log.Fatal(http.ListenAndServe(":8083", handler))
}
