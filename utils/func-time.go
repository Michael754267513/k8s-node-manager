/*
		Handpay ServiceMesh

           创建时间: 2020年11月25日15:55:24

	       少侠好武功,一起Giao起来
	  	 我说一Giao,你说Giao
		   一 Giao ？？？？

*/

package utils

import (
	"time"

	"gitlab.handpay.com.cn/it/k8s-node-manager/global"
)

func Tracker() func() {
	start_time := time.Now()
	return func() {
		tracker_time := time.Since(start_time)
		global.Log.Info("本次日志归档时间:", tracker_time)
	}
}
