# k8s-node-manager

k8s node 节点管理

1. 凌晨4点进行容器日志归档
2. 遍历/opt/logs 目录， 截取 namespace  subsystem  podName
3. 通过k8sapi 检查pods是否存在，不存在进行podName 的归档，归档结构同日志目录

# 归档目录
完整目录: /opt/logs/yzt-env02/yzt-auth/yzt-auth-5d6d8b5f74-tjjjc/*.log
1. node节点日志目录  /opt/logs/
2. 命名空间    yzt-env02
3. 服务名     yzt-auth
4. 容器名     yzt-auth-5d6d8b5f74-tjjjc
5. 最终日志   *.log

# 归档规则
1. 定时任务凌晨3点到4点业务低峰期进行处理
2. 3点 分钟和秒 为随机时间段，防止同一时间处理造成拥堵
3. 每次检索日志目录，通过检查pod 是否存在进行日志处理，
4. 处理规则挂载备份服务器的存储到容器的pvc，根据config.yaml里面配置


# 日志归档目录
```
      - name: backup-center-logs
        persistentVolumeClaim:
          claimName: backup-center-logs

```
