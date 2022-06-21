/*
		Handpay ServiceMesh

           创建时间: 2020年11月25日15:55:24

	       少侠好武功,一起Giao起来
	  	 我说一Giao,你说Giao
		   一 Giao ？？？？

*/

package main

import (
	"encoding/json"
	"fmt"

	archive_log "gitlab.handpay.com.cn/it/k8s-node-manager/pkg/archive-log"
)

func main() {
	pods := archive_log.GetNsPods("./logs")
	b, _ := json.Marshal(pods)
	fmt.Println(string(b))
}
