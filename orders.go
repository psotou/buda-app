package main

import (
	"fmt"
	"os"
	"src/buda-app/buda"
	//"strconv"
    //"strings"
)

func main() {
	var key string = os.Getenv("KEY")
	var secret string = os.Getenv("SECRET")
	var crypto string = os.Args[1] + "-clp"

	budaClient, err := buda.NewAPIClient(key, secret)

	if err != nil {
		panic(err)
	}

    orders, _ := budaClient.GetOrdersByMarket(crypto)

    fmt.Println(orders)
}