
## 简介

定时检查公网IP, 发生变化后更新(腾讯云/阿里云)域名配置.

## 配置

根据需要申请相应API密钥, 然后修改`main.go`, 配置API访问密钥以及需要绑定的域名:

```golang
const (
    // 阿里云密钥/域名配置
	aliAccessKeyId      = ""
	aliAccessKeySecrete = ""
	aliDomain           = "" // example.com
	aliSubDomain        = "" // pi
    // 腾讯云密钥/域名配置
	tencentAccessKeyId      = ""
	tencentAccessKeySecrete = ""
	tencentDomain           = "" // example.com
	tencentSubDomain        = "" // pi
)
```

- [阿里云API密钥管理](https://usercenter.console.aliyun.com/#/manage/ak)
- [腾讯云API密钥管理](https://console.cloud.tencent.com/cam/capi)

## 编译

```bash
go build .
```
会在当前目录下生成可执行文件`myddns`.

如果需要运行在树莓派上,可以使用以下交叉编译命令:

```bash
env GOOS=linux GOARCH=arm GOARM=7 go build .
```

## 运行

使用任何方式将`myddns`运行起来即可, 比如`supervisord`:

1. 安装`supervisor`
2. 将下面文件放入`/etc/supervisor/conf.d`

    > myddns.ini
    ```ini
    [program:myddns]
    command=/path/to/myddns
    ;startsecs=1      ; # of secs prog must stay up to be running (def. 1)
    ;startretries=3   ; max # of serial start failures when starting (default 3)
    redirect_stderr=true
    stdout_logfile=/var/log/myddns.log
    ```

3. `supervisorctl reload`
