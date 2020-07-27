### 我在 GitHub 上的统计

![louis's github stats](https://github-readme-stats.vercel.app/api?username=oldthreefeng&show_icons=true&hide_border=false)

<!--events start -->

### 我的近期动态

⭐️ Star [个人主页](https://github.com/oldthreefeng/oldthreefeng) 后会自动更新，最近更新时间：`2020-07-27 15:09:21`

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
*  [kuberbetes ingress ssl证书配置](https://www.fenghong.tech/blog/kubernetes/kubernetes-ingress-ssl/)

  > [TOC]
 可以直接在 Ingress 中配置 HTTPS 证书，使得你的网站支持 HTTPS 协议。
使用openssl创建自用证书或使用acme创建免费的证书
*  [vim-go插件安装](https://www.fenghong.tech/blog/go/go-vim/)

  > [TOC]
前言  vim是linux系统下常用的代码编辑器，默认情况下不支持go的代码高亮和语法检查，不过可以通过安装vim插件来支持go的开发 Vim-go是当前使用最为广泛的用于搭建Golang开发环境的vim插件，这里我同样使用vim-go作为核心和基础进行环境搭建的。vim-go利 用开源Vim插件管理器安装，gmarik/Vundle
*  [git-rebase合并多commit记录](https://www.fenghong.tech/blog/ops/git-rebase/)

  > [toc]
新建测试仓库 为了不破坏现有的仓库，我们首先创建一个新建一个实验repo，所有操作都在该仓库下操作。创建命令如下：
$ mkdir rebase-repo $ cd rebase-repo $ git init rebase-repo $ git commit --allow-empty -m &quot;init&quot; $ git remote add origin git@code
*  [记一次curl版本yum升级至7.71.0](https://www.fenghong.tech/blog/ops/curl-data-raw/)

  > [TOC]
背景  偶然有一次机会使用到了curl 命令行中 --data-raw 选项， 但是提示是curl: option --data-raw: is unknown, 网上查询了蛮多， 也没有写什么解决方案，估摸着是版本太低的缘故。
 目前的centos 7 的yum仓库版本, 最新版应该到了 7
*  [jenkins+java+docker+kubernetes实现自动化部署](https://www.fenghong.tech/blog/ops/jenkins-java-docker-kubernetes/)

  > [TOC]
背景  基于springboot项目, 部署在kubernetes集群, 并实现自动化构建部署


<!--events end -->
