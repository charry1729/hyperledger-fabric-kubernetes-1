package users

// import(
// 	"github.com/charry1729/hyperledger-fabric-kubernetes-1/t3-services/back-end/models"
// )

func Destroy(id string) error  {
	//var err error
	var exists bool

	for index ,user := range mockUsers{
		if user.ID == id{
			mockUser  = append(mockUsers[:index],mockUsers[index+1:]...)
			exists = true
		}
	}

	if !exists{
		return errors.New("unable to delete user was not found")

	}
	return nil
}