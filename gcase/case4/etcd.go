package case4

import (
	"fmt"
	"log"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdConfig struct {
	Endpoints          []string
	UserName, Password string
	Timeout            time.Duration
}

type moduleEtcd struct {
	identify        string
	Config          *EtcdConfig
	Client          *clientv3.Client
	ElectionMap     map[string]*EtcdElection
	lockElectionMap *sync.RWMutex
}

func newModuleEtcd() *moduleEtcd {
	return &moduleEtcd{
		identify:        "{etcd}",
		ElectionMap:     make(map[string]*EtcdElection, 0),
		lockElectionMap: &sync.RWMutex{},
	}
}

/* load
****************************************************************/

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

/* election
****************************************************************/

func (m *moduleEtcd) OnElection(name, val string) (err error) {
	var election *EtcdElection
	election, err = NewElection(name, m.Client)
	if err != nil {
		err = fmt.Errorf("NewElection err :%s", err)
		return
	}
	err = election.Campaign(val)
	if err != nil {
		err = fmt.Errorf("Campaign err :%s", err)
		return
	}
	m.lockElectionMap.Lock()
	m.ElectionMap[name] = election
	m.lockElectionMap.Unlock()
	return
}

func (m *moduleEtcd) GetElection(name string) (election *EtcdElection, ok bool) {
	m.lockElectionMap.RLock()
	defer m.lockElectionMap.RUnlock()
	election, ok = m.ElectionMap[name]
	return
}

func (m *moduleEtcd) ResignElection(name string, all bool) {
	m.lockElectionMap.Lock()
	defer m.lockElectionMap.Unlock()
	if all {
		for _, election := range m.ElectionMap {
			if election != nil {
				election.Resign()
			}
		}
		m.ElectionMap = make(map[string]*EtcdElection, 0)
	} else {
		election, ok := m.ElectionMap[name]
		if ok && election != nil {
			election.Resign()
		}
		delete(m.ElectionMap, name)
	}
	return
}

/*
****************************************************************/
