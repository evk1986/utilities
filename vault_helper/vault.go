package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/hashicorp/vault/api"
	"net/http"
	"time"
)

var token = flag.String("token", "", "vault token for auth")
var vaultAddr = flag.String("addr", "", "vault addr")
var path = flag.String("secret_path", "", "secret path")


var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func main() {
	flag.Parse()
	//var authPath = "/v1/auth/github/login"
	tls := &api.TLSConfig{Insecure: true}
	config := &api.Config{
		Address: *vaultAddr,
	}

	_ = config.ConfigureTLS(tls)

	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	client.SetToken(*token)
	secret, err := client.Logical().Read(*path)
	if err != nil {
		fmt.Println(err)
		return
	}
	m, ok := secret.Data["data"]
	if !ok {
		fmt.Printf("%T %#v\n", secret.Data["data"], secret.Data["data"])
		return
	}

	rawSecret, err := base64.StdEncoding.DecodeString(m.(string))

	fmt.Print(string(rawSecret))
}
