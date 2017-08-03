package ibase

import (
	"fmt"
	"reflect"
)

type VehicleInfo struct {
	// ID         bson.ObjectId `bson:"_id,omitempty"`
	VehicleId string `bson:"编号"`
	Date      string `bson:"日期"`
	Type      string `bson:"类型"`
	Brand     string `bson:"型号"`
	Color     string `bson:"颜色"`
}

func main() {
	vinfo := &VehicleInfo{
		VehicleId: "123456a",
		Date:      "20140101a",
		Type:      "Trucka",
		Brand:     "Forda",
		Color:     "Whitea",
	}
	vinfo2 := &VehicleInfo{
		VehicleId: "123456",
		Date:      "20140101",
		Type:      "Truck",
		Brand:     "Ford",
		Color:     "White",
	}
	vt := reflect.TypeOf(vinfo).Elem()
	vv := reflect.ValueOf(vinfo).Elem()
	vt2 := reflect.TypeOf(vinfo2).Elem()
	vv2 := reflect.ValueOf(vinfo2).Elem()
	for i := 0; i < vt.NumField(); i++ {
		f := vt.Field(i)
		chKey := f.Tag.Get("bson")
		fmt.Printf("%q => %q \n", chKey, vv.FieldByName(f.Name).String())
		fmt.Println("")
		fmt.Printf("%q => %q ", f.Name, f.Type)
		for i := 0; i < vt2.NumField(); i++ {
			f2 := vt2.Field(i)
			if f.Name == f2.Name && f.Type == f2.Type {
				structFieldValue := vv2.FieldByName(f2.Name)
				if !structFieldValue.CanSet() {
					fmt.Errorf("struct 字段 %s 不能set,请检查!!!", f2.Name)
				} else {
					structFieldValue.Set(vv.FieldByName(f.Name))
				}
			}
		}
	}
	fmt.Printf("===================%+v\n", vinfo2)
}
