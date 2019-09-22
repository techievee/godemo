package user

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/techievee/godemo/models"
	"github.com/techievee/godemo/storage"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type UserModule struct {
	logger  *log.Logger
	backend *storage.RedisClient
}

type Result struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

func (u *UserModule) InsertUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.logger.Print("Error")
	}
	log.Println(string(body))

	fmt.Println(r.Body)
	usr := models.User{}

	err = json.Unmarshal([]byte(body), &usr)
	if err != nil {
		u.logger.Print(err)
	}

	fmt.Println(usr)
	if usr.Email != "" && usr.FirstName != "" && usr.LastName != "" && usr.MobileID != "" {
		usr.UUID = uuid.New().String()
		err := u.backend.HmSetUser(usr)
		if err != nil {
			u.logger.Print(err)
			fmt.Print(err)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	res := Result{

		Result: "Success",
		Error:  "",
	}

	json.NewEncoder(w).Encode(res)
	//Insert the data to the User

}

func (u *UserModule) GetAllUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain:charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	usr, err := u.backend.Hgetallmap()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error occured"))
		return
	}

	b, err := json.Marshal(usr)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	//fmt.Println(string(b))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(b))
}

func (u *UserModule) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		starttime := time.Now()
		defer u.logger.Printf("Request processed in %s", time.Now().Sub(starttime))
		next(w, r)

	}
}

func (u *UserModule) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/user/insert", u.InsertUser)
	mux.HandleFunc("/user/getall", u.GetAllUser)
}

func NewUserModule(log *log.Logger, backend *storage.RedisClient) *UserModule {
	return &UserModule{
		logger:  log,
		backend: backend,
	}
}
