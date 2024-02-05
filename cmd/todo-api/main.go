package main

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/BREACH1247/todo-api/internal/api"
	"github.com/BREACH1247/todo-api/internal/repository"
	"github.com/BREACH1247/todo-api/internal/routes"
	"github.com/BREACH1247/todo-api/internal/scylladb"
)

func main() {
	
	session, err := scylladb.SetupScyllaDB()
	if err != nil {
		log.Fatalf("Error setting up ScyllaDB: %v", err)
	}
	defer session.Close()

	repo := repository.NewTodoRepository(session)

	router := gin.Default()

	api := api.NewTodoAPI(repo)

	routes.SetupRoutes(router, api)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
