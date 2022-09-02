package case2

import (
	"bytes"
	"log"
	"os/exec"
)

func Do() {
	cmd := exec.Command("go", "version")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	if err != nil {
		log.Println("cmd run err:", cmd.Args, err)
		return
	}
	log.Println("out :", string(stdout.Bytes()))
	log.Println("outerr :", string(stderr.Bytes()))
}
