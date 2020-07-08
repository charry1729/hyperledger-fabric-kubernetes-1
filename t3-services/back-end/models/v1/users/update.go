package users

import(
	"github.com/charry1729/hyperledger-fabric-kubernetes-1/t3-services/back-end/models"
)

type UpdateOpts struct{
	Replace bool
}


func Update(id string,usr *models.User,opts *UpdateOpts)(*models.User,error)  {
	//var err error

	if opts == nil {
		opts = &UpdateOpts{}
	}
	var exists bool

	for index ,user := range mockUsers{
		if user.ID == id{

			if opts.Replace {
				usr.ID = mockUsers[index].ID
				usr.Password = mockUsers[index].Password
				usr.Email = mockUsers[index].Email
				mockUsers[index] = *usr	

			}else {
				if len(usr.FirstName)>0{
					mockUsers[index].FirstName = usr.FirstName
				}
				if len(usr.LastName)>0{
					mockUsers[index].LastName = usr.LastName
				}

			}
			//existingID := mockUsers[index].id
			//existingPassword := mockUsers[index].Password

			exists = true
			usr= &mockUsers[index]
		}
	}

	if !exists{
		return nil,errors.New("unable to update user was not found")

	}
	return mockUsers[index],nil
}