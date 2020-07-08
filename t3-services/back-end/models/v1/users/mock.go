package users

import(
	"github.com/charry1729/hyperledger-fabric-kubernetes-1/t3-services/back-end/models"
)

var mockUsers models.users
func init(){
	usr, _ := models.NewUser("Nick","Kds","chary@gmial.com,","123")

	mockUsers = models.Users{
		*usr,
	}
}