/*
		Handpay ServiceMesh

           创建时间: 2020年11月25日15:55:24

	       少侠好武功,一起Giao起来
	  	 我说一Giao,你说Giao
		   一 Giao ？？？？

*/

package archive_log

import (
	"os"
	"os/exec"
	"strings"

	"gitlab.handpay.com.cn/it/k8s-node-manager/global"
	"gitlab.handpay.com.cn/it/k8s-node-manager/pkg/archive-log/getpods"
	"gitlab.handpay.com.cn/it/k8s-node-manager/pkg/k8s-api"
	"gitlab.handpay.com.cn/it/k8s-node-manager/utils"
)

// 归档
func BackupK8sLogs(logsPath, archivePath string) {
	defer utils.Tracker()()
	podlog := getpods.GetNsPods(logsPath)
	for _, ns := range podlog {
		global.Log.Info("开始处理日志路径的日志：", ns.Namespace)
		if !k8s_api.IsExistNamespace(ns.Namespace) {
			global.Log.Info("命名空间不存在，开始处理日志", ns.Namespace)
			BackupFiles(logsPath+"/"+ns.Namespace, archivePath, ns.Namespace)
			continue
		}
		for _, subsystems := range ns.SubSystem {
			global.Log.Info("开始处理日志", subsystems.SubSystem)
			if !k8s_api.IsExistSubSystem(ns.Namespace, subsystems.SubSystem) {
				global.Log.Info("子系统不存在，开始处理日志", subsystems.SubSystem)
				BackupFiles(logsPath+"/"+ns.Namespace+"/"+subsystems.SubSystem, archivePath+"/"+ns.Namespace, subsystems.SubSystem)
				continue
			}
			for _, pod := range subsystems.Pods {
				global.Log.Info("开始处理日志pod日志", ns.Namespace, "  ", pod)
				if !k8s_api.IsExistPod(ns.Namespace, pod) {
					global.Log.Info("pod不存在开始处理pod日志", ns.Namespace, "  ", pod)
					BackupFiles(logsPath+"/"+ns.Namespace+"/"+subsystems.SubSystem+"/"+pod,
						archivePath+"/"+ns.Namespace+"/"+subsystems.SubSystem, pod)

				}
			}
		}
	}

}

// 开始备份文件

func BackupFiles(src, dest, enddir string) {
	if !PathExists(dest) {
		err := os.Mkdir(dest, os.FileMode(os.ModeDir))
		if err != nil {
			global.Log.Info("目录创建失败")
			global.Log.Error(err)
		}
		//os.Mkdir(windowsReplacePath(dest), os.FileMode(fs.ModeDir))
	}
	// rename 步伐夸分区执行
	//err := os.Rename(src, dest+"/"+enddir)
	cmd := exec.Command("mv", src, dest+"/"+enddir)
	_, err := cmd.Output()
	//err := os.Rename(windowsReplacePath(src), windowsReplacePath(dest+"\\"+enddir))
	global.Log.Info("源目录:", src, " -- 目标目录", dest)
	if err != nil {
		global.Log.Warn("移动文件失败")
		global.Log.Error(err)
	}
}

// 判断所给路径文件/文件夹是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false
	}
	return false
}

// windows机器下面使用
func windowsReplacePath(path string) string {
	path = strings.Replace(path, "/", "\\", -1)
	return path
}
