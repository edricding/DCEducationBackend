package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"backend/internal/config"
	"backend/internal/modules/universities"
)

func NewRouter(cfg config.Config, db *sqlx.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(CORSMiddleware(cfg.CORSAllowOrigins))

	// health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := r.Group("/api/v1")

	// universities module wiring
	uRepo := universities.NewRepo(db)
	uSvc := universities.NewService(uRepo)
	uHandler := universities.NewHandler(uSvc)

	ug := v1.Group("/universities")
	{
		ug.GET("", uHandler.Search)     // list search
		ug.GET("/:id", uHandler.GetByID) // detail
		ug.GET("/u-name-cn", uHandler.ListAllNameCN)
		ug.GET("/options-u-name-cn", uHandler.OptionsCN)


	}

	return r
}


