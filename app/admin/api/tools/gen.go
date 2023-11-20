package tools

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ruoyi-go/app/admin/model/tools"
	tools2 "ruoyi-go/app/admin/service/tools"
	"ruoyi-go/utils"
	"ruoyi-go/utils/R"
	"strconv"
)

func GenList(context *gin.Context) {
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	var tableName = context.DefaultQuery("tableName", "")
	var tableComment = context.DefaultQuery("tableComment", "")

	var beginTime = context.DefaultQuery("params[beginTime]", "")
	var endTime = context.DefaultQuery("params[endTime]", "")

	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: tools.GenTable{
			Table_Name:   tableName,
			TableComment: tableComment,
		},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	context.JSON(http.StatusOK, tools2.SelectGenList(param, true))
}

func GenInfo(context *gin.Context) {
	tableId := context.Param("tableId")
	result := tools2.GenInfoService(tableId)
	context.JSON(http.StatusOK, R.ReturnSuccess(result))
}

func GenDbList(context *gin.Context) {
	pageNum, _ := strconv.Atoi(context.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	var tableName = context.DefaultQuery("tableName", "")
	var tableComment = context.DefaultQuery("tableComment", "")

	var beginTime = context.DefaultQuery("params[beginTime]", "")
	var endTime = context.DefaultQuery("params[endTime]", "")
	var param = tools.SearchTableDataParam{
		PageNum:  pageNum,
		PageSize: pageSize,
		Other: tools.GenTable{
			Table_Name:   tableName,
			TableComment: tableComment,
		},
		Params: tools.Params{
			BeginTime: beginTime,
			EndTime:   endTime,
		},
	}
	context.JSON(http.StatusOK, tools2.GenDbList(param))
}

func GenColumnInfo(context *gin.Context) {
	tableId := context.Param("tableId")
	log.Println(tableId)
}

func ImportTable(context *gin.Context) {
	userId, _ := context.Get("userId")
	var tables = context.DefaultQuery("tables", "")
	table := utils.SplitStr(tables)
	result := tools2.SelectDbTableListByNames(table)
	tools2.ImportGenTable(result, utils.GetInterfaceToInt(userId))
	context.JSON(http.StatusOK, R.ReturnSuccess("成功"))
}

func GenEdit(context *gin.Context) {
	var req tools.EditGenTableVO
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusOK, R.ReturnFailMsg("获取参数失败"))
	} else {
		res := tools2.UpdateGenTableService(req)
		context.JSON(http.StatusOK, res)
	}
}

func GenDelete(context *gin.Context) {
	tableIds := context.Param("tableIds")
	tools2.DeleteGenTableByIds(tableIds)
	tools2.DeleteGenTableColumnByIds(tableIds)
}

func PreviewGenTable(context *gin.Context) {
	tableId := context.Param("tableId")
	result := tools2.PreviewGenTableCode(tableId)
	context.JSON(http.StatusOK, R.ReturnSuccess(result))
}

func GenDownload(context *gin.Context) {
	tableName := context.Param("tableName")
	log.Println(tableName)
}

func GenBatchCode(context *gin.Context) {
	var tables = context.DefaultQuery("tables", "")
	table := utils.SplitStr(tables)
	const fileName = "RuoYi-Go.zip"
	filePath := tools2.GenBatchCode(table, fileName)
	utils.DonwloadFile(context, fileName, filePath)
	utils.DirExistAndDel(filePath)
}

func Gen(context *gin.Context) {
	tableName := context.Param("tableName")
	log.Println(tableName)
	context.JSON(http.StatusOK, R.ReturnFailMsg("暂不支持"))
}
