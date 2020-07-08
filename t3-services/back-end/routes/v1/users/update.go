package users


import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/charry1729/hyperledger-fabric-kubernetes-1/t3-services/back-end/models"
	UsersModel "github.com/charry1729/hyperledger-fabric-kubernetes-1/t3-services/back-end/models/v1/users"
)

func Update() http.HandlerFunc{

	return func (w http.ResponseWriter, r *http.Request){

		vars := mux.Vars(r)
		id := vars["id"]
		// id,err := strconv.Atoi(vars["id"])
		// if err!=nil{
		// 	http.Error(w, "no ID found in the URL")
		// 	return
		// }
		var opts UsersModel.UpdateOpts
		log.Println(opts)
		var user models.User
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		err := decoder.Decode(&user)
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if r.Method == "PUT"{
			opts.Replace = true
		}

		updatedUser, err := UserModel.Update(id.&user)
		if err!= nil {
			http.Error(w.Error(),http.StatusInternalServerError)
			return
		}
		packet,err := json.Marshal(updatedUser)
		if err!= nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		w.Write(packet)


	}
}