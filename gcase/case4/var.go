package case4

import "time"

const (
	ETCD_ENDPOINTS_ENV = "DARWIN_ETCD_ENDPOINTS"
	ETCD_AUTH_ENV      = "DARWIN_ETCD_AUTH"

	delimiter     = ";"
	authDelimiter = ":"

	EtcdClientTimeout = 30 * time.Second
)
