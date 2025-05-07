package main

import (
	"context"
	"log"
	"time"

	"github.com/docker-cli-golang-lab/logs"
	"github.com/docker-cli-golang-lab/routes"
	"github.com/docker/docker/client"

	// _ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// var err error

func main() {

	initTimeZone()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	logs.InitLogger()
	logs.InitLoggerRequest()
	defer logs.CloseLogReq()
	defer logs.Close()

	// 1. Create Docker Client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Failed to create Docker client: %v", err)
	}
	defer cli.Close()

	_, err = cli.Ping(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to Docker daemon: %v", err)
	}
	log.Println("Connected to Docker daemon.")

	// DB_HOST := os.Getenv("DB_HOST")
	// DB_DATABASE := os.Getenv("DB_DATABASE")
	// DB_USERNAME := os.Getenv("DB_USERNAME")
	// DB_PASSWORD := os.Getenv("DB_PASSWORD")
	// DB_PORT := os.Getenv("DB_PORT")

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
	// 	DB_HOST, DB_USERNAME, DB_PASSWORD, DB_DATABASE, DB_PORT)

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	logs.Error(err)
	// 	panic(err)
	// }

	// databases.DB = db

	// if err != nil {
	// 	fmt.Println("statuse: ", err)
	// }

	// sqlDB, err := databases.DB.DB()
	// if err != nil {
	// 	log.Fatal("Error getting sqlDB from GORM: ", err)
	// }
	// defer sqlDB.Close()

	// uri := os.Getenv("MONGODB_URI")
	// serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// databases.MongoDb, err = mongo.Connect(context.TODO(), opts)
	// if err != nil {
	// 	panic(err)
	// }

	// defer databases.MongoDb.Disconnect(context.TODO())

	//setup routes
	r := routes.SetupRouter()
	// running
	r.Run(":8080")
}

func initTimeZone() {
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = location
}
