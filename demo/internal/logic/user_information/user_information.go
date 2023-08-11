package user_information

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sUserInformation struct{}

func New() *sUserInformation {
	return &sUserInformation{}
}

func init() {
	service.RegisterUserInformation(New())
}

// 查询读者信息方法
func (s *sUserInformation) UserInformationQuery(ctx context.Context, req *v1.ReaderInformationReq) (res *v1.ReaderInformationRes, err error) {
	object := g.Model("UserInformation").Ctx(ctx).WhereLike("UserName", "%"+req.UserName+"%")
	if req.UserIP != "" {
		object = object.Where("UserIP", req.UserIP)
	}
	if req.ID != 0 {
		object = object.Where("ID", req.ID)
	}
	user, err := object.All()
	if err != nil {
		return
	}
	if user == nil {
		res = &v1.ReaderInformationRes{
			Message: "该用户不存在！",
			User:    &[]v1.UserInformation{},
		}
		return
	}
	// 为了支持模糊查询
	ArrUser := make([]v1.UserInformation, 0)
	for _, people := range user {
		ele := v1.UserInformation{
			Id:         gconv.Int(people["Id"]),
			UserIP:     gconv.String(people["UserIP"]),
			UserName:   gconv.String(people["UserName"]),
			Email:      gconv.String(people["Email"]),
			CurrentNum: gconv.Int(people["CurrentNum"]),
			HistoryNum: gconv.Int(people["HistoryNum"]),
		}
		ArrUser = append(ArrUser, ele)
	}
	res = &v1.ReaderInformationRes{
		Message: "读者信息如下：",
		User:    &ArrUser,
	}
	return
}

// 修改读者信息方法
func (s *sUserInformation) UserInformationModify(ctx context.Context, req *v1.ReaderModifyReq) (res *v1.ReaderModifyRes, err error) {
	return
}

// 添加读者信息方法
func (s *sUserInformation) UserInformationAdd(ctx context.Context, req *v1.ReaderAddReq) (res *v1.ReaderAddRes, err error) {
	// 添加之前判断一下该用户是否已经存在，通过UserIP、UserName和Email联合判断
	// select * from UserInformation where UserIP = req.UserIP and UserName = req.UserName and Email = req.Email
	number, err := g.Model("UserInformation").Fields("count(1)").Ctx(ctx).Where(g.Map{"UserIP": req.UserIP, "UserName": req.UserName, "Email": req.Email}).All()
	if err != nil {
		return nil, err
	}
	if gconv.Int(number) > 1 {
		res = &v1.ReaderAddRes{
			Message:   "用户信息已存在,请勿重复添加！",
			UserAdded: v1.UserInformation{},
		}
		return
	} else {
		// _, Err := g.Model("bookinformation").Insert(g.Map{"BookName": in.Date.Name, "ISBN": in.Date.ISBN,
		//		"Author": in.Date.Author, "Publishers": in.Date.Publishers, "BookTypeID": in.Date.BookTypeID, "Amount": in.Date.Amount})
		data := g.Map{"UserIP": req.UserIP, "UserName": req.UserName, "Email": req.Email, "CurrentNum": req.CurrentNum, "HistoryNum": req.HistoryNum}
		_, err1 := g.Model("UserInformation").Data(data).Insert()
		if err1 != nil {
			return
		}
		flag, err2 := g.Model("UserInformation").Where(data).All()
		if err2 != nil {
			return
		}
		if flag == nil {
			res = &v1.ReaderAddRes{
				Message:   "用户信息添加失败！",
				UserAdded: v1.UserInformation{},
			}
			return
		}
		res = &v1.ReaderAddRes{
			Message: "用户信息添加成功，信息如下：",
			UserAdded: v1.UserInformation{
				Id:         gconv.Int(flag[0]["ID"]),
				UserIP:     gconv.String(flag[0]["UserIP"]),
				UserName:   gconv.String(flag[0]["UserName"]),
				Email:      gconv.String(flag[0]["Email"]),
				CurrentNum: gconv.Int(flag[0]["CurrentNum"]),
				HistoryNum: gconv.Int(flag[0]["HistoryNum"]),
			},
		}
	}
	return
}
