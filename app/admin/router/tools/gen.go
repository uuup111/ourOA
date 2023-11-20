package tools

import (
	"github.com/gin-gonic/gin"
	"ruoyi-go/app/admin/api/tools"
	"ruoyi-go/utils"
)

func InitGen(e *gin.Engine) {
	v := e.Group("tool")
	{
		auth := v.Group("")
		auth.Use(utils.JWTAuthMiddleware())
		{
			auth.GET("/gen/list", tools.GenList)
			auth.GET("/gen/:tableId", tools.GenInfo)
			auth.GET("/gen/db/list", tools.GenDbList)
			auth.GET("/gen/column/:tableId", tools.GenColumnInfo)
			auth.POST("/gen/importTable", tools.ImportTable)
			auth.PUT("/gen", tools.GenEdit)
			auth.DELETE("/gen/:tableIds", tools.GenDelete)
			auth.GET("/gen/preview/:tableId", tools.PreviewGenTable)
			auth.GET("/gen/download/:tableName", tools.GenDownload)
			auth.GET("/gen/genCode/:tableName", tools.Gen)
			auth.GET("/gen/synchDb/:tableName", tools.Gen)
			auth.GET("/gen/batchGenCode", tools.GenBatchCode)
		}
	}
}
