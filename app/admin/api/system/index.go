package system

import (
	"github.com/gin-gonic/gin"
	"ruoyi-go/utils/R"
)

// 后台获取 首页数据

func IndexData() {
	//
}

// IndexHandler 测试代码
func IndexHandler(context *gin.Context) {
	R.ReturnSuccess("Hello ruoyi go")
}
