package file_store

// New 创建用户配置目录下的 TOML 存储文件。
func New(softName string, fileName string) (*Store, error) {
	return newStore(softName, fileName)
}
