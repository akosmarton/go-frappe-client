package frappe

import (
	"log"
	"os"
	"testing"
)

func getClient() *Client {
	url := os.Getenv("URL")
	key := os.Getenv("KEY")
	secret := os.Getenv("SECRET")

	return &Client{
		URL:    url,
		Key:    key,
		Secret: secret,
	}
}

func TestClient_GetAll(t *testing.T) {
	client := getClient()

	var fields []string
	var filters []Filter

	_, err := client.GetAll("Lead", fields, filters, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_Post(t *testing.T) {
	client := getClient()

	doc := NewDocument()
	doc["lead_name"] = "Name"

	_, err := client.Post("Lead", doc)
	if err != nil {
		log.Fatal(err)
	}
}

func TestClient_Get(t *testing.T) {
	client := getClient()

	doc := NewDocument()
	doc["lead_name"] = "Name"

	r1, err := client.Post("Lead", doc)
	if err != nil {
		log.Fatal(err)
	}

	var filter []string
	r2, err := client.Get("Lead", r1.GetAsString("name"), filter)
	if err != nil {
		log.Fatal(err)
	}

	if r1.GetAsString("name") != r2.GetAsString("name") {
		log.Fatal("Lead mismatch")
	}
}

func TestClient_Put(t *testing.T) {
	client := getClient()

	doc := NewDocument()
	doc["lead_name"] = "Name"

	r1, err := client.Post("Lead", doc)
	if err != nil {
		log.Fatal(err)
	}

	r1.Set("lead_name", "NewName")
	r2, err := client.Put("Lead", r1.GetAsString("name"), r1)
	if err != nil {
		log.Fatal(err)
	}

	if r1.GetAsString("lead_name") != r2.GetAsString("lead_name") {
		log.Fatal("Lead name mismatch")
	}
}
