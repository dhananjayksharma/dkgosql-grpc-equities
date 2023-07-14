package handlers

import (
	"grpcequitiesapi/internals/middleware"
	"grpcequitiesapi/pkg/v1/models/merchants"
	"grpcequitiesapi/pkg/v1/models/orderprocessed"
	"grpcequitiesapi/pkg/v1/models/users"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	CONST_MERCHANTS      = "/merchants"
	CONST_ORDERPROCESSED = "/orderprocessed"
)

func SetupRouter(merchantService merchants.MerchantService, userService users.UserService, orderService orderprocessed.OrderProcessedService) *gin.Engine {
	r := gin.Default()
	corsConfig := CORS()

	r.Use(corsConfig)
	healthHandler := NewHealthHandler()

	// swagger:route GET /health checking for the server
	r.GET("/health", healthHandler.Health)

	// NewMerchantHandler
	merchantHandler := NewMerchantHandler(merchantService)
	// NewUserHandler
	userHandler := NewUserHandler(userService)

	orderProcessedHandler := NewOrderProcessedHandler(orderService)
	{
		v1Group := r.Group(CONST_MERCHANTS)
		{
			// swagger:route Group /secured route group for application
			secured := v1Group.Group("/secured").Use(middleware.Auth())
			{
				secured.PUT("/merchant/:code", merchantHandler.UpdateMerchantByID)
				secured.POST(CONST_MERCHANTS, merchantHandler.CreateMerchant)
				secured.GET(CONST_MERCHANTS, merchantHandler.GetMerchantList) //.Use(auth.GetClaim(c))
				secured.GET("/members/:code", userHandler.ListMembersByCode)
			}
			v1Group.POST("/:code/member", userHandler.CreateMerchantMember)
			v1Group.GET(CONST_MERCHANTS, merchantHandler.GetMerchantList)
			v1Group.POST(CONST_MERCHANTS, merchantHandler.CreateMerchant)
			v1Group.POST("/member/login", userHandler.LoginMember)
			v1Group.GET("/member/refresh", userHandler.RefreshToken)

		}
		// orderprocessed
		v2Group := r.Group(CONST_ORDERPROCESSED)
		v2Group.GET(CONST_ORDERPROCESSED, orderProcessedHandler.GetOrderProcessedList)
		v2Group.GET("/listbyid/:userid", orderProcessedHandler.ListOrderProcessedByID)
		v2Group.POST(CONST_ORDERPROCESSED, orderProcessedHandler.CreateOrderProcessed)
		v2Group.PUT(CONST_ORDERPROCESSED, orderProcessedHandler.UpdateOrderProcessedByID)
		v2Group.POST(CONST_ORDERPROCESSED+"/bulk/:userid", orderProcessedHandler.BulkOrderProcessedByUserId)
	}
	return r
}

func CORS() gin.HandlerFunc {
	config := cors.Config{}
	config.AllowHeaders = []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"}
	// config.AllowAllOrigins = true
	config.AllowBrowserExtensions = true
	config.AllowCredentials = true
	config.AllowWildcard = true
	config.AllowOrigins = []string{"*"}
	config.MaxAge = time.Hour * 12
	return cors.New(config)
}
