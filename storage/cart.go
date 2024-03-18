package storage

import (
	"encoding/json"
	"io/fs"
	"os"
)

type SSHConfig struct {
	Host string
	Port string
	User string
}

type Cart struct {
	Items []SSHConfig
}

func getHomeDir() string {
	str, _ := os.UserHomeDir()

	var fullDir = str + "/.config/ssh-cart/"

	_ = os.MkdirAll(fullDir, fs.ModePerm)

	return fullDir
}

func GetCart() Cart {
	homeDir := getHomeDir()
	directory := homeDir + "/ssh-cart.json"
	file, err := os.Open(directory)
	if err != nil {
		return Cart{}
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	cart := Cart{}
	err = decoder.Decode(&cart)
	if err != nil {
		return Cart{}
	}

	return cart
}

func SaveCart(cart Cart) {
	homeDir := getHomeDir()
	directory := homeDir + "/ssh-cart.json"
	file, err := os.Create(directory)
	if err != nil {
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cart)
	if err != nil {
		return
	}

	return

}
func AddItem(item SSHConfig) {
	cart := GetCart()
	cart.Items = append(cart.Items, item)
	SaveCart(cart)
}

func RemoveItem(item SSHConfig) {
	cart := GetCart()

	for i, v := range cart.Items {
		if v.Host == item.Host && v.Port == item.Port && v.User == item.User {
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			break
		}
	}
	SaveCart(cart)
}
