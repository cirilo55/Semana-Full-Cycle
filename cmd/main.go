package main

import (
	"github.com/codeedu/imersao/codepix-go/Infrastructure/db"
	"github.com/codeedu/imersao/codepix-go/application/grpc"
	"github.com/jinzhu/gorm"
	"os"
)

var database *gorm.DB

func main(){
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
