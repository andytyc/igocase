package case4

import (
	"log"

	"github.com/joho/godotenv"
)

func Do() {
	if err := godotenv.Load(".env-etcd.sh"); err != nil {
		log.Panicln("No env file found {注意: 配置文件在程序运行目录下} :", err)
	}
	etcdDefault.load(&EtcdConfig{})

	do_lease_0()
	do_txn_0()
}
