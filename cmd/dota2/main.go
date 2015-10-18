package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/endophage/godota2/api"
	"os"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	key := os.Args[1]
	a := api.NewDota2Api(key)
	m, err := a.Matches(0)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	fmt.Printf("Matches: %v\n", m)
}
