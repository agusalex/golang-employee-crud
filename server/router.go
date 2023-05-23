package server

import (
	handlers2 "github.com/agusalex/golang-employee-crud/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	_ = router.SetTrustedProxies(nil)
	router.Use(gin.Logger())
	api := router.Group("/api")
	{
		api.GET("/ping", handlers2.PingHandler)
		api.GET("/health", handlers2.HealthHandler)
		v1 := api.Group("/v1")
		{
			v1.GET("/members", handlers2.GetMembersHandler)
			v1.GET("/members/:id", handlers2.GetMemberHandler)
			v1.PUT("/members/:id", handlers2.PutMemberHandler)
			v1.POST("/members", handlers2.PostMembersHandler)
			v1.DELETE("/members/:id", handlers2.DeleteMemberHandler)
			v1.GET("/search/members", handlers2.SearchMembersHandler)
			v1.GET("/tags", handlers2.GetTagsHandler)
		}
	}

	return router
}
