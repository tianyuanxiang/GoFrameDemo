package controller

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
)

type uPersonalProfile struct{}

var UPersonalProfile = new(uPersonalProfile)

func (u *uPersonalProfile) QueryPersonalProfile(ctx context.Context, in *v1.PersonalProfileQueryReq) (out *v1.PersonalProfileQueryRes, err error) {
	num, err := service.PersonsalProfile().QueryPersonalProfile(ctx, in)
	if err != nil {
		return
	}
	out = &v1.PersonalProfileQueryRes{
		Message:                    num.Message,
		PersonalInformationDisplay: num.PersonalInformationDisplay,
	}
	return
}

func (u *uPersonalProfile) ModifyPersonalProfile(ctx context.Context, in *v1.PersonalProfileModifyReq) (out *v1.PersonalProfileModifyRes, err error) {
	num, err := service.PersonsalProfile().ModifyPersonalProfile(ctx, in)
	if err != nil {
		return
	}
	out = &v1.PersonalProfileModifyRes{
		Message:          num.Message,
		UserIP:           num.UserIP,
		ModifiedUserName: num.ModifiedUserName,
		ModifiedEmail:    num.ModifiedEmail,
	}
	return
}
