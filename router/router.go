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

		Router.NotFoundHandler = http.HandlerFunc(notFound)

		fs := http.FileServer(http.Dir("static"))
		Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
		Router.Path("/favicon.ico").Handler(fs)

		Router.Path("/").HandlerFunc(home)
		Router.Path("/about").HandlerFunc(about)
		Router.Path("/contact").HandlerFunc(contact)
		Router.Path("/blog").HandlerFunc(construction)

		/// use {order} for geckos / snakes
		Router.Path("/{order}").HandlerFunc(cards)
		Router.Path("/{order}/{type}").HandlerFunc(cards)
		Router.Path("/{order}/{type}/our-animals").HandlerFunc(ourAnimals)
		Router.Path("/{order}/{type}/holdbacks").HandlerFunc(holdbacks)
		Router.Path("/{order}/{type}/for-sale").HandlerFunc(forSale)
		Router.Path("/{order}/{type}/{id}").HandlerFunc(animal)
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
