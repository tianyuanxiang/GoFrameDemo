// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GfUser is the golang structure for table gf_user.
type GfUser struct {
	Id        uint        `json:"id"        ` // UID
	Passport  string      `json:"passport"  ` // 账号
	Password  string      `json:"password"  ` // MD5密码
	Nickname  string      `json:"nickname"  ` // 昵称
	Avatar    string      `json:"avatar"    ` // 头像地址
	Status    int         `json:"status"    ` // 状态 0:启用 1:禁用
	Gender    int         `json:"gender"    ` // 性别 0: 未设置 1: 男 2: 女
	CreatedAt *gtime.Time `json:"createdAt" ` // 注册时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}
