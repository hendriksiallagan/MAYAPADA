package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"

	// news
	newsHttpDeliver "github.com/mayapada/news/delivery/http"
	newsRepo "github.com/mayapada/news/repository"
	newsUcase "github.com/mayapada/news/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/mayapada/middleware"
)

func init() {
	// viper.SetConfigFile(`config.json`)
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

	// dbHost := viper.GetString(`database.host`)
	// dbPort := viper.GetString(`database.port`)
	// dbUser := viper.GetString(`database.user`)
	// dbPass := viper.GetString(`database.pass`)
	// dbName := viper.GetString(`database.name`)
	dbHost := viper.GetString(`MYSQL_HOST`)
	dbPort := viper.GetString(`MYSQL_PORT`)
	dbUser := viper.GetString(`MYSQL_USER`)
	dbPass := viper.GetString(`MYSQL_PASSWORD`)
	dbName := viper.GetString(`MYSQL_DBNAME`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	dsn := fmt.Sprintf("%s?parseTime=true", connection)
	dbConn, err := sql.Open(`mysql`, dsn)
	// if err != nil && viper.GetBool("debug") {
	if err != nil && viper.GetBool(`DEBUG`) {
		fmt.Println(err)
	}

	err = dbConn.Ping()
	if err != nil {
		os.Exit(1)
	}
	defer dbConn.Close()

	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	// timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Hour
	timeoutContext := time.Duration(viper.GetInt(`CONTEXT_TIMEOUT`)) * time.Hour

	// repos
	newsRepo := newsRepo.NewNewsRepository(dbConn)

	// news
	newsUcase := newsUcase.NewNewsUsecase(newsRepo, timeoutContext)
	newsHttpDeliver.NewNewsHttpHandler(e, newsUcase)

	// e.Start(viper.GetString("server.address"))
	e.Start(viper.GetString(`PORT`))
}
