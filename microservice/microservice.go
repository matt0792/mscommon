package microservice

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type Microservice struct {
	container   *dig.Container
	providers   []interface{}
	engine      *gin.Engine
	middlewares []gin.HandlerFunc
}

func NewMicroservice() *Microservice {
	return &Microservice{
		container: dig.New(),
		engine:    gin.New(),
	}
}

func (ms *Microservice) AddLayer(constructor interface{}) *Microservice {
	ms.providers = append(ms.providers, constructor)
	return ms
}

func (ms *Microservice) Use(mw gin.HandlerFunc) *Microservice {
	ms.middlewares = append(ms.middlewares, mw)
	return ms
}

func (ms *Microservice) Build() error {
	for _, p := range ms.providers {
		if err := ms.container.Provide(p); err != nil {
			return err
		}
	}

	ms.engine.Use(gin.Recovery())
	for _, mw := range ms.middlewares {
		ms.engine.Use(mw)
	}

	// Routes
	err := ms.container.Invoke(func(registrars []RouteRegistrar) {
		for _, r := range registrars {
			r.RegisterRoutes(ms.engine)
		}
	})
	if err != nil {
		return err
	}

	return nil
}

func (ms *Microservice) Run(addr string) {
	if err := ms.Build(); err != nil {
		log.Fatalf("failed to build microservice: %v", err)
	}
	srv := &http.Server{
		Addr:    addr,
		Handler: ms.engine,
	}
	log.Printf("Starting service on %s", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
