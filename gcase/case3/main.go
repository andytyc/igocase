package case3

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Do() {
	// 默认文件名称: .env
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found {注意: 配置文件在程序运行目录下} :", err)
	}
	// 指定一个文件
	// if err := godotenv.Load(".env-mongodb"); err != nil {
	// 	log.Fatalln("No .env-mongodb file found {注意: 配置文件在程序运行目录下} :", err)
	// }
	// 指定多个文件
	// if err := godotenv.Load(".env-mongodb", ".env-etcd"); err != nil {
	// 	log.Fatalln("No .env-mongodb file found {注意: 配置文件在程序运行目录下} :", err)
	// }

	MONGODB_URI := os.Getenv("MONGODB_URI")
	if MONGODB_URI == "" {
		log.Fatalln("You must set your 'MONGODB_URI' environmental variable")
	}
	log.Println("MONGODB_URI :", MONGODB_URI)

	Endpoints := os.Getenv("Endpoints")
	if Endpoints == "" {
		log.Fatalln("You must set your 'Endpoints' environmental variable")
	}
	log.Println("Endpoints :", Endpoints)

	UserName := os.Getenv("UserName")
	if UserName == "" {
		log.Fatalln("You must set your 'UserName' environmental variable")
	}
	log.Println("UserName :", UserName)

	Password := os.Getenv("Password")
	if Password == "" {
		log.Fatalln("You must set your 'Password' environmental variable")
	}
	log.Println("Password :", Password)
}
