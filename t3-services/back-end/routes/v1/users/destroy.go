package users

func Destroy()  {
	return func(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		id := vars["id"]
		if err:=UserModel.Destroy(id); err!= nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Success"))



	}
}