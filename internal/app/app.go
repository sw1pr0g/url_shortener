package app

import (
	"flag"
	"fmt"
	"github.com/sw1pr0g/url_shortener/config"
)

func Run(cfg *config.Config) {
	usePostgres := flag.Bool("d", false, "Use PostgreSQL for data storage")
	flag.Parse()

	if *usePostgres {
		fmt.Println("Using PostgreSQL for data storage")
	} else {
		fmt.Println("Using storage for data storage")
	}

	fmt.Println("App")
}
