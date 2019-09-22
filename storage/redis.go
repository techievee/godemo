package storage

import (
	"github.com/fatih/structs"
	"github.com/go-redis/redis"
	"github.com/techievee/godemo/models"
	"log"
	"strconv"
)

/*----------------------Standard Redis methods---------------------------------  */
type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(endpoint string, password string, db string) *RedisClient {
	DbNumber, _ := strconv.Atoi(db)
	client := redis.NewClient(&redis.Options{
		Addr:     endpoint,
		DB:       DbNumber,
		Password: password,
	})
	return &RedisClient{client: client}
}

func (r *RedisClient) Client() *redis.Client {
	return r.client
}

func (r *RedisClient) Check() (string, error) {
	return r.client.Ping().Result()
}

func (r *RedisClient) BgSave() (string, error) {
	return r.client.BgSave().Result()
}

/*-----------------------------------------------------*/

//Get the user and insert into the Hashmap of Redis, using the HMset function
//Perform the Bgsave immediately to persist

func (r *RedisClient) HmSetUser(usr models.User) error {

	usrM := structs.Map(usr)
	err := r.client.HMSet("user:"+usr.UUID, usrM).Err()
	if err != nil {
		log.Printf("Redis: Error: %v", err)
	}
	r.client.BgSave()

	return nil
}

//Get all the data from the User:* key hashmap and return as slice of structure
func (r *RedisClient) Hgetallmap() ([]models.User, error) {

	var cursor uint64
	var keys []string
	var err error
	var data []models.User
	var n int
	for {

		keys, cursor, err = r.client.Scan(cursor, "user:*", 1000).Result()
		if err != nil {
			panic(err)
		}
		n += len(keys)
		for _, value := range keys {
			//fmt.Println(key,value)
			usr := models.User{}
			m, err := r.client.HGetAll(value).Result()
			if err != nil {
				continue
			}

			for col, val := range m {
				switch col {
				case "UUID":
					usr.UUID = val
				case "MobileID":
					usr.MobileID = val
				case "Email":
					usr.Email = val
				case "FirstName":
					usr.FirstName = val
				case "LastName":
					usr.LastName = val
				}
			}
			data = append(data, usr)
		}

		if cursor == 0 {
			break
		}
	}

	return data, err

}
