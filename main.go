package main

import (
	"context"
	"os"

	"github.com/alcarria/data-marketplace-chaincode-rest/controller"
	"github.com/alcarria/data-marketplace-chaincode-rest/rest"
	"github.com/alcarria/data-marketplace-chaincode-rest/utils"
)

func main() {
	ctx := context.Background()
	logger := utils.CreateLogger("data-marketplace-chaincode-rest")

	config, err := utils.LoadConfig()
	if err != nil {
		logger.Printf("no-port-specified-using-default-9090")
		config.Port = 9090
	}

	fabricSetup := controller.FabricSetup{
		// Network parameters
		//OrdererID:  "orderer.example.com",
		//OrdererURL: "blockchain-orderer:31010",
		OrdererID:  "orderer.example.com",
		OrdererURL: "orderer.example.com:7050",

		// Channel parameters
		//ChannelID:     "dmp",
		//ChannelConfig: "/shared/dmp.tx",
		ChannelID:     "mychannel",
		ChannelConfig: "/channel-artifacts/channel.tx",
		

		// Chaincode parameters
		//ChainCodeID:     "dmp",
		//ChaincodeGoPath: os.Getenv("GOPATH"),
		//ChaincodePath:   "github.com/alcarria/data-marketplace-chaincode",
		//OrgAdmin:        "Admin",
		//OrgName:         "Org1",
		//ConfigFile:      "/shared/artifacts/config.yaml",
		//UserName:        "Admin",
		ChainCodeID:     "data-marketplace-chaincode",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/alcarria/data-marketplace-chaincode",
		OrgAdmin:        "Admin",
		OrgName:         "Org1",
		ConfigFile:      "/shared/artifacts/config.yaml",
		UserName:        "Admin",
	}

	// Initialization of the Fabric SDK from the previously set properties
	// err = fSDKSetup.Initialize()
	// if err != nil {
	// 	fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
	// 	panic("Unable to initialize the Fabric SDK")
	// }

	controller := controller.NewPeerController(ctx, logger, fabricSetup)
	handler := rest.NewCCHandler(ctx, logger, controller)
	server := rest.NewCCServer(ctx, logger, handler, config)
	logger.Printf("starting-server", server.Start())
}
