package case4

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

func getEtcdConfig() (endpoints []string, userName string, passwd string, err error) {
	etcds := os.Getenv(ETCD_ENDPOINTS_ENV)
	if etcds == "" {
		err = fmt.Errorf("etcd endpoints should be config first in env as %s", ETCD_ENDPOINTS_ENV)
		return
	}
	endpoints = strings.Split(etcds, delimiter)

	auth := os.Getenv(ETCD_AUTH_ENV)
	if auth == "" {
		log.Println("etcd", etcds, "未配置权限认证信息")
		return
	}
	log.Println("etcd", etcds, "配置权限认证信息", auth)

	funcEtcdAuthInfo := func(auth string) (name string, passwd string, err error) {
		up := strings.Split(auth, authDelimiter)
		if len(up) != 2 {
			err = fmt.Errorf("etcd authorization info %s invalid", auth)
			return
		}

		name = up[0]
		passwd = up[1]
		if name == "" {
			err = fmt.Errorf("etcd authorization username invalid")
			return
		}
		return
	}

	userName, passwd, err = funcEtcdAuthInfo(auth)
	if err != nil {
		return
	}
	return
}

func getEtcdClient(endpoints []string, userName, passwd string, timeout time.Duration) (cli *clientv3.Client, err error) {
	// logConfig := zap.NewDevelopmentConfig()
	// logConfig.OutputPaths = append(logConfig.OutputPaths, "./etcdcout.log")
	// logConfig.ErrorOutputPaths = append(logConfig.ErrorOutputPaths, "./etcdcerr.log")
	return clientv3.New(clientv3.Config{
		Endpoints:            endpoints,
		Username:             userName,
		Password:             passwd,
		DialTimeout:          timeout * time.Second,
		DialKeepAliveTimeout: timeout * time.Second,
		DialOptions:          []grpc.DialOption{grpc.WithBlock()},
		// LogConfig:            &logConfig,
	})
}
