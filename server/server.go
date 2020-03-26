package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	goHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	handlers "github.com/thatbeardo/go-play/handlers/products"
)

// StartServer sets up routes and begins the server
func StartServer(logger *log.Logger) {
	productHandler := handlers.NewProducts(logger)
	server := setupRoutes(productHandler)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel
	logger.Println("Recieved terminate, commencing graceful shutdown", sig)

	timeoutContext, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	cancelFunc()
	server.Shutdown(timeoutContext)
}

func setupRoutes(productHandler *handlers.Products) *http.Server {

	serveMux := mux.NewRouter()
	options := middleware.RedocOpts{SpecURL: "./api/swagger.yaml"}
	sh := middleware.Redoc(options, nil)

	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productHandler.GetProducts)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/api/swagger.json", http.FileServer(http.Dir("./")))

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.UpdateProducts)
	putRouter.Use(productHandler.ProductValidationMiddleware)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", productHandler.AddProducts)
	postRouter.Use(productHandler.ProductValidationMiddleware)

	deleteRouter := serveMux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.DeleteProduct)

	cors := goHandlers.CORS(
		goHandlers.AllowedOrigins([]string{"*"}),
		goHandlers.AllowedMethods(
			[]string{
				http.MethodGet,
				http.MethodPut,
				http.MethodPost,
				http.MethodDelete,
			}),
		goHandlers.AllowedHeaders([]string{"Content-Type"}),
	)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      cors(serveMux),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	return server
}
