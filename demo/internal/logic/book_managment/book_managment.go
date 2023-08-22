package book_managment

import (
	"context"
	"demo/api/v1"
	"demo/internal/dao"
	"demo/internal/service"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
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

func (s *sUser) Insert(ctx context.Context, in v1.BookInsertReq) (out *v1.BookInsertRes, err error) {
	// 判断in中的各字段在表中是否完全存在
	// select count(1) from library where name = '翦商' and ISBN = '132456789' and translator = 'qiqi'
	//and date = '2020-01-23' and publisher_id = 110;
	flag, err := g.Model("bookinformation").Ctx(ctx).Fields("count(1)", "ID").
		Where("BookName", in.Date.Name).Where("ISBN", in.Date.ISBN).Where("Author", in.Date.Author).Where("Publishers", in.Date.Publishers).Where("BookTypeID", in.Date.BookTypeID).Group("ID").All()
	if err != nil {
		return out, err
	}
	// 要插入的数据已经存在于表中
	if flag != nil {
		out = &v1.BookInsertRes{
			Message: "数据已存在，不可添加",
			Date:    v1.BookInformation{},
		}
		return
	}
	_, Err := g.Model("bookinformation").Insert(g.Map{"BookName": in.Date.Name, "ISBN": in.Date.ISBN,
		"Author": in.Date.Author, "Publishers": in.Date.Publishers, "BookTypeID": in.Date.BookTypeID, "Amount": in.Date.Amount})
	if Err != nil {
		err = Err
		return
	}
	// 图书类型表中的该类型图书数量 + 1
	updateData := g.Map{
		"Amount": &gdb.Counter{
			Field: "Amount",
			Value: 1,
		},
	}
	_, err = g.DB().Ctx(ctx).Update(ctx, "booktype", updateData, "BookTypeID", in.Date.BookTypeID)
	if err != nil {
		return
	}
	// 取新插入数据的ID
	MaxID, Err2 := g.Model("bookinformation").Fields("MAX(ID)").Value()

	All, Err3 := g.Model("bookinformation").Ctx(ctx).Where("ID", gconv.Int(MaxID)).All()
	if Err2 != nil {
		err = Err3
		return
	}
	out = &v1.BookInsertRes{
		Message: "插入的信息如下",
		Date: v1.BookInformation{
			Id:         gconv.Int(All[0]["ID"]),
			Name:       gconv.String(All[0]["BookName"]),
			ISBN:       gconv.String(All[0]["ISBN"]),
			Author:     gconv.String(All[0]["Author"]),
			Publishers: gconv.String(All[0]["Publishers"]),
			BookTypeID: gconv.Int(All[0]["BookTypeID"]),
			Amount:     gconv.Int(All[0]["Amount"]),
		},
	}
	return
}

// 得到书名和出版日期，判断有没有
// 都隶属于sUser这个结构体

func (s *sUser) Query(ctx context.Context, name string, ISBN string) (out *v1.BookQueryRes, err error) {
	// all里面保存的是：图书数量和id
	// 是0就不执行AND (`publisher_id`=PublisherId)，
	// 只执行WHERE (`name`='name') OR (`ISBN`='')
	// 不是0的话，就把这句加上，并且在Where("name", "许三观卖血记")与Where("publisher_id", pub)之间构造and
	object := g.Model("bookinformation").WhereLike("BookName", "%"+name+"%")
	if ISBN != "" {
		object = object.WhereLike("ISBN", "%"+ISBN+"%")
	}
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
		out = &v1.BookQueryRes{
			Message:     "查询结果为空",
			Information: []v1.BookInformation{},
			Flag:        false,
		}
		return
	}
	arrResult := make([]v1.BookInformation, 0)
	// 如果只有一个目标，可以直接改
	if len(Arr) == 1 {
		// 取出一个全局变量
		globalVariable = gconv.Int(Arr[0]["ID"])
		BookSingle := v1.BookInformation{
			Id:         gconv.Int(Arr[0]["ID"]),
			Name:       gconv.String(Arr[0]["BookName"]),
			ISBN:       gconv.String(Arr[0]["ISBN"]),
			Author:     gconv.String(Arr[0]["Author"]),
			Publishers: gconv.String(Arr[0]["Publishers"]),
			BookTypeID: gconv.Int(Arr[0]["BookTypeID"]),
			Amount:     gconv.Int(Arr[0]["Amount"]),
		}
		arrResult = append(arrResult, BookSingle)
		out = &v1.BookQueryRes{
			Message:     "图书信息如下",
			Information: arrResult,
			Flag:        true,
		}
	} else { // 多条记录
		for _, res := range Arr {
			BookMulti := v1.BookInformation{
				Id:         gconv.Int(res["ID"]),
				Name:       gconv.String(res["BookName"]),
				ISBN:       gconv.String(res["ISBN"]),
				Author:     gconv.String(res["Author"]),
				Publishers: gconv.String(res["Publishers"]),
				BookTypeID: gconv.Int(res["BookTypeID"]),
				Amount:     gconv.Int(res["Amount"]),
			}
			arrResult = append(arrResult, BookMulti)
		}
		out = &v1.BookQueryRes{
			Message:     "图书信息如下",
			Information: arrResult,
			Flag:        true,
		}
	}
	return
}

