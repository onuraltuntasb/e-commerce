package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/onuraltuntasb/e-commerce/internal/auth"
	"github.com/onuraltuntasb/e-commerce/internal/driver"
	"github.com/onuraltuntasb/e-commerce/internal/handlers"
	"github.com/onuraltuntasb/e-commerce/internal/routes"
)

func main() {

	mode := flag.String("mode", "", "-mode=dev")

	flag.Parse()

	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	envDir := currDir + "/internal/env/"

	if *mode == "dev" {
		if err := godotenv.Load(envDir + ".env.dev"); err != nil {
			log.Fatal("Error loading .env file")
		}
	} else if *mode == "prod" {
		if err := godotenv.Load(envDir + ".env.prod"); err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		fmt.Println("No mode or invalid option has been given, app is starting as dev mode!")
		if err := godotenv.Load(envDir + ".env.dev"); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	dsn := os.Getenv("DSN")

	conn, err := driver.ConnectSQL(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.SQL.Close()

	log.Println("Connected to database!")

	//repository pattern to use repo operations in handlers (reciving pattern)
	repo := handlers.NewRepo(conn)
	handlers.NewHandlers(repo)

	//auth
	auth.NewAuth(&auth.AuthType{
		Issuer:                 os.Getenv("JWT_ISSUER"),
		Audience:               os.Getenv("JWT_AUDIENCE"),
		Secret:                 os.Getenv("JWT_SECRET"),
		TokenExpiry:            time.Minute * 15,
		RefreshExpiry:          time.Hour * 24,
		CookiePath:             "/",
		CookieAccessTokenName:  os.Getenv("JWT_ACCESS_TOKEN_COOKIE_NAME"),
		CookieRefreshTokenName: os.Getenv("JWT_REFRESH_TOKEN_COOKIE_NAME"),
		//__host asks setting from https (secure connection)
		//CookieName:    "__Host-refresh_token",
		CookieDomain: os.Getenv("JWT_COOKIE_DOMAIN"),
	})

	portNumber := os.Getenv("PORT")

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes.Routes(),
	}
	

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
