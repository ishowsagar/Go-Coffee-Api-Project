package controllers

import (
	"encoding/json"
	"net/http"
	"server/helpers"
	"server/services"

	"github.com/go-chi/chi/v5"
)

var coffee services.Coffee

// Get all coffees
func GetAllCoffees(w http.ResponseWriter,r *http.Request) {
	all,err := coffee.GetAllCoffees()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJson(w,http.StatusOK,helpers.Envelop{"coffees":all})
}

// Get Coffee by id slug ({id}👇👇)
func GetCoffeeByID(w http.ResponseWriter,r *http.Request) {
	id := chi.URLParam(r,"id")
	coffee,err := coffee.GetCoffeeByID(id)

	// if caught error getting coffee
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// otherwise successfully got coffee by calling that func query --> send res back to client✅✅
	helpers.WriteJson(w,http.StatusOK,helpers.Envelop{"coffee":coffee})
}


//@PUT ---- Update Coffee by id slug ({id}👇👇)
func UpdateCoffeeByID(w http.ResponseWriter,r *http.Request) {
	id := chi.URLParam(r,"id")
	var coffeeUpdateInputVar services.Coffee
	err:= json.NewDecoder(r.Body).Decode(&coffeeUpdateInputVar)
	// if caught error decoding r.Body into coffeeupdateinputvar
	if err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	updatedCoffee,err:= coffee.UpdateCoffee(id,coffeeUpdateInputVar)
	// if caught error getting coffee
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// otherwise successfully got coffee by calling that func query --> send res back to client✅✅
	helpers.WriteJson(w,http.StatusOK,helpers.Envelop{"updated_coffee":updatedCoffee})
}

// creates coffee
func CreateCoffee(w http.ResponseWriter,r *http.Request) {
	var coffeeData services.Coffee
	//! what decoder doing is --> # decoding whats coming from req.Body and injecting into coffeeData
	err := json.NewDecoder(r.Body).Decode(&coffeeData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// at this point --> ✅✅successfully have decoded data into coffeeData instance which has all the required data
	
	// & Since we got coffee data, now we can pass to func that query it to createCoffee
	coffeeCreated,err := coffee.CreateCoffee(coffeeData) // ✅✅ creates coffee 
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// sending response back to client
	helpers.WriteJson(w,http.StatusOK,helpers.Envelop{"coffee":coffeeCreated})
}

func DeleteCoffee(w http.ResponseWriter,r *http.Request) {
	id := chi.URLParam(r,"id")
	err := coffee.DeleteCoffeeByID(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// sending response back to client
	helpers.WriteJson(w,http.StatusOK,helpers.Envelop{"status":"succesfully deleted🛑🛑!"})

}