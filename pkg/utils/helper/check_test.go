package helper

import (
	"fmt"
	"testing"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service"
)

// func TestCheck(t *testing.T) {
// 	result :=HashPassword("123")
// 	t.Logf("%s", result)
// }

// func Test_config(t *testing.T) {
// 	config, _ := config.LoadConfig()
// 	t.Logf("%s", config)
// 	fmt.Println("---------",config)
// }

// func Test_token(t *testing.T){
// 	phone,err:=service.FetchPhoneFromToken( "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJQaG9uZSI6Ijk3NDQ3MDM1NTgifQ.VvEB7YffTW_8taP6cpZBnPDgTtMfz6ZSnXoFw4jq6bI", "tempervary12345jlkh")
// 	if err!=nil{
// 		fmt.Println(err)
// 	}
// 	fmt.Println(phone)

// }

func Test_dataload(){
	repository.
	CheckUserByPhone("9744703558")
}