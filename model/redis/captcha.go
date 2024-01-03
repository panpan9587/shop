package redis

import "time"

func AddCaptcha(mobile, code string) (err error) {
	err = Client.Set(mobile, code, time.Second*60).Err()
	return
}
func GetCaptcha(mobile string) string {
	return Client.Get(mobile).Val()
}
