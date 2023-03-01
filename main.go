package main

import (
	"sync"

	database "github.com/himanshu07070/newsletter/Database"
	utils "github.com/himanshu07070/newsletter/Utils"

	arc "github.com/himanshu07070/newsletter/ArcReactor"
	input "github.com/himanshu07070/newsletter/Input"
	routes "github.com/himanshu07070/newsletter/Routes"
	"github.com/joho/godotenv"
)

var (
	Username = ""
	Password = ""
	Host     = ""
	Port     = 0
	DBName   = ""
)

func init() {
	utils.InitializeLogger("Logs/log.log")
	database.InitializeMongoConnection(Username, Password, Host, Port, DBName)
	godotenv.Load(".env")
}

func startserver(data chan input.MetaData, wg *sync.WaitGroup) {
	utils.Logger.Info("Starting Server ....")
	defer wg.Done()
	r := routes.Router(data)
	err := r.Run("localhost:8000")
	if err != nil {
		utils.Logger.Error(err.Error())
		return
	}

}
func main() {
	data := make(chan input.MetaData)
	var wg sync.WaitGroup
	wg.Add(1)
	go arc.ReactorEngine(data, &wg)
	wg.Add(1)
	go startserver(data, &wg)
	wg.Wait()
}
