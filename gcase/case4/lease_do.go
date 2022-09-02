package case4

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func do_lease_0() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	key := "/test0"
	val := "c4677guisy7666s82usse233"

	leaseKeepTime := 5 // 5s
	var leaseid clientv3.LeaseID

	resp, err := etcdDefault.Client.Grant(ctx, int64(leaseKeepTime))
	if err != nil {
		log.Fatalln("err :", err)
	}
	leaseid = resp.ID
	log.Println("leaseid :", leaseid)

	_, err = etcdDefault.Client.Put(ctx, key, val, clientv3.WithLease(leaseid), clientv3.WithPrevKV())
	if err != nil {
		log.Fatalln("err :", err)
	}

	keepChan, err := etcdDefault.Client.KeepAlive(ctx, leaseid)
	if err != nil {
		log.Fatalln("err :", err)
	}

	subCtx, subCancel := context.WithDeadline(ctx, time.Now().Add(time.Second*20))
	defer subCancel()

	for {
		select {
		case <-subCtx.Done():
			log.Println("sub ctx Done", subCtx.Err())
			return
		case <-ctx.Done():
			log.Fatalln("ctx Done", ctx.Err())
		case <-etcdDefault.Client.Ctx().Done():
			log.Fatalln("client Done", etcdDefault.Client.Ctx().Err())
		case _, ok := <-keepChan:
			if !ok {
				log.Fatalln("续租失败")
			} else {
				log.Println("续租成功")
			}
		case <-time.After(time.Second * time.Duration(leaseKeepTime*2)):
			log.Fatalln("续租超时")
		}
	}
}
