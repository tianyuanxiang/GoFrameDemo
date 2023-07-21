package user

import (
	"context"
	v1 "demo/api/hello/v1"
	"demo/internal/dao"
	"demo/internal/service"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"
)

type (
	sUser struct{}
)

// 接口的具体实现注入
func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

var globalVariable = 0

// 新增
func (s *sUser) Insert(ctx context.Context, in v1.InsertReq) (out *v1.InsertRes, err error) {
	// 判断in中的各字段在表中是否完全存在
	// select count(1) from library where name = '翦商' and ISBN = '132456789' and translator = 'qiqi'
	//and date = '2020-01-23' and publisher_id = 110;
	flag, err := g.Model("library").Ctx(ctx).Fields("count(1)", "id").Where("name", in.Date.Name).Where("ISBN", in.Date.ISBN).Where(
		"translator", in.Date.Translator).Where("date", in.Date.Date).Where("publisher_id", in.Date.Publisher_id).Group("id").All()
	if err != nil {
		return out, err
	}
	// 要插入的数据已经存在于表中
	if flag != nil {
		out = &v1.InsertRes{
			Message: "数据已存在，不可添加",
			Date:    v1.BookInformation{},
		}
		return
	}
	_, Err := g.Model("library").Insert(g.Map{"name": in.Date.Name, "ISBN": in.Date.ISBN,
		"translator": in.Date.Translator, "date": in.Date.Date, "publisher_id": in.Date.Publisher_id})
	if Err != nil {
		err = Err
		return
	}
	// 取新插入数据的ID
	MaxID, Err2 := g.Model("library").Fields("MAX(id)").Value()

	All, Err3 := g.Model("library").Ctx(ctx).Where("id", gconv.Int(MaxID)).All()
	if Err2 != nil {
		err = Err3
		return
	}
	out = &v1.InsertRes{
		Message: "插入的信息如下",
		Date: v1.BookInformation{
			Id:           gconv.Int(All[0]["id"]),
			Name:         gconv.String(All[0]["name"]),
			ISBN:         gconv.String(All[0]["ISBN"]),
			Translator:   gconv.String(All[0]["translator"]),
			Date:         gconv.String(All[0]["date"]),
			Publisher_id: gconv.Int(All[0]["publisher_id"]),
		},
	}
	return
}

// 得到书名和出版日期，判断有没有
// 都隶属于sUser这个结构体
func (s *sUser) Query(ctx context.Context, name string, ISBN string, PublisherId int) (out *v1.Q_Res, err error) {
	// all里面保存的是：图书数量和id
	// all, err := g.Model("library").Ctx(ctx).Where("name", name).WhereOr("ISBN", ISBN).Where("publisher_id", PublisherId).All()

	// 是0就不执行AND (`publisher_id`=PublisherId)，
	// 只执行WHERE (`name`='name') OR (`ISBN`='')
	// 不是0的话，就把这句加上，并且在Where("name", "许三观卖血记")与Where("publisher_id", pub)之间构造and
	object := g.Model("library").Where("name", name)
	if PublisherId != 0 {
		object = object.Where("publisher_id", PublisherId)
	}
	object = object.WhereOr("ISBN", ISBN)
	all, err := object.All()
	if err != nil {
		return out, err
	}
	//创建map数组类型的切片数组，其中每个map数组负责存储像ISBN:78654132的元素
	Arr := make([]map[string]interface{}, 0)
	for _, element := range all {
		// 创建map数组存储每一个map,如arr1[ISBN]=78654132的元素
		arr1 := make(map[string]interface{}, 0)
		for key, value := range element {
			arr1[gconv.String(key)] = value
		}
		Arr = append(Arr, arr1)
	}
	if len(Arr) == 0 {
		g.RequestFromCtx(ctx).Response.ResponseWriter.WriteHeader(500)
		out = &v1.Q_Res{
			Message: "查询结果为空",
			Date:    []v1.BookInformation{},
			Flag:    false,
		}
		return
	}
	arrResult := make([]v1.BookInformation, 0)
	// 如果只有一个目标，可以直接改
	if len(Arr) == 1 {
		// 取出一个全局变量
		globalVariable = gconv.Int(Arr[0]["id"])
		BookSingle := v1.BookInformation{
			Id:           gconv.Int(Arr[0]["id"]),
			Name:         gconv.String(Arr[0]["name"]),
			ISBN:         gconv.String(Arr[0]["ISBN"]),
			Translator:   gconv.String(Arr[0]["translator"]),
			Date:         gconv.String(Arr[0]["date"]),
			Publisher_id: gconv.Int(Arr[0]["publisher_id"]),
		}
		arrResult = append(arrResult, BookSingle)
		out = &v1.Q_Res{
			Message: "图书信息如下",
			Date:    arrResult,
			Flag:    true,
		}
	} else { // 多条记录
		for _, res := range Arr {
			BookMulti := v1.BookInformation{
				Id:           gconv.Int(res["id"]),
				Name:         gconv.String(res["name"]),
				ISBN:         gconv.String(res["ISBN"]),
				Translator:   gconv.String(res["translator"]),
				Date:         gconv.String(res["date"]),
				Publisher_id: gconv.Int(res["publisher_id"]),
			}
			arrResult = append(arrResult, BookMulti)
		}
		out = &v1.Q_Res{
			Message: "图书信息如下",
			Date:    arrResult,
			Flag:    true,
		}
	}
	return
}

