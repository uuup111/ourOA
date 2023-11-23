package router

import (
	"github.com/gin-gonic/gin"
	 api "ruoyi-go/app/admin/api/system"
	"ruoyi-go/app/pkg/utils"
)

func initSysUser(e *gin.Engine) {
	// 路由权限相关
	v2 := e.Group("system")
	{
		auth := v2.Group("")
		auth.Use(utils.JWTAuthMiddleware())
		{
			// 查询
            auth.GET("/user/list", api.GetPageLimitSysUser)
            // 添加
            auth.POST("/user", api.CreateSysUser)
            // 修改
            auth.PUT("/user", api.UpdateSysUser)
            // 删除
            auth.DELETE("/user/:ids", api.DeleteSysUserByIds)
            // 获取详情
            auth.GET("/user/:id", api.GetSysUser)
		}
	}
}
