package main

import (
	"fmt"
	"golang-testing/domain"
	"golang-testing/internal/config"
	"golang-testing/internal/delivery/http"
	"golang-testing/internal/delivery/http/middleware"
	"golang-testing/internal/repository/mysql"
	"golang-testing/internal/usecase/project"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`.env`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {

	db, _ := config.DBConnection()

	db.AutoMigrate(domain.Project{})

	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	projectRepository := mysql.NewMysqlProjectRepository(db)

	timeoutContext := time.Duration(viper.GetInt("APP_TIMEOUT")) * time.Second
	projectUsecase := project.NewProjectService(projectRepository, timeoutContext)

	http.NewProjectHandler(e, projectUsecase)
	log.Fatal(e.Start(":" + viper.GetString("APP_PORT")))
}