// 修改
func (s *sUser) Update(ctx context.Context, in v1.UpdateReq) (outUpdated *v1.UpdateRes, err error) {
	val := 0
	if in.Information.Id != 0 {
		val = in.Information.Id
	} else {
		val = globalVariable
	}

	// 开启事务
	// 当给定的闭包方法返回的error为nil时，闭包执行结束后当前事务自动执行Commit提交操作；否则自动执行Rollback回滚操作。
	dao.GfUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 为了防止空字段覆盖原始字段，提前设置一个map对象备份（倒腾）
		data := gmap.New()
		// 不为""，即证明有值，也就是需要修改，其余的不用修改，晕了
		if gconv.String(in.Information.Name) != "" {
			data.Set("name", in.Information.Name)
		}
		if gconv.String(in.Information.ISBN) != "" {
			data.Set("ISBN", in.Information.ISBN)
		}
		if gconv.String(in.Information.Translator) != "" {
			data.Set("translator", in.Information.Translator)
		}
		if gconv.String(in.Information.Date) != "" {
			data.Set("date", in.Information.Date)
		}
		if gconv.Int(in.Information.Publisher_id) != 0 {
			data.Set("publisher_id", in.Information.Publisher_id)
		}
		_, err = g.Model("library").Data(data.Map()).Where("id", val).Update()
		if err != nil {
			return err
		} else {
			return nil
		}
	})
	// update library set name = '背影',ISBN = '132456789',translator = 'qiqi', date = now(),publisher_id = '110' where id = 4;
	all2, err2 := g.Model("library").Ctx(ctx).Where("id", val).All()
	if err2 != nil {
		err = err2
		return
	}
	outUpdated = &v1.UpdateRes{
		Message: "修改后的信息如下",
		Date: v1.BookInformation{
			Id:           gconv.Int(all2[0]["id"]),
			Name:         gconv.String(all2[0]["name"]),
			ISBN:         gconv.String(all2[0]["ISBN"]),
			Translator:   gconv.String(all2[0]["translator"]),
			Date:         gconv.String(all2[0]["date"]),
			Publisher_id: gconv.Int(all2[0]["publisher_id"]),
		},
	}
	return
}

// 删除
func (s *sUser) Delete(ctx context.Context, in v1.DeleteReq) (OutDeleted *v1.DeleteRes, err error) {
	flag, err := s.Query(ctx, in.Name, in.ISBN, in.PublisherID)
	if err != nil {
		return
	}
	// 如果图书信息不存在
	if !flag.Flag {
		OutDeleted = &v1.DeleteRes{
			Date:    []v1.BookInformation{},
			Message: "图书信息不存在",
		}
		return
	}
	// 如果图书信息存在
	if flag.Flag {
		ArrResult2 := make([]v1.BookInformation, 0)
		// 如果为一条图书信息
		if len(flag.Date) == 1 {
			_, err2 := g.Model("library").Where("publisher_id", flag.Date[0].Publisher_id).Delete()
			if err2 != nil {
				return
			}
			bookSingle := v1.BookInformation{
				Id:           flag.Date[0].Id,
				Name:         flag.Date[0].Name,
				ISBN:         flag.Date[0].ISBN,
				Translator:   flag.Date[0].Translator,
				Date:         flag.Date[0].Date,
				Publisher_id: flag.Date[0].Publisher_id,
			}
			ArrResult2 = append(ArrResult2, bookSingle)
			OutDeleted = &v1.DeleteRes{
				Message: "删除信息如下",
				Date:    ArrResult2,
			}
		}
		// 如果图书信息为多条
		if len(flag.Date) > 1 {
			// flag.Date已经是一个数组了，那就直接返回可以吗
			OutDeleted = &v1.DeleteRes{
				Message: "请指定图书publisher_id，再次进行删除",
				Date:    flag.Date,
			}
		}
	}
	return
}

// func (s *sUser)AssignValue()
