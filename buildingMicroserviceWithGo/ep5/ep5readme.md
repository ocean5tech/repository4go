go mod init github.com/ocean5tech/repository4go/buildingMicroserviceWithGo/ep5

Gorilla framework

参考
https://goproxy.io/zh/   

遇到如下错误
PS C:\drive_D\repository4go\buildingMicroserviceWithGo\ep5> go mod tidy
go: finding module for package github.com/gorilla/mux
github.com/ocean5tech/repository4go/buildingMicroserviceWithGo/ep5 imports
        github.com/gorilla/mux: module github.com/gorilla/mux: Get "https://proxy.golang.org/github.com/gorilla/mux/@v/list": dial tcp 142.251.42.241:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.

是因为proxy.golang.org被墙了，所以需要设定环境

Bash (Linux or macOS)
# 配置 GOPROXY 环境变量
export GOPROXY=https://proxy.golang.com.cn,direct
# 还可以设置不走 proxy 的私有仓库或组，多个用逗号相隔（可选）
export GOPRIVATE=git.mycompany.com,github.com/my/private

PowerShell (Windows)
# 配置 GOPROXY 环境变量
$env:GOPROXY = "https://proxy.golang.com.cn,direct"
# 还可以设置不走 proxy 的私有仓库或组，多个用逗号相隔（可选）
$env:GOPRIVATE = "git.mycompany.com,github.com/my/private"

使配置长久生效 （推荐）

Mac/Linux
# 设置你的 bash 环境变量
echo "export GOPROXY=https://proxy.golang.com.cn,direct" >> ~/.profile && source ~/.profile

# 如果你的终端是 zsh，使用以下命令
echo "export GOPROXY=https://proxy.golang.com.cn,direct" >> ~/.zshrc && source ~/.zshrc

Windows
1. 右键 我的电脑 -> 属性 -> 高级系统设置 -> 环境变量
2. 在 “[你的用户名]的用户变量” 中点击 ”新建“ 按钮
3. 在 “变量名” 输入框并新增 “GOPROXY”
4. 在对应的 “变量值” 输入框中新增 “https://proxy.golang.com.cn,direct”
5. 最后点击 “确定” 按钮保存设置