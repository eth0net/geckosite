package router

import (
	"log"
	"net/http"
	"sync"

	"github.com/eth0net/geckosite/router/handlers"
	"github.com/eth0net/geckosite/static"
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

		Router.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

		fs := http.FileServer(http.FS(static.FS))

		s3Handler := http.StripPrefix("/s3/", http.HandlerFunc(handlers.S3))
		Router.PathPrefix("/s3/{bucket}").Handler(s3Handler)

		Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
		Router.Path("/favicon.ico").Handler(fs)
		Router.Path("/sitemap.xml").Handler(fs)

		Router.Path("/").HandlerFunc(handlers.Home)
		Router.Path("/about").HandlerFunc(handlers.About)
		Router.Path("/contact").HandlerFunc(handlers.Contact)
		Router.Path("/blog").HandlerFunc(handlers.Construction)

		/// use {order} for geckos / snakes
		Router.Path("/{order}/").HandlerFunc(handlers.Cards)
		Router.Path("/{order}/{type}/").HandlerFunc(handlers.Cards)
		Router.Path("/{order}/{type}/{status:[a-z-]+}/").HandlerFunc(handlers.Animals)
		Router.Path("/{order}/{type}/{id:[a-z0-9-]{36}}").HandlerFunc(handlers.Animal)
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
