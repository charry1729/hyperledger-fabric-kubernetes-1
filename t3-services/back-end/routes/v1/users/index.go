package users

import (
	"encoding/json"
	"net/http"
	UsersModel "github.com/charry1729/hyperledger-fabric-kubernetes-1/t3-services/back-end/models/v1/users.go"
)

func Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		users,err := UserModel.Index{
	
		if err!= nil {
				http.Error(w,err.Error(), http.StatusInternalServerError)
		}
		packet,err := json.Marshal(users)
		if err!= nil {
			http.Error(w,err.Error(), http.StatusInternalServerError)
		}

		w.Write(packet)
			
		}
	}
}