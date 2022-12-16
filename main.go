package main

import (
	"context"
	"fmt"
	"github.com/aryasadeghy/go-mic/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const PORT = ":8080"

func main() {
	fmt.Printf("Starting the application on http://localhost%s", PORT)
	prdLogg := log.New(os.Stdout, "product-api", log.LstdFlags)
	// router
	router := gin.Default()
	v1 := router.Group("/api/v1")

	pr := handlers.NewProduct(prdLogg)
	prRoute := v1.Group("/products")
	{
		prRoute.GET("/", pr.GetProducts)
		prRoute.POST("/", pr.AddProduct)
		prRoute.PUT("/:id", pr.UpdateProduct)
	}

	// server  configuration
	srv := &http.Server{
		Addr:         PORT,
		Handler:      router,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// graceful shutdown
	sc := make(chan os.Signal)
	signal.Notify(sc, os.Interrupt)
	signal.Notify(sc, os.Kill)

	sig := <-sc
	prdLogg.Println("Received terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(tc)
}
