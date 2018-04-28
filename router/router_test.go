package router

import (
  "testing"
  "encoding/json"
  "fmt"
)

type Route struct {
  Name string
  Path string
  Modal string
}

type Modal struct {
  Name string
  Routes []Route
}

type Contract struct {
  Modals []Modal
  Routes []Route
}

func TestRouter(t * testing.T) {
  jsonText := []byte( `{
    "modals": [
      {
        "name": "Account",
        "routes": [
          {
            "name": "login"
          }
        ]
      },
      {
        "name": "Home"
      }
    ],
    "routes": [
      {
        "modal": "Home",
        "name": "About"
      }
    ]
  }` )
  var jsonObject Contract
  json.Unmarshal(jsonText, &jsonObject)

  text, _ := json.Marshal(jsonObject)
  fmt.Println(string(text))
}
