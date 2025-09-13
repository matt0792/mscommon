package microservice

import "github.com/gin-gonic/gin"

type RouteRegistrar interface {
	RegisterRoutes(*gin.Engine)
}
