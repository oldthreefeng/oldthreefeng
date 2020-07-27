### 我在 GitHub 上的统计

![louis's github stats](https://github-readme-stats.vercel.app/api?username=oldthreefeng&show_icons=true&hide_border=false)

<!--events start -->

### 我的近期动态

⭐️ Star [个人主页](https://github.com/oldthreefeng/oldthreefeng) 后会自动更新，最近更新时间：`2020-07-27 15:13:57`

📝*  [sealos 一键安装 kubernetes1.18.0](https://www.fenghong.tech/blog/kubernetes/sealos-install/)

  > [TOC]
 一个二进制工具加一个资源包，不依赖haproxy keepalived ansible等重量级工具，一条命令就可实现kubernetes高可用集群构建，无论是单节点还是集群，单master还是多master，生产还是测试都能很好支持！简单不意味着阉割功能，照样能全量支持kubeadm所有配置。 立即获取sealos
 效果图 dashboard
grafana
环境准备 主机名 设置永久主机名称，然后重新登录:
$ hostnamectl set-hostname k8s-master # 将 master 替换为当前主机名 $ cat /etc/redhat-release CentOS Linux release 7
*  [kubernetes ingress rewrite配置](https://www.fenghong.tech/blog/kubernetes/kubernetes-ingress-rewrite/)

  > [TOC]
 可以直接在 Ingress 中配置 rewrite
*  [lvm磁盘在线扩缩容](https://www.fenghong.tech/blog/ops/lvm-reduce-extend/)

  > [TOC]
 背景: 公司安装centos7
*  [阿里云日志方案说明](https://www.fenghong.tech/blog/ops/log-for-aliyun/)

  > [TOC]
一 现状与需求 现状与问题:  日志文件分散在各个应用服务器，开发人员必须远程登录才能查看日志，不利于服务器安全管控，加大生产服务器的风险； 服务器上各项目日志配置很随意，文件分布杂乱，没有统一的规范和管理； 日志文件占用服务器大量的硬盘空间，如不及时清理会发生硬盘占满，影响系统的正常运行; 对于超过百兆的日志文件根本没法打开和关键字搜索，不利于问题的快速定位和排查; 集群和分布式的系统需要查看多个服务器的日志; 日志保存的时间不统一，不能长时间保存日志
*  [frp内网穿透http获取客户端真实ip](https://www.fenghong.tech/blog/ops/frp-get-realip/)

  > [TOC]
 背景: frp是一款内网穿透神器, 本地的http可以通过frp穿透至外网,详情可以查看官方网站Github
*  [httpProxy代理导致kubernetes出现的异常错误](https://www.fenghong.tech/blog/kubernetes/kubernetes-error-proxy/)

  > [TOC]
 在创建好kubernetes集群后， 内网访问便是稀松平常的事情。 由于使用了http_proxy代理，但是未设置no_proxy导致各种错误。为什么设置代理，因为 github头像老是不显示。
 排查思路 内网的另外一台主机(这台主机用的是provixy代理)使用curl命令进行排查，查看集群version信息，没有任何异常返回200。
$ curl -ik https://apiserver


<!--events end -->
