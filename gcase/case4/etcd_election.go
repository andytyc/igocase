package case4

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

var (
	EtcdElectPreStr = "election"
)

type EtcdElection struct {
	ctx    context.Context
	cancel context.CancelFunc
	cli    *clientv3.Client
	name   string
	Leader chan string
	sess   *concurrency.Session
	elect  *concurrency.Election
}

func NewElection(name string, cli *clientv3.Client) (em *EtcdElection, err error) {
	name = strings.TrimLeft(name, "/")
	name = "/" + EtcdElectPreStr + "/" + name
	em = &EtcdElection{
		cli:    cli,
		name:   name,
		Leader: make(chan string),
	}
	em.ctx, em.cancel = context.WithCancel(context.TODO())
	return
}

func (e *EtcdElection) Resign() (err error) {
	if e.elect != nil {
		err = e.elect.Resign(e.ctx)
	}
	e.cancel()
	return
}

func (e *EtcdElection) Campaign(val string) (err error) {
	go func() {
		for {
			select {
			case <-e.ctx.Done():
				return
			default:
				ctx, cancel := context.WithCancel(e.ctx)
				e.sess, err = concurrency.NewSession(e.cli)
				if err != nil {
					err = fmt.Errorf("Campaign NewSession err :%s", err)
					log.Println("选举错误", e.name, val, err)
					cancel()
					continue
				}
				e.elect = concurrency.NewElection(e.sess, e.name)
				go func() {
				retry:
					o := e.elect.Observe(ctx)
					for {
						select {
						case <-ctx.Done():
							return
						case resp, ok := <-o:
							if !ok {
								err := errors.New("Observe get campaign response error")
								log.Println("选举错误", e.name, val, err)
								goto retry
							}
							e.Leader <- string(resp.Kvs[0].Value)
						}
					}
				}()
				if cerr := e.elect.Campaign(ctx, val); cerr != nil {
					err = fmt.Errorf("Campaign err :%s", cerr)
					log.Println("选举错误", e.name, val, err)
					cancel()
					continue
				}
				log.Println("选举成功", e.name, val)
				select {
				case <-ctx.Done():
				case <-e.sess.Done():
					err = errors.New("Session done")
					log.Println("选举错误", e.name, val, err)
					cancel()
					continue
				}
				cancel()
			}
		}
	}()
	return
}
