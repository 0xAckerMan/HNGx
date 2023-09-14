package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
    "errors"

	"github.com/0xAckerMan/HNGx/0x01-people/cmd/data"
)


func (app *Application) healthcheck(w http.ResponseWriter, r *http.Request){
    data := map[string]string {
        "status": "Active",
        "Environment": app.Env,
        "Version": Version,
    }
    js, err := json.Marshal(data)
    if err != nil{
        http.Error(w, http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
        return
    }
    js = append(js, '\n')
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)

}

func (app *Application) getCreateUser(w http.ResponseWriter, r *http.Request){
    if r.Method == http.MethodGet{
        user :=[]data.Information{
            {
                Id: 1,
                Name: "Kores",
            },
            {
                Id: 2,
                Name: "Hopp",
            },
        }
        jsonUser, err:= json.Marshal(user) 
        if err != nil{
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }
        jsonUser = append(jsonUser, '\n')
        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonUser)
    }
    if r.Method == http.MethodPost{
        var input struct{
            Name string `json:"name"`
        }
        err := app.readJson(w, r, &input)
        if err != nil{
            http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        }
        user:= &data.Information{
            Name: input.Name,
        }

        err = app.UserModel.Insert(user)
        if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		headers := make(http.Header)
		headers.Set("Location", fmt.Sprintf("api/%d", user.Id))

        err = app.writeJson(w, http.StatusCreated, envelope{"user": user}, headers)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
        fmt.Fprintf(w, "%v\n", input)
    }
}

func (app *Application) UserGetUpdateDelete (w http.ResponseWriter, r *http.Request){
    switch r.Method {
    case http.MethodGet:
        app.getUser(w, r)
    case http.MethodPost:
        app.createUser(w, r)
    case http.MethodPut:
        app.updateUser(w, r)
    case http.MethodDelete:
        app.deleteUser(w, r)
    default:
        http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
    }

}

func (app *Application) getUser(w http.ResponseWriter, r *http.Request){
    id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user, err := app.UserModel.Get(idInt)
	if err != nil {
		switch {
		case errors.Is(err, errors.New("record not found")):
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	if err := app.writeJson(w, http.StatusOK, envelope{"user": user}, nil); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
    // id := r.URL.Path[len("/api/"):]
    // idInt, err := strconv.ParseInt(id, 10, 64)
    // if err != nil{
    //     http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    // }
    // user := data.Information{
    //     Id: idInt,
    //     Name: "Joel",
    // }

    // if err := app.writeJson(w, http.StatusOK, user); err != nil{
    //     http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    // }
}

func (app *Application) createUser(w http.ResponseWriter, r *http.Request){

}

func (app *Application) updateUser(w http.ResponseWriter, r *http.Request){

}

func (app *Application) deleteUser(w http.ResponseWriter, r *http.Request){

}
