package main

import (
	"context"
	"fmt"
	"log"

	nat "github.com/libp2p/go-libp2p-nat"
)

func main() {
	nat, err := nat.DiscoverNAT(context.Background())
	//nat, err :=nat .DiscoverGateway()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	for _, mapItem := range nat.Mappings() {
		fmt.Printf("mapItem is %v\n", mapItem)
	}
}
