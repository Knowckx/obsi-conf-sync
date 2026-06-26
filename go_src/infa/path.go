// go/inner/infa2/path.go
package infa

import (
	"errors"
	"os"
	"path/filepath"
)

// errGoModNotFound 表示从指定目录向上查找时，没有找到 go.mod。
// 这是一个可预期错误，用于区分“项目根目录不存在”和“系统调用失败”。
var errGoModNotFound = errors.New("go.mod not found")

// FindProjectRoot 查找项目根目录。
//
// 查找策略：
//  1. 优先从当前工作目录开始，向上查找 go.mod。
//     这适合开发环境、go run、IDE 运行、测试运行等场景。
//  2. 如果当前工作目录下找不到 go.mod，则从可执行文件所在目录开始继续查找。
//     这适合程序被编译后，从二进制文件所在位置运行的场景。
//  3. 如果两边都找不到 go.mod，则返回可执行文件所在目录。
//     这适合生产环境，因为生产环境通常不会携带 go.mod。
func FindProjectRoot() (string, error) {
	// 1. 优先从当前工作目录开始查找。
	//
	// 在开发环境中，当前工作目录通常就是项目目录，或者项目目录的子目录。
	// 比如：
	//   go run ./cmd/app
	//   go test ./...
	//   IDE 直接运行
	//
	// 这些情况下，从 os.Getwd() 开始查找 go.mod 通常最准确。
	if wd, err := os.Getwd(); err == nil {
		root, err := findGoModRoot(wd)
		if err == nil {
			return root, nil
		}

		// 如果不是“找不到 go.mod”，说明发生了真实错误。
		// 例如权限问题、路径不可访问等，这类错误应该直接返回。
		if !errors.Is(err, errGoModNotFound) {
			return "", err
		}
	}

	// 2. 当前工作目录没有找到 go.mod，则尝试从可执行文件所在目录查找。
	//
	// 在编译后的程序中，当前工作目录可能不是程序所在目录。
	// 例如用户从其他目录执行：
	//
	//   /opt/myapp/app
	//
	// 此时 os.Getwd() 可能是用户当前所在目录，而不是 /opt/myapp。
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	exeDir := filepath.Dir(exePath)

	// 从可执行文件所在目录继续向上查找 go.mod。
	//
	// 如果二进制文件仍然放在项目目录或项目子目录中，
	// 这里仍然可以找到项目根目录。
	if root, err := findGoModRoot(exeDir); err == nil {
		return root, nil
	} else if !errors.Is(err, errGoModNotFound) {
		return "", err
	}

	// 3. 两种方式都找不到 go.mod，则认为当前处于生产环境。
	//
	// 生产环境中通常不会部署 go.mod。
	// 此时返回可执行文件所在目录，可以作为配置文件、静态资源等文件的基准路径。
	return exeDir, nil
}

// findGoModRoot 从 start 目录开始，逐级向上查找 go.mod。
// 一旦找到 go.mod，就返回它所在的目录。
// 如果查找到文件系统根目录仍然没有找到，则返回 errGoModNotFound。
func findGoModRoot(start string) (string, error) {
	// 先转成绝对路径，避免相对路径在向上递归时产生歧义。
	dir, err := filepath.Abs(start)
	if err != nil {
		return "", err
	}

	for {
		goModPath := filepath.Join(dir, "go.mod")

		// 判断当前目录下是否存在 go.mod。
		if _, err := os.Stat(goModPath); err == nil {
			return dir, nil
		} else if !errors.Is(err, os.ErrNotExist) {
			// 如果错误不是“文件不存在”，说明发生了真实系统错误。
			// 例如权限不足、路径损坏等，需要返回给调用方。
			return "", err
		}

		// 获取父目录。
		parent := filepath.Dir(dir)

		// 如果父目录等于当前目录，说明已经到达文件系统根目录。
		//
		// Unix 示例：
		//   filepath.Dir("/") == "/"
		//
		// Windows 示例：
		//   filepath.Dir("C:\\") == "C:\\"
		if parent == dir {
			return "", errGoModNotFound
		}

		// 继续向上一层目录查找。
		dir = parent
	}
}