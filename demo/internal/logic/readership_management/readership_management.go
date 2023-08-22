package readership_management

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sReaderShipManagement struct{}

func New() *sReaderShipManagement {
	return &sReaderShipManagement{}
}

func init() {
	service.RegisterReaderShipManagement(New())
}

// 查询读者信息

func (s *sReaderShipManagement) ReaderShipQuery(ctx context.Context, req *v1.ReadershipQueryReq) (res *v1.ReadershipQueryRes, err error) {
	UserData, err := g.Model("userinformation").Ctx(ctx).All()
	if err != nil {
		return
	}
	CarrierArray := make([]v1.UserProfile, 0)
	for _, num := range UserData {
		K := v1.UserProfile{
			UserIP:     gconv.String(num["UserIP"]),
			UserName:   gconv.String(num["UserName"]),
			Email:      gconv.String(num["Email"]),
			CurrentNum: gconv.Int(num["CurrentNum"]),
			HistoryNum: gconv.Int(num["HistoryNum"]),
		}
		CarrierArray = append(CarrierArray, K)
	}
	res = &v1.ReadershipQueryRes{
		Message:          "用户信息如下：",
		UserProfileGroup: &CarrierArray,
	}
	return
}

// 修改读者信息

func (s *sReaderShipManagement) ReaderShipModify(ctx context.Context, req *v1.ReadershipModifyReq) (res *v1.ReadershipModifyRes, err error) {
	if req.UserName == "" && req.Email != "" {
		_, err = g.Model("userinformation").Ctx(ctx).Data(g.Map{"Email": req.Email}).Where("UserIP", req.UserIP).Update()
		if err != nil {
			return
		}
	}
	if req.UserName != "" {
		object := gmap.New()
		object.Set("UserName", req.UserName)
		if req.Email != "" {
			object.Set("Email", req.Email)
		}
		_, err = g.Model("userinformation").Ctx(ctx).Data(object).Where("UserIP", req.UserIP).Update()
		if err != nil {
			return
		}
		_, err = g.Model("bookborrowinformation").Ctx(ctx).Data(g.Map{"UserName": req.UserName}).Where("UserIP", req.UserIP).Update()
		if err != nil {
			return
		}
		ModifiedRes, err2 := g.Model("userinformation").Ctx(ctx).Where("UserIP", req.UserIP).All()
		if err2 != nil {
			return
		}
		res = &v1.ReadershipModifyRes{
			Message: "修改后的用户信息如下：",
			Modified: v1.UserProfile{
				UserIP:     gconv.String(ModifiedRes[0]["UserIP"]),
				UserName:   gconv.String(ModifiedRes[0]["UserName"]),
				Email:      gconv.String(ModifiedRes[0]["Email"]),
				CurrentNum: gconv.Int(ModifiedRes[0]["CurrentNum"]),
				HistoryNum: gconv.Int(ModifiedRes[0]["HistoryNum"]),
			},
		}
	}
	return
}
