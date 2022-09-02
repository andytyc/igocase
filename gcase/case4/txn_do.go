package case4

import (
	"context"
	"log"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func do_txn_0() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	key := "/test1"
	val_no := "-1"
	val_yes := "1"

	client := etcdDefault.Client
	txn := client.Txn(ctx)
	txn = txn.If(clientv3.Compare(clientv3.Value(key), "=", val_no))
	txn = txn.Then(clientv3.OpPut(key, val_yes))
	txn = txn.Else(clientv3.OpGet(key, clientv3.WithPrefix()))
	_, err := txn.Commit()
	if err != nil {
		log.Fatalln("err :", err)
	}
	log.Println("ok")
}
