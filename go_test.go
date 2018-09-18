package test

import (
	"github.com/yangsai/blance/strategy"
	"testing"
)


func Benchmark_ys(t *testing.B) {
	sers := []strategy.Servers{
		{
			Service: "视频去重",
			Weight: 2,
			Provider: "ali",
		},
		{
			Service: "视频去重2",
			Weight: 2,
			Provider: "bd",
		},
	}
	b, err := strategy.NewBlance(sers)
	if err != nil {
		t.Fatalf("newblance error:%s", err.Error())
	}
	for i:=0; i<20; i++ {
		a := b.GetServer()
		t.Logf("测试通过%v \r\n", a)
	}
}
