package system

import (
    "strconv"
	"github.com/gin-gonic/gin"
	"ruoyi-go/app/admin/model/tools"
	"log"
	"net/http"
	"ruoyi-go/app/pkg/utils/R"
    model "ruoyi-go/app/admin/model/system"
)

func CreateSysUser(context *gin.Context) {
	userId, _ := context.Get("userId")
	log.Println(userId)
	var req model.SysUser
    if err := context.ShouldBindJSON(&req);err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("获取参数失败"))
	} else {
		res := model.CreateSysUserService(req)
		context.JSON(http.StatusOK, res)
	}
}

func DeleteSysUser(c *gin.Context) {
	var id = context.Param("id")
    result := model.DeleteSysUserService(id)
    context.JSON(http.StatusOK, result)
}

func DeleteSysUserByIds(c *gin.Context) {
    var ids = context.Param("ids")
    result := model.DeleteSysUserByIdsService(ids)
    context.JSON(http.StatusOK, result)
}

func UpdateSysUser(c *gin.Context) {
    var req model.SysUser
    if err := context.ShouldBindJSON(&req);err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("获取参数失败"))
	} else {
		res := model.UpdateSysUserService(req)
		context.JSON(http.StatusOK, res)
	}
}

func GetSysUser(c *gin.Context) {
    idstr:=context.Param("id")    //查询路径Path参数
    id, err := strconv.ParseInt(idstr, 10, 64)
    if err!=nil{
        context.JSON(http.StatusOK, R.ReturnFailMsg("获取参数失败"))
    }else{
        res := model.GetSysUserService(id)
    	context.JSON(http.StatusOK, res)
    }
}

func GetPageLimitSysUser(c *gin.Context) {
    pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
    pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

    beginTime := context.DefaultQuery("params[beginTime]", "")
    endTime := context.DefaultQuery("params[endTime]", "")

    var param = tools.SearchTableDataParam{
        PageNum:  pageNum,
        PageSize: pageSize,
        Other: model.SysUser{
        },
        Params: tools.Params{
            BeginTime: beginTime,
            EndTime:   endTime,
        },
    }

    result := model.GetPageLimitSysUserService(param)
    context.JSON(http.StatusOK, result)
}
