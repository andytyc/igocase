package gcase

import (
	"github.com/andytyc/igocase/conf"
	"github.com/andytyc/igocase/gcase/case0"
	"github.com/andytyc/igocase/gcase/case1"
	"github.com/andytyc/igocase/gcase/case2"
	"github.com/andytyc/igocase/gcase/case3"
)

func Do() {
	switch conf.FlagCaseNum.Val {
	case conf.FlagCaseNum0:
		case0.Do()
	case conf.FlagCaseNum1:
		case1.Do()
	case conf.FlagCaseNum2:
		case2.Do()
	case conf.FlagCaseNum3:
		case3.Do()
	}
}
