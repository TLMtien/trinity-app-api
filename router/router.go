package router

import (
	"trinity_app/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	Controller  controllers.Controllers
	Router      *gin.Engine
	RouterGroup *gin.RouterGroup
}

func (r *Router) Init() {
	router := gin.Default()
	r.Router = router

	r.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.RouterGroup = router.Group("/api")

	r.RouterGroup.GET("/voucher", r.Controller.CheckValidVoucher)
	r.RouterGroup.POST("/voucher", r.Controller.GenerateVoucher)
	r.RouterGroup.GET("/campaign/:id", r.Controller.CheckEligibilityCampaign)
	r.RouterGroup.POST("/purchase", r.Controller.Purchase)
}

func (r *Router) Start(address string) error {
	r.Init()
	return r.Router.Run(address)
}

func NewRouter(ctl controllers.Controllers) *Router {
	return &Router{Controller: ctl}
}
