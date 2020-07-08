package users
import (
	"net/http"
	"encoding/json"
	"github.com/charry1729/hyperledger-fabric-kubernetes-1/t3-services/back-end/models"
	UsersModel "github.com/charry1729/hyperledger-fabric-kubernetes-1/t3-services/back-end/models/v1/users"
)
func Store() http.HandlerFunc {
	return func(w http.ResponseWriter , r *http.Request){
		var user models.User
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		err := decoder.Decode(&user)

		if err!= nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newUser,err := UsersModel.Store(user.FirstName,user.LastName,user.Email,user.Password)
		if err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}

		packet,err := json.Marshal(newUser)
		if err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		w.Write(packet)
	}
}