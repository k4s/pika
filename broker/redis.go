package broker

import (
	"strconv"
	"strings"

	"gopkg.in/redis.v3"
)

func newRedisClient(addr string, pwd string, db int64) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr, //"localhost:6379"
		Password: pwd,  // no password set
		DB:       db,   // use default DB
	})
	return client
}

type PikaRedisClient struct {
	client *redis.Client
}

func NewPikaRedisClient(redisAddr string) *PikaRedisClient {
	//get db from redisAddr
	addrl := strings.Split(redisAddr, "/")
	db, _ := strconv.Atoi(addrl[len(addrl)-1])
	// if redisAddr contain the password
	if isExist := strings.Contains(redisAddr, "@"); isExist {
		pwd := strings.Split(strings.Split(redisAddr, "@")[0], "://")[1]
		addr := strings.Split(strings.Split(redisAddr, "@")[1], "/")[0]
		return &PikaRedisClient{
			client: newRedisClient(addr, pwd, int64(db)),
		}
	} else {
		addr := strings.Split(strings.Split(redisAddr, "://")[1], "/")[0]
		return &PikaRedisClient{
			client: newRedisClient(addr, "", int64(db)),
		}
	}

}

func (p *PikaRedisClient) Publish(topic string, pmsg string) error {
	err := p.client.Publish(topic, pmsg).Err()
	return err
}

func (p *PikaRedisClient) Subscribe(topic string) (*redis.PubSub, error) {
	pubsub, err := p.client.Subscribe(topic)
	return pubsub, err
}

func (p *PikaRedisClient) RPush(listName string, values ...string) error {
	err := p.client.RPush(listName, values...).Err()
	return err
}

func (p *PikaRedisClient) RPop(listName string) (string, error) {
	val, err := p.client.RPop(listName).Result()
	return val, err
}

func (p *PikaRedisClient) LPush(listName string, values ...string) error {
	err := p.client.LPush(listName, values...).Err()
	return err
}

func (p *PikaRedisClient) LPop(listName string) (string, error) {
	val, err := p.client.LPop(listName).Result()
	return val, err
}
