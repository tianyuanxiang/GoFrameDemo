package u_personalprofile

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sPersonsalProfile struct{}

func New() *sPersonsalProfile {
	return &sPersonsalProfile{}
}

func init() {
	service.RegisterPersonsalProfile(New())
}

// 用户查询个人信息

func (s *sPersonsalProfile) QueryPersonalProfile(ctx context.Context, req *v1.PersonalProfileQueryReq) (res *v1.PersonalProfileQueryRes, err error) {
	// 分开操作，先查询个人信息，再查询图书数量
	// 查询个人信息
	PersonalProfile, err := g.Model("userinformation").Ctx(ctx).Where("UserIP", req.UserIP).Where("UserName", req.UserName).All()
	if err != nil {
		return
	}
	// 查询所借阅的图书数量，既然是两个查询，就不用非得联表了
	// select BookName, count(1) from bookborrowinformation where UserIP=19822350 and Flag=1 group by BookName;
	NumberOfLoans, err := g.Model("bookborrowinformation b").Fields("BookName, count(1)").Where("UserIP", req.UserIP).Where("Flag", 1).Group("BookName").All()
	if err != nil {
		return
	}
	// Data最后以数组的形式返回
	ret := garray.New()
	for _, record := range NumberOfLoans {
		nodes := gmap.New()
		nodes.Set(gconv.String(record["BookName"]), gconv.Int(record["count(1)"]))
		ret.Append(nodes)
	}
	// 终于可以返回了
	DataPadding := v1.PersonalInformation{
		UserIP:     gconv.String(PersonalProfile[0]["UserIP"]),
		UserName:   gconv.String(PersonalProfile[0]["UserName"]),
		Email:      gconv.String(PersonalProfile[0]["Email"]),
		CurrentNum: gconv.Int(PersonalProfile[0]["CurrentNum"]),
		Data:       ret,
		HistoryNum: gconv.Int(PersonalProfile[0]["HistoryNum"]),
	}
	res = &v1.PersonalProfileQueryRes{
		Message:                    "该用户的信息如下：",
		PersonalInformationDisplay: DataPadding,
	}
	return
}

func (s *sPersonsalProfile) ModifyPersonalProfile(ctx context.Context, req *v1.PersonalProfileModifyReq) (res *v1.PersonalProfileModifyRes, err error) {
	data := gmap.New()
	if req.UserName != "" {
		data.Set("UserName", req.UserName)
	}
	if req.Email != "" {
		data.Set("Email", req.Email)
	}
	_, err = g.Model("userinformation").Ctx(ctx).Data(data.Map()).Where("UserIP", req.UserIP).Update()
	if err != nil {
		return
	}
	all, err := g.Model("userinformation").Ctx(ctx).Where("UserIP", req.UserIP).All()
	res = &v1.PersonalProfileModifyRes{
		Message:          "修改后的个人资料如下：",
		UserIP:           gconv.String(all[0]["UserIP"]),
		ModifiedUserName: gconv.String(all[0]["UserName"]),
		ModifiedEmail:    gconv.String(all[0]["Email"]),
	}
	return
}
