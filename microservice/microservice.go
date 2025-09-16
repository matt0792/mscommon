package microservice

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type BackgroundWorker interface {
	Start()
}

type Microservice struct {
	Container   *dig.Container
	providers   []interface{}
	engine      *gin.Engine
	middlewares []gin.HandlerFunc
}

func NewMicroservice() *Microservice {
	return &Microservice{
		Container: dig.New(),
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

func (ms *Microservice) AddController(constructor interface{}) *Microservice {
	err := ms.Container.Provide(
		constructor,
		dig.Group("controllers"),
	)
	if err != nil {
		log.Fatalf("failed to provide controller: %v", err)
	}
	return ms
}

func (ms *Microservice) AddBackground(constructor interface{}) *Microservice {
	err := ms.Container.Provide(
		constructor,
		dig.Group("backgrounds"),
	)
	if err != nil {
		log.Fatalf("failed to provide background: %v", err)
	}
	return ms
}

func (ms *Microservice) startBackgrounds() {
	type BackgroundsIn struct {
		dig.In
		Backgrounds []BackgroundWorker `group:"backgrounds"`
	}

	err := ms.Container.Invoke(func(bi BackgroundsIn) {
		for _, bg := range bi.Backgrounds {
			go bg.Start()
		}
	})
	if err != nil {
		log.Printf("failed to start backgrounds: %v", err)
	}
}

func (ms *Microservice) Build() error {
	for _, p := range ms.providers {
		if err := ms.Container.Provide(p); err != nil {
			return err
		}
	}

	ms.engine.Use(gin.Recovery())
	for _, mw := range ms.middlewares {
		ms.engine.Use(mw)
	}

	type ControllersIn struct {
		dig.In
		Controllers []RouteRegistrar `group:"controllers"`
	}

	return ms.Container.Invoke(func(ci ControllersIn) {
		for _, r := range ci.Controllers {
			r.RegisterRoutes(ms.engine)
		}
	})
}

func (ms *Microservice) Run(addr string) {
	if err := ms.Build(); err != nil {
		log.Fatalf("failed to build microservice: %v", err)
	}

	ms.startBackgrounds()

	srv := &http.Server{
		Addr:    addr,
		Handler: ms.engine,
	}
	log.Printf("Starting service on %s", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
