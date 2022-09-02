package case4

import (
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var moduleEtcdDefault = newModuleEtcd()

type EtcdConfig struct {
	Endpoints          []string
	UserName, Password string
	Timeout            time.Duration
}

type moduleEtcd struct {
	identify string
	Config   *EtcdConfig
	Client   *clientv3.Client
}

func newModuleEtcd() *moduleEtcd {
	return &moduleEtcd{
		identify: "{etcd}",
	}
}

func (m *moduleEtcd) load(config *EtcdConfig) (err error) {
	tagmsg := m.identify + "加载"

	err = m.getEtcdConfig(config)
	if err != nil {
		log.Panicln(tagmsg, "获取etcd失败", m.Config.Endpoints, m.Config.UserName, m.Config.Password, err)
	}

	retry := 0
	for {
		err = m.getEtcdClient()
		if err != nil {
			retry++
			if retry > 3 {
				log.Panicln(tagmsg, "连接失败.结束重试", m.Config.Endpoints, m.Config.UserName, m.Config.Password, err)
				break
			}
			log.Println(tagmsg, "连接失败.3秒后重试", m.Config.Endpoints, m.Config.UserName, m.Config.Password, err)
			time.Sleep(time.Second * 3)
			return
		}
		log.Println(tagmsg, "连接成功", m.Config.Endpoints, m.Config.UserName, m.Config.Password)
		break
	}
	return
}

func (m *moduleEtcd) getEtcdConfig(config *EtcdConfig) (err error) {
	m.Config = config
	if m.Config == nil {
		m.Config = &EtcdConfig{}
	}
	if m.Config.Endpoints != nil && len(m.Config.Endpoints) > 0 {
		return
	}
	m.Config.Endpoints, m.Config.UserName, m.Config.Password, err = getEtcdConfig()
	if err != nil {
		return
	}
	return
}

func (m *moduleEtcd) getEtcdClient() (err error) {
	if m.Config.Timeout == 0 {
		m.Config.Timeout = EtcdClientTimeout
	}
	m.Client, err = getEtcdClient(m.Config.Endpoints, m.Config.UserName, m.Config.Password, m.Config.Timeout)
	if err != nil {
		err = fmt.Errorf("Etcd getEtcdClient [%v - %s - %s] err :%s", m.Config.Endpoints, m.Config.UserName, m.Config.Password, err)
		return
	}
	return
}
