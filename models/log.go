/*
		Handpay ServiceMesh

           创建时间: 2020年11月25日15:55:24

	       少侠好武功,一起Giao起来
	  	 我说一Giao,你说Giao
		   一 Giao ？？？？

*/

package models

type PodLogPath struct {
	Namespace string      `json:"namespace"` // 命名空间
	SubSystem []SubSystem `json:"sub_system"`
}
type SubSystem struct {
	SubSystem string   `json:"sub_system"`
	Pods      []string `json:"pods"`
}
