package case3

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Do() {
	if err := godotenv.Load(); err != nil {
		log.Panicln("No .env file found {注意: .env 文件在程序运行目录下}")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Panicln("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	log.Println("MONGODB_URI :", uri)
}
