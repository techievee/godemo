package storage

import (
	"encoding/json"
	"fmt"
	guuid "github.com/google/uuid"
	"github.com/techievee/godemo/models"
	"os"
	"testing"
)

var r *RedisClient

//Add test in front ok key for testing
const prefix = "test"

func TestMain(m *testing.M) {
	r = NewRedisClient("localhost:6379", "", "0")
	reset()
	c := m.Run()
	reset()
	os.Exit(c)
}

//Mock insertion of user using the written function
func TestHmSet(t *testing.T) {
	usr := models.User{
		UUID:      guuid.New().String(),
		MobileID:  "123456798",
		Email:     "vinod@email.com",
		FirstName: "Vinod",
		LastName:  "Kumar",
	}

	err := r.HmSetUser(usr)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	t.Logf("Success")
}

//Mock reading of user details using the written function
func TestHgetAll(t *testing.T) {

	usr, err := r.Hgetallmap()
	if err != nil {
		t.Errorf("Error: %v ", err)
	}

	b, err := json.Marshal(usr)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Println(string(b))
	t.Logf("Result=%v,Success", b)
}

//Delete all the DB with test prefix
func reset() {
	keys := r.client.Keys("test" + ":*").Val()
	for _, k := range keys {
		r.client.Del(k)
	}
}
