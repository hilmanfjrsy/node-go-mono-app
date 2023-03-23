package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

func SetCache(name string, data []byte, expiredDuration time.Duration) error {
	filename := "cache/" + name + "_cache.json"
	if _, err := os.Stat(filename); err != nil {
		d := fmt.Sprintf(`{"expiredAt": "%s","data": %s}`,
			time.Now().Add(expiredDuration).Format("2006-01-02 15:04:05"),
			string(data),
		)

		if os.IsNotExist(err) {
			err := os.WriteFile(filename, []byte(d), 0644)
			return err
		}
	}
	return nil
}

func GetCache(name string) ([]byte, error) {
	filename := "cache/" + name + "_cache.json"
	data := map[string]interface{}{}
	file, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		return []byte{}, err
	}

	expiredAt, err := time.Parse("2006-01-02 15:04:05", data["expiredAt"].(string))
	if err != nil {
		return []byte{}, err
	}
	if expiredAt.Before(time.Now()) {
		err := os.Remove(filename)
		if err != nil {
			return []byte{}, err
		}
		return []byte{}, errors.New("cache expired")
	}

	b, err := json.Marshal(data["data"])
	if err != nil {
		return []byte{}, err
	}

	return b, nil
}
