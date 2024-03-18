package main

import (
	"ShoppingCart/storage"
	"ShoppingCart/ui"
	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
)

func main() {
	app := &cli.App{
		Name:  "Shopping cart",
		Usage: "Manage Manage your SSH connections in style",

		Authors: []*cli.Author{
			{
				Name:  "NietThijmen",
				Email: "thijmen@rierink.dev",
			},
		},

		Commands: []*cli.Command{
			{
				Name:  "add",
				Usage: "Add a new item to the cart",

				Action: func(c *cli.Context) error {
					username := ui.Input("Enter the username: ", 100)
					host := ui.Input("Enter the host: ", 100)
					port := ui.Input("Enter the port: ", 1000)

					log.Infof("Username: %s, Host: %s, Port: %s", username, host, port)

					storage.AddItem(storage.SSHConfig{
						Host: host,
						Port: port,
						User: username,
					})
					return nil
				},
			},

			{
				Name:  "remove",
				Usage: "Remove an item from the cart",

				Action: func(c *cli.Context) error {

					cart := storage.GetCart()
					options := make([]ui.SelectOption, len(cart.Items))
					for i, item := range cart.Items {
						options[i] = ui.SelectOption{
							Key:   strconv.FormatInt(int64(i), 10),
							Value: item.User + "@" + item.Host + ":" + item.Port,
						}
					}

					key := ui.Select(options)
					id, _ := strconv.Atoi(key)
					item := cart.Items[id]
					storage.RemoveItem(item)

					log.Infof("Removed: %s", item.User+"@"+item.Host+":"+item.Port)

					return nil
				},
			},

			{
				Name:  "connect",
				Usage: "Connect to an item in the cart",

				Action: func(c *cli.Context) error {

					cart := storage.GetCart()
					options := make([]ui.SelectOption, len(cart.Items))
					for i, item := range cart.Items {
						options[i] = ui.SelectOption{
							Key:   strconv.FormatInt(int64(i), 10),
							Value: item.User + "@" + item.Host + ":" + item.Port,
						}
					}

					key := ui.Select(options)
					id, _ := strconv.Atoi(key)
					item := cart.Items[id]

					log.Infof("Connecting to: %s", item.User+"@"+item.Host+":"+item.Port)
					connect(item)
					log.Info("Connection closed, shutting down...")

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
