package system

import (
    "ruoyi-go/app/pkg/utils"
    "ruoyi-go/app/pkg/utils/R"
)

// Model sys_user  用户信息表
type SysUser struct {
     UserId int64 `json:"userId" gorm:"primaryKey;column:user_id"`//用户ID
     DeptId int64 `json:"deptId" gorm:"dept_id"`//部门ID
     UserName string `json:"userName" gorm:"user_name"`//用户账号
     NickName string `json:"nickName" gorm:"nick_name"`//用户昵称
     UserType string `json:"userType" gorm:"user_type"`//用户类型（00系统用户）
     Email string `json:"email" gorm:"email"`//用户邮箱
     Phonenumber string `json:"phonenumber" gorm:"phonenumber"`//手机号码
     Sex string `json:"sex" gorm:"sex"`//用户性别（0男 1女 2未知）
     Avatar string `json:"avatar" gorm:"avatar"`//头像地址
     Password string `json:"password" gorm:"password"`//密码
     Status string `json:"status" gorm:"status"`//帐号状态（0正常 1停用）
     DelFlag string `json:"delFlag" gorm:"del_flag"`//删除标志（0代表存在 2代表删除）
     LoginIp string `json:"loginIp" gorm:"login_ip"`//最后登录IP
     LoginDate time.Time `json:"loginDate" gorm:"login_date"`//最后登录时间
     CreateBy string `json:"createBy" gorm:"create_by"`//创建者
     CreateTime time.Time `json:"createTime" gorm:"create_time"`//创建时间
     UpdateBy string `json:"updateBy" gorm:"update_by"`//更新者
     UpdateTime time.Time `json:"updateTime" gorm:"update_time"`//更新时间
     Remark string `json:"remark" gorm:"remark"`//备注
}


func (SysUser) TableName() string {
  return "sys_user"
}

//创建
func CreateSysUserService(data SysUser) any {
    err := utils.MysqlDb.Create(&data).Error
	if err!=nil{
        return R.ReturnFailMsg("保存失败")
	}
	return R.ReturnSuccess(nil)
}

//根据ID删除
func DeleteSysUserService(id string) any {
    err := utils.MysqlDb.Delete(&SysUser{},"id = ?",id).Error
	if err!=nil{
        return R.ReturnFailMsg("删除失败")
    }

    return R.ReturnSuccess(nil)
}

//根据ID批量删除
func DeleteSysUserByIdsService(ids string) any {
	if len(ids) == 0{
         return R.ReturnFailMsg("参数获取失败")
    }

    err := utils.MysqlDb.Delete(&SysUser{},"id in ?",ids).Error
    if err!=nil{
       return R.ReturnFailMsg("批量删除失败")
    }

    return R.ReturnSuccess(nil)
}

//根据id 更新 ，排除零值
func UpdateSysUserService(data SysUser) any {
     err := utils.MysqlDb.Updates(&data).Error
    if err!=nil{
        return R.ReturnFailMsg("更新失败")
    }
    return R.ReturnSuccess(nil)
}

//根据id获取model
func GetSysUserService(id int64) any {

    if id == 0{
        return R.ReturnFailMsg("参数获取失败")
    }
    var data SysUser
    err := utils.MysqlDb.Where("id = ?", id).First(&data).Error
    if err!=nil{
        return R.ReturnFailMsg("获取数据失败")
    }else{
        return R.ReturnSuccess(data)
    }

}

//获取所有的model
func GetListSysUserService() any {
    var list []SysUser
    err := utils.MysqlDb.Find(&list).Error
	if err!=nil {
        return R.ReturnFailMsg("获取数据失败")
    }else{
        return R.ReturnSuccess(list)
    }
}

//按条件分页查询 limit offset ,参数用指针&, 数据会自动填充到req对象
func GetPageLimitSysUserService(params tools.SearchTableDataParam) tools.TableDataInfo {

    var pageNum = params.PageNum
	var pageSize = params.PageSize

	var total int64
	db := utils.MysqlDb.Model(SysUser{})
	// 可以自定义搜索方式

	var rows []SysUser

	if err := db.Count(&total).Error; err != nil {
		return tools.Fail()
	}
	offset := (pageNum - 1) * pageSize
	db.Order("id DESC").Offset(offset).Limit(pageSize).Find(&rows)
	if rows == nil {
        return tools.Fail()
    } else {
        return tools.Success(rows, total)
    }
}