// 修改
func (s *sUser) Update(ctx context.Context, in v1.BookUpdateReq) (outUpdated *v1.BookUpdateRes, err error) {
	val := 0
	if in.Ret != 0 {
		val = in.Ret
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
			data.Set("BookName", in.Information.Name)
		}
		if gconv.String(in.Information.ISBN) != "" {
			data.Set("ISBN", in.Information.ISBN)
		}
		if gconv.String(in.Information.Author) != "" {
			data.Set("Author", in.Information.Author)
		}
		if gconv.String(in.Information.Publishers) != "" {
			data.Set("Publishers", in.Information.Publishers)
		}
		if gconv.Int(in.Information.BookTypeID) != 0 {
			data.Set("BookTypeID", in.Information.BookTypeID)
		}
		if gconv.Int(in.Information.Amount) != 0 {
			data.Set("Amount", in.Information.Amount)
		}
		_, err = g.Model("bookinformation").Data(data.Map()).Where("ID", val).Update()
		if err != nil {
			return err
		} else {
			return nil
		}
	})
	// update library set name = '背影',ISBN = '132456789',translator = 'qiqi', date = now(),publisher_id = '110' where id = 4;
	all2, err2 := g.Model("bookinformation").Ctx(ctx).Where("ID", val).All()
	if err2 != nil {
		err = err2
		return
	}
	outUpdated = &v1.BookUpdateRes{
		Message: "修改后的信息如下",
		Date: v1.BookInformation{
			Id:         gconv.Int(all2[0]["ID"]),
			Name:       gconv.String(all2[0]["BookName"]),
			ISBN:       gconv.String(all2[0]["ISBN"]),
			Author:     gconv.String(all2[0]["Author"]),
			Publishers: gconv.String(all2[0]["Publishers"]),
			BookTypeID: gconv.Int(all2[0]["BookTypeID"]),
			Amount:     gconv.Int(all2[0]["Amount"]),
		},
	}
	return
}

// 删除
func (s *sUser) Delete(ctx context.Context, in v1.BookDeleteReq) (OutDeleted *v1.BookDeleteRes, err error) {
	flag, err := s.Query(ctx, in.Name, in.ISBN)
	if err != nil {
		return
	}
	// 如果图书信息不存在
	if !flag.Flag {
		OutDeleted = &v1.BookDeleteRes{
			Information: []v1.BookInformation{},
			Message:     "图书信息不存在",
		}
		return
	}
	// 如果图书信息存在
	if flag.Flag {
		ArrResult2 := make([]v1.BookInformation, 0)
		// 如果为一条图书信息
		if len(flag.Information) == 1 {
			_, err2 := g.Model("bookinformation").Where("ISBN", flag.Information[0].ISBN).Delete()
			if err2 != nil {
				return
			}
			BookSingle := v1.BookInformation{
				Id:         flag.Information[0].Id,
				Name:       flag.Information[0].Name,
				ISBN:       flag.Information[0].ISBN,
				Author:     flag.Information[0].Author,
				Publishers: flag.Information[0].Publishers,
				BookTypeID: flag.Information[0].BookTypeID,
				Amount:     flag.Information[0].Amount,
			}
			ArrResult2 = append(ArrResult2, BookSingle)
			OutDeleted = &v1.BookDeleteRes{
				Message:     "删除信息如下",
				Information: ArrResult2,
			}
		}
		// 如果图书信息为多条
		if len(flag.Information) > 1 {
			// flag.Date已经是一个数组了，那就直接返回可以吗
			OutDeleted = &v1.BookDeleteRes{
				Message:     "请指定图书ISBN，再次进行删除",
				Information: flag.Information,
			}
		}
	}
	return
}

// func (s *sUser)AssignValue()
