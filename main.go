package main // import "github.com/hashicorp/vault"

import (
	"os"
	"fmt"
  "github.com/newrelic/go-agent" 	//import newrelic agent
	"github.com/hashicorp/vault/cli"
)

func main() {

	// get newrelic license key
	licence := os.Getenv("licence")

	// Create a newrelic config.
	config := newrelic.NewConfig("vault", licence)

	// newrelic not enabled if there is no licence
	if ("" != licence) {
		config.Enabled = true
		fmt.Println("New relic will be enabled")
	} else {
		config.Enabled = false
		fmt.Println("New relic will be disabled as no licence key is provided.")
	}

	// Log newrelic to stdout
	config.Logger = newrelic.NewLogger(os.Stdout)

	// Create the newrelic application
  _, err := newrelic.NewApplication(config)

  // exit if an error has occurred in creating the newrelic application
	if err != nil {
			fmt.Println("An error has occurred starting New relic. Exiting Vault.")
			fmt.Println(err)
			os.Exit(1)
	}

	os.Exit(cli.Run(os.Args[1:]))
}
