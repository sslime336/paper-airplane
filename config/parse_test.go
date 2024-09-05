package config

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	appconf := ParseConfig[App]("../misc/config.example.yaml")
	fmt.Println(appconf)
}
