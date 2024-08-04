package di

import (
	"log"

	"github.com/shivaraj-shanthaiah/godoc-user/config"
	"github.com/shivaraj-shanthaiah/godoc-user/pkg/db"
	"github.com/shivaraj-shanthaiah/godoc-user/pkg/doctors"
	"github.com/shivaraj-shanthaiah/godoc-user/pkg/handler"
	"github.com/shivaraj-shanthaiah/godoc-user/pkg/repo"
	"github.com/shivaraj-shanthaiah/godoc-user/pkg/server"
	"github.com/shivaraj-shanthaiah/godoc-user/pkg/services"
)

func Init() {
	config.LoadConfig()
	db := db.ConnectDB()
	client, err := doctors.ClientDial()
	if err != nil {
		log.Fatalf("something went wrong %v", err)
	}

	userRepo := repo.NewUserRepo(db)
	userSvc := services.NewUserService(userRepo, client)
	userHandler := handler.NewUserHandler(userSvc)
	server.NewgrpcServer(userHandler)

}
