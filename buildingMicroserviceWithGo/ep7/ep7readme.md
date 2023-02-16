go mod init github.com/ocean5tech/repository4go/buildingMicroserviceWithGo/ep7

swagger的安装 - 从source安装
1. goswagger.io/install.html
2. 建立目录 C:\drive_D\go-swagger
3. git clone https://github.com/go-swagger/go-swagger
4. go install 需要mod.go 文件，所以要进入下一层go-swagger文件夹
   C:\drive_D\go-swagger\go-swagger>
5. go install ./cmd/swagger
6. 安装到环境变量的GOPATH的bin下面中去了
   C:\Users\201744672\go\bin\swagger.exe


   https://goswagger.io/use/spec/meta.html


1. ！！！下面这两句之间，不能有空白行，否则生成空白swaggerymal文件！！！
        // swagger:meta
        package handlers

2. 用下面的例子来改造swagger注释

        // Package classification Petstore API.
        //
        // the purpose of this application is to provide an application
        // that is using plain go code to define an API
        //
        // This should demonstrate all the possible comment annotations
        // that are available to turn go code into a fully compliant swagger 2.0 spec
        //
        // Terms Of Service:
        //
        // there are no TOS at this moment, use at your own risk we take no responsibility
        //
        //     Schemes: http, https
        //     Host: localhost
        //     BasePath: /v2
        //     Version: 0.0.1
        //     License: MIT http://opensource.org/licenses/MIT
        //     Contact: John Doe<john.doe@example.com> http://john.doe.com
        //
        //     Consumes:
        //     - application/json
        //     - application/xml
        //
        //     Produces:
        //     - application/json
        //     - application/xml
        //
        //     Security:
        //     - api_key:
        //
        //     SecurityDefinitions:
        //     api_key:
        //          type: apiKey
        //          name: KEY
        //          in: header
        //     oauth2:
        //         type: oauth2
        //         authorizationUrl: /oauth2/auth
        //         tokenUrl: /oauth2/token
        //         in: header
        //         scopes:
        //           bar: foo
        //         flow: accessCode
        //
        //     Extensions:
        //     x-meta-value: value
        //     x-meta-array:
        //       - value1
        //       - value2
        //     x-meta-array-obj:
        //       - name: obj
        //         value: field
        //
        // swagger:meta

3. makefile 在windows下不好用，执行下面的命令
     3-1. go env -w GO111MODULE=off
     3-2. 在这个路径下，C:\drive_D\repository4go\buildingMicroserviceWithGo\ep7\handlers
            swagger generate spec -o ./swagger.yaml --scan-models


git路径
https://github.com/go-swagger/go-swagger

redoc
https://github.com/redocly/redoc

https://github.com/go-openapi/runtime/tree/master/middleware