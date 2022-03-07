// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package model

import (
	comModel "gfast/app/common/model"
	"gfast/app/system/model/internal"
)

// SysRole is the golang structure for table sys_role.
type SysRole internal.SysRole

// Fill with you ideas below.
//分页请求参数
type SelectPageReq struct {
	RoleName string `p:"roleName"` //参数名称
	Status   string `p:"status"`   //状态
	comModel.PageReq
}

//修改状态参数
type StatusSetReq struct {
	RoleId uint `p:"roleId" v:"required#角色ID不能为空"`
	Status uint `p:"status" v:"required#状态不能为空"`
}

//角色数据授权参数
type DataScopeReq struct {
	RoleId    uint   `p:"roleId" v:"required#角色ID不能为空"`
	DataScope uint   `p:"dataScope" v:"required#权限范围不能为空"`
	DeptIds   []uint `p:"deptIds"`
}