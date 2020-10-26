package router

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/raziel2244/geckosite/router/handlers"
)

var (
	// Router stores the initialised router object.
	Router *mux.Router
	once   sync.Once
)

// Init creates the router object and sets up routes.
func Init() *mux.Router {
	once.Do(func() {
		Router = mux.NewRouter()

		Router.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

		fs := http.FileServer(http.Dir("static"))

		s3Handler := http.StripPrefix("/s3/", http.HandlerFunc(handlers.S3))
		Router.PathPrefix("/s3/{bucket}").Handler(s3Handler)

		Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
		Router.Path("/favicon.ico").Handler(fs)

		Router.Path("/").HandlerFunc(handlers.Home)
		Router.Path("/about").HandlerFunc(handlers.About)
		Router.Path("/contact").HandlerFunc(handlers.Contact)
		Router.Path("/blog").HandlerFunc(handlers.Construction)

		/// use {order} for geckos / snakes
		Router.Path("/{order}").HandlerFunc(handlers.Cards)
		Router.Path("/{order}/{type}").HandlerFunc(handlers.Cards)
		Router.Path("/{order}/{type}/{id}").HandlerFunc(handlers.Animals)
	})

	return Router
}

// Serve starts a http server using the router.
func Serve() {
	log.Println("Starting server on http://localhost")
	if err := http.ListenAndServe(":80", Router); err != nil {
		log.Panic("failed to start server from router", err)
	}
}

// InitAndServe initialises the router and starts the http server.
func InitAndServe() {
	Init()
	Serve()
}
