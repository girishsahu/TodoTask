package main

import (
	"os"
	"time"
	"fmt"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"TodoTask/handler"
	"gopkg.in/mgo.v2"
	"github.com/tylerb/graceful"
)

type Configuration struct{
	MongoDBHosts string
	AuthDatabase string
	AuthUserName string
	AuthPassword string
}


func main() {

	configFile, _ := os.Open("conf.json")
	defer configFile.Close()
	decoder := json.NewDecoder(configFile)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
    fmt.Println()



    
	e := echo.New()

	file, _ := os.Create("Task.log")
	
	//e.Logger.SetLevel(log.ERROR)
	loggerConfig := middleware.DefaultLoggerConfig
	loggerConfig.Output = file
	e.Use(middleware.LoggerWithConfig(loggerConfig))

	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))

	mongoDBDialInfo := &mgo.DialInfo{
	Addrs:    []string{configuration.MongoDBHosts},
	Timeout:  60 * time.Second,
	Database: configuration.AuthDatabase,
	Username: configuration.AuthUserName,
	Password: configuration.AuthPassword,
}

	// Database connection
	//db, err := mgo.Dial("localhost:27017")
	db, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		 e.Logger.Fatal(err)

		
	}

	// Create indices
	if err = db.Copy().DB("TodoTask").C("users").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		 log.Fatal(err)

		
	}

	// Initialize handler
	h := &handler.Handler{DB: db}

	// Routes
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
	e.POST("/tasks", h.CreateTask)
	e.POST("/fetchTasks", h.FetchTasks)
	e.POST("/updateTask/:id",h.UpdateTask)
	e.POST("/completeTask/:id",h.CompleteTask)

    //Logger.SetOutput(io.Writer)
	// Start server

	e.Server.Addr = ":8000"

	// Serve it like a boss
	graceful.ListenAndServe(e.Server, 5*time.Second)
     
	 

	 //e.Logger.Fatal(e.Start(":8000"))
      
    
	
}
