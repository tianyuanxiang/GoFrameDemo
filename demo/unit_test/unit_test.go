package unit_test

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2" // 导入MySQL驱动
)

func TestDC(t *testing.T) {
	//s1 := "红楼梦"
	//s2 := ""
	// SELECT * FROM `library` WHERE  sss ='ISBN' or sss = 'name';

	// count, err := g.Model("library").Where("name", sss).WhereOr("ISBN", sss).Count()
	//if err != nil {
	//	return
	//}
	//fmt.Println(count)

	// SELECT count(1),id FROM library where name='' or ISBN = '988732164' group by id;
	// count, err := g.Model("library").Fields("count(1),id").Where("name", s1).WhereOr("ISBN", s2).Group("id").All()

	// update library set Name = '罗辑思维',ISBN = '132456789' where id = id;
	//Name := "人生海海"
	//ISBN := "464646465"
	//translator := "麦家"
	//date := "2010-11-23"
	//publisher_id := 310
	//_, err := g.Model("library").Data(g.Map{"name": Name, "ISBN": ISBN, "translator": translator,
	//	"date": date, "publisher_id": publisher_id}).Where("id", 1).Update()
	//flag, err := g.Model("library").Fields("count(1)", "id").Where("name", Name).Where("ISBN", ISBN).Where(
	//	"translator", translator).Where("date", date).Where("publisher_id", publisher_id).Group("id").All()
	////all2, _ := g.Model("library").Where("id", 1).All()
	//fmt.Println(flag)
	//_, Err := g.Model("library").Insert(g.Map{"name": Name, "ISBN": ISBN,
	//	"translator": translator, "date": date, "publisher_id": publisher_id})
	//all, err := g.Model("library").Data(g.Map{"name": "罗辑思维", "ISBN": "132456789"}).Where("id", id).Update()
	//if err != nil {
	//	return
	//}
	MaxID, Err2 := g.Model("library").Fields("MAX(id)").Value()
	fmt.Println(MaxID)
	// data, err := g.Model("library").Where("id", id).All()
	if Err2 != nil {

		return
	}
}

//	func TestUid(t *testing.T) {
//		get, err := g.Cfg().Get(gctx.New(), "kafka")
//		if err != nil {
//			return
//		}
//		fmt.Println(get)
//
// }
type BookInformation struct {
	Id           int    ` dc:"book_id" json:"id"`
	Name         string ` json:"book_name"`
	ISBN         string `json:"Book_ISBN"`
	Translator   string `json:"translator"`
	Date         string ` json:"publish_date"`
	Publisher_id int    `json:"publisher_Id"`
}
type InsertRes struct {
	Message string          `json:"Message"`
	Date    BookInformation `json:"date"`
}

func TestU(t *testing.T) {
	//	s := garray.New()
	//all, _ := g.Model("library").Where("name", "光荣与梦想").WhereOr("ISBN", "").WhereOr("publisher_id", "").All()
	//fmt.Println(all)
	//all, _ := g.Model("library").Where("name", "光荣与梦想").WhereOr("ISBN", "").WhereOr("publisher_id", 0).All()

	// pub := 输入的publisher_id
	pub := 10086
	obj := g.Model("library").Where("name", "许三观卖血记")
	// 是0就不执行AND (`publisher_id`=10086)，
	// 只执行WHERE (`name`='许三观卖血记') OR (`ISBN`='')
	// 不是0的话，就把这句加上，并且在Where("name", "许三观卖血记")与Where("publisher_id", pub)之间构造and
	if pub != 0 {
		obj = obj.Where("publisher_id", pub)
	}
	obj = obj.WhereOr("ISBN", "")
	// 不管之前怎么样，在这里统一才执行
	all, _ := obj.All()
	//var Se []v1.BookInformation
	fmt.Println(all)
	// s3 := make([]map[string]interface{}, 0)
	// s2 := make(map[string]interface{}, 1)
	//for j, record := range all {
	//	// for j, value := range record {
	//	value := record[gconv.String(j)]
	//	fmt.Println(value)
	//}

}

// fmt.Println(s3)
// var array_name [array_size]data_type
// 集合类型的数组

// 集合类型的数组，InsertRes是一个集合
//var Se []v1.InsertRes
//
//Sr1 := new(v1.InsertRes)
//Sr1.Message = "number 1"
//Sr1.Date.ISBN = "132134564"
//Sr1.Date.Name = "az"
//Sr1.Date.Publisher_id = 456456
//
//Sr2 := new(v1.InsertRes)
//Sr2.Message = "number 2"
//Sr2.Date.ISBN = "9637894625"
//Sr2.Date.Name = "tj"
//Sr2.Date.Publisher_id = 123456
//
//Se = append(Se, *Sr1)
//Se = append(Se, *Sr2)
//
//fmt.Println(Se)
//
//fmt.Println(Se[0])
//fmt.Println(Se[1])
//
//fmt.Println(Se[0].Date.Name)
//fmt.Println(Se[0].Date.ISBN)

// fmt.Println(se[0].Date)
// fmt.Println(se[0].Message)
// var Se []v1.InsertRes
// Sr1 := new(v1.InsertRes)
// all, err := g.Model("library").Ctx(gctx.New()).Fields("date,name").Where("name = '光荣与梦想'").All()
//
//	if err != nil {
//		return
//	}
//
// fmt.Println(all)
//
//	for i, record := range all {
//		fmt.Println(i)
//		append(Se, record.Map())
//
// }
//
// fmt.Println(s)
func TestK(t *testing.T) {

	//pub, err := g.DB().Ctx(gctx.New()).Model("bookborrowinformation").Where("ID", 2).All()
	//if err != nil {
	//	return
	//}
	////var Se []v1.BookInformation
	//fmt.Println(pub)
	//BookTypeID := 102
	//
	//pub, err := g.DB().Ctx(gctx.New()).Model("booktype").
	//	Where("BookTypeID", BookTypeID).All()
	pub, err := g.Model("booktype").All()
	fmt.Println(pub)
	if err != nil {
		return
	}
}
