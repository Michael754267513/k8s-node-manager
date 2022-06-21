/*
		Handpay ServiceMesh

           创建时间: 2020年11月25日15:55:24

	       少侠好武功,一起Giao起来
	  	 我说一Giao,你说Giao
		   一 Giao ？？？？

*/

package getpods

import (
	"fmt"
	"io/ioutil"
	"os"

	"gitlab.handpay.com.cn/it/k8s-node-manager/models"
)

func GetNsPods(log_path string) (podlog []models.PodLogPath) {
	if !PathExists(log_path) {
		return podlog
	}
	filenames, err := ioutil.ReadDir(log_path)
	if err != nil {
	}
	for _, filename := range filenames {
		if filename.IsDir() {
			// 获取namespace
			podlog = append(podlog, models.PodLogPath{
				Namespace: getString(filename),
				SubSystem: GetSubSystems(log_path, getString(filename)),
			})
		}
	}
	return podlog
}

func GetSubSystems(log_path string, namespace string) (podlog []models.SubSystem) {
	filenames, err := ioutil.ReadDir(fmt.Sprint(log_path, "/", namespace))
	if err != nil {
	}
	for _, filename := range filenames {
		if filename.IsDir() {
			// 获取namespace
			podlog = append(podlog, GetSubSystem(log_path, namespace, getString(filename)))
		}
	}
	return podlog
}

func GetSubSystem(log_path string, namespace string, subsystem string) (sub models.SubSystem) {
	filenames, _ := ioutil.ReadDir(fmt.Sprint(log_path, "/", namespace, "/", subsystem))
	sub.SubSystem = subsystem
	sub.Pods = GetPods(filenames)
	return sub
}

func GetPods(f []os.FileInfo) (s []string) {
	for _, file := range f {
		s = append(s, getString(file))
	}
	return s
}

// os.fileInfo 转换成string
func getString(f os.FileInfo) string {
	return f.Name()

}

//  判断目录是否存在

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
