go mod init github.com/ocean5tech/repository4go/buildingMicroserviceWithGo/ep4

git pull

git checkout -b  episode_4

curl localhost:9090 -d '{}'  可以send一个json

正则表达式
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)


定制error
    var ErrProductNotFound = fmt.Errorf("Product not found")

查询数组
func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}