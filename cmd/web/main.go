package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/williamneokh/booking/pkg/config"
	"github.com/williamneokh/booking/pkg/handlers"
	"github.com/williamneokh/booking/pkg/render"
	"log"
	"net/http"
	"time"
)

var portNum2 = ":8080"
var app config.Appconfig
var session *scs.SessionManager

func main() {

	// Change this to true when inn production
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc := render.TemplateSet()
	app.TemplateCache = tc
	app.UserCache = false
	render.NewTemplate(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	fmt.Println("Server is running on http://localhost" + portNum2)

	srv := &http.Server{
		Addr:    portNum2,
		Handler: routes(&app),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
