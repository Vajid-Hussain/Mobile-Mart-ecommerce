package helper

import (
	"testing"
)

// type forTest struct{
// 	collection interfaces.IUserRepo
// }
// var functions forTest

// func Setup(repo interfaces.IUserRepo){
// 	functions.collection=repo
// }

// func TestCheck(t *testing.T) {
// 	result := HashPassword("9876")
// 	t.Logf("%s --------------", result)
// 	fmt.Println(result)
// 	fmt.Println([]byte(result))
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

// func Test_dataload(t *testing.T){
// 	err:=functions.collection.CheckUserByPhone("9744703558")
// 	if err!=nil{
// 		fmt.Println(err ,"something is wrong")
// 	}
// }

func Test_CompairPassword(t *testing.T) {
	err := CompairPassword("9876", "$2a$10$vs/ljK7YN.s..kq/.MJX8egzqG8aO5XFKZFq6wmIcQLJrMHgUTp9O")
	t.Logf(err.Error())
}
