package users
import(
	"github.com/charry1729/hyperledger-fabric-kubernetes-1/t3-services/back-end/models"
)

func Index()(users *models.Users,err error){
	users = &mockUsers
	return
}

