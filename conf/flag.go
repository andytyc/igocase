package conf

import (
	"log"
	"strconv"
)

var FlagCaseNum = newFlagTypeCaseNum()

func FlagCheck() {
	FlagCaseNum.check()
}

/*
************************************************************/

const (
	FlagCaseNum0 = iota
	FlagCaseNum1
	FlagCaseNum2
	FlagCaseNum3
	FlagCaseNum4
)

var FlagCaseNumUsageMap = map[int]string{
	FlagCaseNum0: "Hello World !",
	FlagCaseNum1: "获取环境变量",
	FlagCaseNum2: "执行命令",
	FlagCaseNum3: "godotenv",
	FlagCaseNum4: "etcd",
}

type FlagTypeCaseNum struct {
	Val        int
	Key        string
	ValDefault int
}

func newFlagTypeCaseNum() *FlagTypeCaseNum {
	return &FlagTypeCaseNum{
		Key:        "case",
		ValDefault: 0,
	}
}

func (f *FlagTypeCaseNum) help() string {
	help := "例子编号,选项参考:"
	for key, usage := range FlagCaseNumUsageMap {
		help += "\n  " + strconv.Itoa(key) + " :" + usage
	}
	return help
}

func (f *FlagTypeCaseNum) Usage() string {
	return f.help()
}

func (f *FlagTypeCaseNum) result() (string, bool) {
	usage, ok := FlagCaseNumUsageMap[f.Val]
	if !ok {
		usage = "不合法," + f.help()
	}
	return "参数 [" + f.Key + " " + strconv.Itoa(f.Val) + "] " + usage, ok
}

func (f *FlagTypeCaseNum) String() string {
	data, _ := f.result()
	return data
}

func (f *FlagTypeCaseNum) check() {
	data, ok := f.result()
	if ok {
		log.Println(data)
	} else {
		log.Panicln(data)
	}
}
