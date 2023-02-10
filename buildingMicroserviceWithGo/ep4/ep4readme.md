go mod init github.com/ocean5tech/repository4go/buildingMicroserviceWithGo/ep4

d, err := json.Marshal(lp) 用来把对象转换成JSON

curl localhost:9090 | jq 可以格式化表示json，但是在windows需要另外设定

vscode的扩展插件 postcode，有选项可以把测试request转成各种语言的代码


tags： 指定json各种item的行为

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

encode和Marshal，功能差不多，性能快很多
