package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"gowild.com/internal"
	"gowild.com/internal/config"
	"gowild.com/internal/handlers"
	"gowild.com/internal/repositories"
	"gowild.com/internal/services"
	database "gowild.com/pkg/db"
)

type appHandlers struct {
	forestHandler   internal.ForestHandler
	mountainHandler internal.MountainHandler
}

func attachHandlers(forestHandler internal.ForestHandler, mountainHandler internal.MountainHandler, mux *http.ServeMux) {

	mux.HandleFunc("v1/forest/animals", forestHandler.Animals)
	//mux.HandleFunc("v1/mountain/animals", nil)
	//mux.HandleFunc("v1/mountain/plants", nil)
}

// This is our "factory" - this is where we set the stage. All repository share the same db pool.
func bootstrap(db database.DB) appHandlers {

	// Mountains
	//mountainAnimalRepository := repositories.NewAnimalRepository(db.Pool)
	//mountainPlantRepository := repositories.NewPlantRepository(db.Pool)
	//mountainService := services.NewMountainService(&mountainAnimalRepository, &mountainPlantRepository)
	//mountainHandler := handlers.NewMountainHandler(&mountainService)

	// Forests
	forestAnimalRepository := repositories.NewAnimalRepository(db.Pool)
	forestPlantRepository := repositories.NewPlantRepository(db.Pool)
	forestService := services.NewForestService(&forestAnimalRepository, &forestPlantRepository)
	forestHandler := handlers.NewForestHandler(&forestService)

	return appHandlers{
		//mountainHandler: &mountainHandler,
		forestHandler: &forestHandler,
	}
}

func main() {

	fmt.Println("gowild starting")

	// Get configuration from environment
	cfg := config.ReadAppConfig()
	fmt.Println("read app config and proceeding with environment setting " + string(cfg.AppEnvironment))

	// Connect to external resources based on config
	db := database.ConnectToPostgresDB(cfg.DBConfig)

	// Create our long living structs
	appServices := bootstrap(db)
	mux := http.NewServeMux()

	// Register handlers
	attachHandlers(appServices.forestHandler, appServices.mountainHandler, mux)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("captured %v, exiting..", sig)
			os.Exit(1)
		}
	}()

	fmt.Println("server starting on port " + cfg.AppPort)
	address := fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)
	// Keep running
	if err := http.ListenAndServe(address, mux); err != nil {
		fmt.Println(err)
	}
}
