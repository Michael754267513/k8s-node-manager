/*
		Handpay ServiceMesh

           创建时间: 2020年11月25日15:55:24

	       少侠好武功,一起Giao起来
	  	 我说一Giao,你说Giao
		   一 Giao ？？？？

*/

package k8s_api

import (
	"context"
	"io/ioutil"

	"gitlab.handpay.com.cn/it/k8s-node-manager/global"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func IsExistNamespace(namespace string) bool {
	k8s_client, err := InitClient()
	global.Log.Info("配置k8s clientset")
	if err != nil {
		global.Log.Errorf("配置k8获取异常")
		global.Log.Error(err)
		return false
	}
	global.Log.Info("获取namespace")
	if _, err := k8s_client.CoreV1().Namespaces().Get(context.Background(), namespace, v1.GetOptions{}); err != nil {
		global.Log.Info("记录错误信息")
		global.Log.Info("错误信息:", err)
		if errors.IsNotFound(err) {
			global.Log.Info(namespace, "不存在namespace", namespace)
			return false
		}
	}
	return true
}
func IsExistPod(namespace, podName string) bool {
	k8s_client, err := InitClient()
	global.Log.Info("配置k8s clientset")
	if err != nil {
		global.Log.Error(err)
		return false
	}
	global.Log.Info("获取pod")
	if _, err := k8s_client.CoreV1().Pods(namespace).Get(context.Background(), podName, v1.GetOptions{}); err != nil {
		global.Log.Info("记录错误信息")
		global.Log.Error(err)
		global.Log.Info("错误信息:", err)
		if errors.IsNotFound(err) {
			global.Log.Info(namespace, "不存在容器", podName)
			return false
		}
	}
	global.Log.Info(namespace, "存在容器", podName)
	return true

}
func IsExistSubSystem(namespace, subsystem string) bool {
	k8s_client, err := InitClient()
	global.Log.Info("配置k8s clientset")
	if err != nil {
		global.Log.Error(err)
		return false
	}
	global.Log.Info("获取子系统：", subsystem)
	if _, err := k8s_client.AppsV1().Deployments(namespace).Get(context.Background(), subsystem, v1.GetOptions{}); err != nil {
		global.Log.Info("记录错误信息")
		global.Log.Info("错误信息:", err)

		if errors.IsNotFound(err) {
			global.Log.Info(namespace, "不存在子系统", subsystem)
			return false
		}
	}
	global.Log.Info(namespace, " 存在子系统", subsystem)
	return true

}

//
// 初始化k8s客户端
func InitClient() (clientset *kubernetes.Clientset, err error) {
	var (
		restConf *rest.Config
	)

	if restConf, err = GetRestConf(); err != nil {
		return
	}

	// 生成clientset配置
	if clientset, err = kubernetes.NewForConfig(restConf); err != nil {
		goto END
	}
END:
	return
}

// 获取k8s restful client配置
func GetRestConf() (restConf *rest.Config, err error) {
	var (
		kubeconfig []byte
	)
	// 读kubeconfig文件
	if kubeconfig, err = ioutil.ReadFile("./admim.config"); err == nil {
		// 生成rest client配置
		if restConf, err = clientcmd.RESTConfigFromKubeConfig(kubeconfig); err != nil {
			goto END
		}
	} else {
		// 从容器SA里面获取配置文件
		restConf, err = rest.InClusterConfig()
		if err != nil {
			goto END
		}
	}

END:
	return
}

func Logger(err error) {
	global.Log.Errorf("获取k8s 配置失败")
	global.Log.Error(err)
}
