package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		err error

		districtScheme DistrictSchemeType
	)

	districtScheme, err = buildDistrictScheme(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(buildLazyPath(districtScheme))
}
