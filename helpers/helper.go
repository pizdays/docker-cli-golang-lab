package helpers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tidwall/pretty"
)

func GetSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func PrintlnJson(v ...interface{}) {
	bytes, _ := json.MarshalIndent(v, "", "\t")
	fmt.Println(string(pretty.Color(pretty.PrettyOptions(bytes, pretty.DefaultOptions), nil)))
}
