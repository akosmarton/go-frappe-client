# go-frappe-client
Frappe Client in Golang

## Example
```golang
package main

import (
	"fmt"

	frappe "github.com/akosmarton/go-frappe-client"
)

func main() {
	client := &frappe.Client{
		URL:    "http://localhost:8000",
		Key:    "api key",
		Secret: "api secret",
	}

	doc, err := client.Get("DocType", "Document", nil)
	if err != nil {
		panic(err)
	}

	err = client.AddTag("DocType", doc, "MyTag")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc)
}
```

## Running Tests
```
# URL="https://demo.erpnext.com" KEY="..." SECRET="..."  go test
```
