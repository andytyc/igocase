package case1

import (
	"log"
	"os"
)

func Do() {
	DARWIN_ETCD_ENDPOINTS := os.Getenv("DARWIN_ETCD_ENDPOINTS")
	DARWIN_ETCD_AUTH := os.Getenv("DARWIN_ETCD_AUTH")
	log.Println("DARWIN_ETCD_ENDPOINTS:", DARWIN_ETCD_ENDPOINTS, "DARWIN_ETCD_AUTH:", DARWIN_ETCD_AUTH)

	PATH := os.Getenv("PATH")
	log.Println("PATH:", PATH)

	ALG_DIR := os.Getenv("ALG_DIR")
	log.Println("ALG_DIR:", ALG_DIR)

	RecognitionSDK_HOME := os.Getenv("RecognitionSDK_HOME")
	log.Println("RecognitionSDK_HOME:", RecognitionSDK_HOME)

	LD_LIBRARY_PATH := os.Getenv("LD_LIBRARY_PATH")
	log.Println("LD_LIBRARY_PATH:", LD_LIBRARY_PATH)
}
