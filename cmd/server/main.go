package main

import (
	_"fmt"
	"net/http"
	"unbeatable-abayomi/go-rest-api/internal/comment"
	transportHTTP "unbeatable-abayomi/go-rest-api/internal/transport/http"
	"unbeatable-abayomi/go-rest-api/internal/transport/http/database"
	log "github.com/sirupsen/logrus"
)

//App - the struct which contains things like pointers
// to database connections
//App - contains application information
type App struct{
	Name string
	Version string

}
//Run - sets up our application
func (app *App) Run() error{
	//fmt.Println("Setting up our App")
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName" : app.Name,
			"AppVersion" : app.Version,
		}).Info("Setting Up Application ")

	

	var err error
	db, err := database.NewDatabase()
	if err != nil {
  return err
	}

	err = database.MigrateDB(db)
	if err != nil{
		return err
	}
	commentService := comment.NewService(db)
	handler  := transportHTTP.NewHandler(commentService)

	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil{
		//fmt.Println("Failed to set up server")
		log.Error("Failed to set up server");
		return err
	}
	return nil
}

func main() {
	//fmt.Println("GO REST API Course")
	log.Info("GO REST API Course")
	app := App{
		Name: "Commenting Service",
		Version: "1.0.0",
	}

	if err := app.Run(); err != nil{
		//fmt.Println("Error starting up our REST API")
		log.Error("Error starting up our REST API")
		//fmt.Println(err)
		log.Fatal(err)
	}
}