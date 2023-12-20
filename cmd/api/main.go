package main

import (
	"fmt"
	"log"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/docs"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/di"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("/home/vajid/Brocamp/Mobile-mart/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("**workded **")

	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("##", err)
		log.Fatal("error at loading the env file using viper")
	}
	fmt.Println("**config done")

	//	@title						Go + Gin Mobile-Mart
	//	@description				Mobile Mart - Your Ultimate Mobile Phone Shopping API 📱🛒
	//	@contact.name				API Support
	//	@securityDefinitions.apikey	BearerTokenAuth
	//	@in							header
	//	@name						Authorization
	//	@securityDefinitions.apikey	Refreshtoken
	//	@in							header
	//	@name						Refreshtoken
	//	@host						localhost:8080
	//	@BasePath					/
	//	@query.collection.format	multi
	// docs.SwaggerInfo.Version = "1.0"

	docs.SwaggerInfo.Title = "Mobile_mart"
	docs.SwaggerInfo.Host = "mobilesmart.vajid.tech"

	server, err := di.InitializeAPI(config)
	if err != nil {
		log.Fatal("error for server creation")
	}

	server.Start()
}

//	@securityDefinitions.apikey	Refreshtoken
//	@in							header
//	@name						Refreshtoken
