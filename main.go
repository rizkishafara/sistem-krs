package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sistem-krs/config"
	"sistem-krs/router"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stnc/pongo4echo"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("couldn't load configuration file. Use container environment variable")
		log.Print(err)
		return
	} else {
		log.Print("success load configuration file")
	}
	config := config.LoadConfig(".")
	os.Setenv("SERVER_HOST", config.ServerHost)
	os.Setenv("SERVER_PORT", fmt.Sprintf("%v", config.ServerPort))
	os.Setenv("DB", config.Db)
	os.Setenv("DB_DSN", config.DbDsn)
	// os.Setenv("SMTP_SENDER", config.SmtpSender)
	// os.Setenv("SMTP_HOST", config.SmtpHost)
	// os.Setenv("SMTP_PORT", fmt.Sprintf("%v", config.SmtpPort))
	// os.Setenv("SMTP_USER", config.SmtpUser)
	// os.Setenv("SMTP_PASS", config.SmtpPass)
	os.Setenv("STARTLS", fmt.Sprintf("%v", config.StarTLS))
	// os.Setenv("HOST", "localhost")
	// os.Setenv("PORT", "1323")
	os.Setenv("SECRET_KEY", "secret")
	os.Setenv("GO_ENV", "development")
}

func main() {
	e := echo.New()
	pongo := pongo4echo.Renderer{}
	if os.Getenv("GO_ENV") != "production" {
		pongo.Debug = true
		e.Debug = true
	}

	e.Renderer = pongo

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} ${uri} - ${method} ${status} ${latency_human}\n",
		CustomTimeFormat: "2006/01/02 15:04:05.00000",
	}))
	cookie := sessions.NewFilesystemStore("", []byte(os.Getenv("SECRET_KEY")))
	// cookie.MaxLength(100000)
	storesess := &session.Config{
		Store: cookie,
	}

	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))))
	e.Use(session.MiddlewareWithConfig(*storesess))

	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:      "1; mode=block",
		ContentTypeNosniff: "nosniff",
		// XFrameOptions:      "ALLOW-FROM http://blumb.id",
	}))

	//public folder
	e.Static("/public", "public")

	router.Auth(e)

	host := fmt.Sprintf("%s:%v", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	fmt.Println("Server started on", host)
	srv := http.Server{
		Addr:    host,
		Handler: e,
	}

	log.Printf("server started on %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// e.Logger.Fatal(err)
		fmt.Println("Error main", err.Error())
	}
}
