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
	"net/http"
	"os"
	"strings"

	"github.com/robfig/cron"
	"gitlab.handpay.com.cn/it/k8s-node-manager/global"
	archive_log "gitlab.handpay.com.cn/it/k8s-node-manager/pkg/archive-log"
	"k8s.io/apimachinery/pkg/util/rand"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health)
	http.ListenAndServe(":9999", mux)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

// 初始化配置 创建日志归档任务
func init() {

	global.K8sNodeManageConfig, _ = global.K8sNodeManageConfig.GetConfig("./config.yaml")
	if global.K8sNodeManageConfig.K8sLogs.EnableK8sLog == true {
		if global.K8sNodeManageConfig.K8sLogs.ArchiveLogTimer == "" ||
			global.K8sNodeManageConfig.K8sLogs.LogPath == "" ||
			global.K8sNodeManageConfig.K8sLogs.ArchiveLogPath == "" {
			global.Log.Error("开启容器日志归档，容器配置参数不全")
			os.Exit(999)
		}
	}
	// 开启定时任务
	c := cron.New()
	c.Start()

	if err := c.AddFunc(randomTime(global.K8sNodeManageConfig.K8sLogs.ArchiveLogTimer),
		func() {
			archive_log.BackupK8sLogs(global.K8sNodeManageConfig.K8sLogs.LogPath,
				global.K8sNodeManageConfig.K8sLogs.ArchiveLogPath)
		}); err != nil {
		global.Log.Warn("创建定时任务失败")
		global.Log.WithError(err)
		os.Exit(999)
	}

}

// 按照时间随机操作
func randomTime(spec string) (s string) {
	seed := fmt.Sprint(rand.IntnRange(0, 59))
	s = strings.Replace(spec, "0", seed, 2)
	return
}
