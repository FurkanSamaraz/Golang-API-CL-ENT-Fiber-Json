package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Time struct {
	Updated    string
	UpdatedISO string
	Updateduk  string
}
type Cru struct {
	Code       string
	Symbol     string
	Rate       string
	Rate_Float float64
}
type Bpi struct {
	USD Cru
	EUR Cru
	GBP Cru
}

type ApiRes struct {
	Time      Time
	ChartName string
	Bpi       Bpi
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		resp, _ := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
		data, _ := ioutil.ReadAll(resp.Body)
		var api ApiRes
		json.Unmarshal(data, &api)
		veri, _ := json.Marshal(api.Bpi)

		return c.Format(veri)
	})
	app.Listen(":8080")

}
