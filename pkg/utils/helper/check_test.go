package helper

import (
	"fmt"
	"testing"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
)

// func TestCheck(t *testing.T) {
// 	result :=HashPassword("123")
// 	t.Logf("%s", result)
// }

func Test_config(t *testing.T) {
	config, _ := config.LoadConfig()
	t.Logf("%s", config)
	fmt.Println("---------",config)
}
