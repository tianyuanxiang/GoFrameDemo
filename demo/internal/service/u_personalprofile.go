// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "demo/api/v1"
)

type (
	IPersonsalProfile interface {
		QueryPersonalProfile(ctx context.Context, req *v1.PersonalProfileQueryReq) (res *v1.PersonalProfileQueryRes, err error)
		ModifyPersonalProfile(ctx context.Context, req *v1.PersonalProfileModifyReq) (res *v1.PersonalProfileModifyRes, err error)
	}
)

var (
	localPersonsalProfile IPersonsalProfile
)

func PersonsalProfile() IPersonsalProfile {
	if localPersonsalProfile == nil {
		panic("implement not found for interface IPersonsalProfile, forgot register?")
	}
	return localPersonsalProfile
}

func RegisterPersonsalProfile(i IPersonsalProfile) {
	localPersonsalProfile = i
}
