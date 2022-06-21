/*
		Handpay ServiceMesh

           创建时间: 2020年11月25日15:55:24

	       少侠好武功,一起Giao起来
	  	 我说一Giao,你说Giao
		   一 Giao ？？？？

*/

package main

import (
	"fmt"
	"os"
)

func main() {

	src := "./logs/risk-air01"
	dest := "./logs/archive"

	err := os.Rename(src, dest)
	fmt.Println(err)
}
