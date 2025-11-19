package myknowledge

import "fmt"

// ** Quy tắc Export trong Go
// Cách viết				Trạng thái				Có thể truy cập từ package khác?
// Name:					Exported (Public)		Có thể dùng ngoài package
// name:					Unexported (Private)	Chỉ dùng trong package hiện tại
// PrintSomething():	 	Public function			Có thể gọi từ file khác
// printSomething(): 		Private function		Chỉ gọi trong package hiện tại

type body struct {
	height int
	width  int
}
type MyProfile struct {
	Name  string
	Phone string
	body  body
}

// khai báo hàm có receiver : (m MyProfile) trước tên hàm sẽ gán cho struct đó có method này
func (m MyProfile) getBodyHelth() float64 {
	return float64(m.body.height*m.body.width) / 100
}
func (m MyProfile) getPhone() string {
	return m.Phone
}

func getMyProfile() MyProfile {
	profile := MyProfile{
		Name:  "Giang",
		Phone: "0334644334",
	}
	return profile
}
func TestPrintStruct() {
	fmt.Print("----------------------------------------------\n")
	// ** cách 1 gọi hàm, và set cố định 1 số filed
	profile := getMyProfile()
	profile.body.height = 169
	profile.body.width = 72

	// ** cách 2: gán trực tiếp vào struct
	// profile := MyProfile{
	// 	Name:  "Giang",
	// 	Phone: "0334644334",
	// 	body: body{
	// 		height: 169,
	// 		width:  72,
	// 	},
	// }

	fmt.Printf("name: %s - phone: %s - body: %v, %v - body rate: %.2f %%\n", profile.Name,
		profile.getPhone(),
		profile.body.height,
		profile.body.width,
		profile.getBodyHelth())
	fmt.Print("==============================================\n")
}
