package letter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllForUser(c *Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := mux.Vars(r)["user"]
		letters, err := c.Datastore.FetchAllForUser(user)
		if err != nil {
			fmt.Println("error encountered fetching letters: ", err.Error())
		}

		fmt.Printf("found letters: %+v", letters)

		b, err := json.Marshal(letters)
		if err != nil {
			fmt.Println(err)
		}

		w.Write(b)
	}
}

func GetLetterById(c *Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the letter id to retrieve
		lid := mux.Vars(r)["id"]

		letter, err := c.Datastore.FetchOne(lid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			resp := map[string]string{
				"id":      lid,
				"msg":     "Unable to fetch letter",
				"details": err.Error(),
				"level":   "error",
			}
			b, _ := json.Marshal(resp)
			w.Write(b)
			return
		}

		b, _ := json.Marshal(letter)
		w.Write(b)
	}
}

func InsertLetter(c *Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// decode the body to get the letter to insert
		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var l Letter
		err = json.Unmarshal(b, &l)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := c.Datastore.Insert(l)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err = json.Marshal(map[string]interface{}{"id": id})
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}
