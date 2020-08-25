package router

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
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

		fs := http.FileServer(http.Dir("static"))
		Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
		Router.Path("/favicon.ico").Handler(fs)

		Router.Path("/").HandlerFunc(home)
		Router.Path("/about").HandlerFunc(about)
		Router.Path("/contact").HandlerFunc(contact)

		/// use {order} for geckos / snakes
		Router.Path("/geckos").HandlerFunc(cards)
		Router.Path("/geckos/{type}").HandlerFunc(cards)
		Router.Path("/geckos/{type}/pets").HandlerFunc(animals)
		Router.Path("/geckos/{type}/breeders").HandlerFunc(animals)
		Router.Path("/geckos/{type}/for-sale").HandlerFunc(animals)
		Router.Path("/geckos/{type}/{id}").HandlerFunc(animal)
	})

	return Router
}

// Serve starts a http server using the router.
func Serve() {
	if err := http.ListenAndServe(":8080", Router); err != nil {
		log.Panic("failed to start server from router", err)
	}
}

// InitAndServe initialises the router and starts the http server.
func InitAndServe() {
	Init()
	Serve()
}
