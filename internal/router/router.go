package router

import (
	"github.com/GuaiZai233/Larvar/internal/auth"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	public := r.Group("/api")
	{
		public.POST("/login", auth.Login)
		public.POST("/reg", auth.Register)
	}
}
