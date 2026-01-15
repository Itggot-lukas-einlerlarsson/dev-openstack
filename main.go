package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
)

func main() {
	// 1. Autentisering (Anpassa efter din DevStack-installation)
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "http://<DIN_DEVSTACK_IP>/identity/v3",
		Username:         "admin",
		Password:         "secret",
		TenantName:       "admin",
		DomainID:         "default",
	}

	provider, err := openstack.AuthenticatedClient(context.Background(), opts)
	if err != nil {
		log.Fatal("Kunde inte autentisera:", err)
	}

	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})

	// 2. Definiera servern
	// OBS: ID:n nedan är exempel. Du hittar rätt ID:n via CLI ( t.ex. 'openstack image list' )
	serverOptions := servers.CreateOpts{
		Name:      "Go-Hello-World",
		ImageRef:  "IMAGE_ID_HÄR", 
		FlavorRef: "FLAVOR_ID_HÄR",
		Networks: []servers.Network{
			{UUID: "NETWORK_ID_HÄR"},
		},
	}

	// 3. Skapa servern
	fmt.Println("Skapar server...")
	server, err := servers.Create(context.Background(), client, serverOptions).Extract()
	if err != nil {
		log.Fatal("Kunde inte skapa server:", err)
	}

	fmt.Printf("Server skapad! Namn: %s, ID: %s\n", server.Name, server.ID)
}ckage main
import(
	"fmt"
	// "github.com/gophercloud/gophercloud/v2"
)

func main(){
	fmt.Println("Smacknoodleon")
}
