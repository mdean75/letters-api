package letter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetMetaForLetter(c *Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			return
		}

		id := mux.Vars(r)["id"]
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		letter, err := c.Datastore.FetchMetaForLetter(id)
		if err != nil {
			fmt.Println("error encountered fetching letter: ", err.Error())
		}

		b, err := json.Marshal(letter)
		if err != nil {
			fmt.Println(err)
		}

		w.Write(b)
	}
}

func GetMetaForUser(c *Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			return
		}

		user := mux.Vars(r)["user"]
		if user == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		letters, err := c.Datastore.FetchMetaForUser(user)
		if err != nil {
			fmt.Println("error encountered fetching letters: ", err.Error())
		}

		b, err := json.Marshal(letters)
		if err != nil {
			fmt.Println(err)
		}

		w.Write(b)
	}
}

func GetAllForUser(c *Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			return
		}

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

func GetLetterContentById(c *Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			return
		}
		// get the letter id to retrieve
		lid := mux.Vars(r)["id"]

		i, err := c.Datastore.FetchContent(lid)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error fetching content"))
			return
		}

		b, _ := json.Marshal(map[string]interface{}{"content": i})
		w.Write(b)
	}
}

func GetLetterById(c *Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			return
		}
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

type Data struct {
	Content   string    `json:"content" bson:"content"`
	CreatedTs time.Time `json:"createdTs" bson:"createdTs"`
	To        string    `json:"to" bson:"to"`
	From      string    `json:"from" bson:"from"`
	Title     string    `json:"title" bson:"title"`
	User      string    `json:"user" bson:"user"`
}

func InsertLetterHTML(c *Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			return
		}

		var M Data

		b, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//var l Letter
		err = json.Unmarshal(b, &M)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if r.Method == http.MethodPost {
			fmt.Println("method post")
		}

		id, err := c.Datastore.InsertHTML(M)
		if err != nil {
			fmt.Println(err)
		}
		w.Write([]byte(fmt.Sprintf("data written to grid bucket: %s", id)))
	}
}

func InsertLetter(c *Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			return
		}

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
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		w.Write(b)
	}
}

func GetLoginRadiusUserDetails(c *Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			return
		}

		// get query param for current token
		token := r.URL.Query().Get("token")
		if token == "" {
			// send error response
			fmt.Println("no token received")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("no token received"))
			return
		}

		// make request to login radius to get the user details
		user, err := requestLoginRadiusUserDetails(c, token)
		if err != nil {
			fmt.Println(err)
			// respond with error
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		fmt.Println(user)

		var primaryEmail string
		for _, address := range user.Email {
			if address.Type == "Primary" {
				primaryEmail = address.Value
				break
			}
		}

		response := map[string]interface{}{"firstname": user.Firstname, "lastname": user.Lastname, "email": primaryEmail}

		b, err := json.Marshal(response)
		if err != nil {
			// now what
			fmt.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}

}

type UserProfile struct {
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Email     []email `json:"email"`
}

type email struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func requestLoginRadiusUserDetails(c *Controller, token string) (UserProfile, error) {
	//req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://api.loginradius.com/identity/v2/auth/account", nil)
	//if err != nil {
	//	return nil, err
	//}

	resp, err := http.Get(fmt.Sprintf("https://api.loginradius.com/identity/v2/auth/account?access_token=%s&apiKey=%s&fields=email,firstname,lastname", token, c.LoginRadiusApiKey))
	if err != nil {
		return UserProfile{}, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return UserProfile{}, err
	}

	var user UserProfile
	err = json.Unmarshal(b, &user)
	if err != nil {
		return UserProfile{}, err
	}

	return user, nil
}
