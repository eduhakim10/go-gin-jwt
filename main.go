package main

import (
	"evendor.com/go/appevendor"
	"evendor.com/go/initializers"
)

func init() {
	initializers.NowDB()
	// appevendor.evendorapp()
}
func main() {
	//appevendor.evendorapp()
	appevendor.EvendorApp()

}
