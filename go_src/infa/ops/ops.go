package ops

import (
	"log"
	"os"
	"testing"

	"github.com/cockroachdb/errors"
)

func init() {
	log.SetOutput(os.Stdout) // 转回普通stdout
	log.SetFlags(log.Ltime) // 时间格式 13:27:44
}

// 初始化场景 快速panic 也可以把三行代码压成一行
func MustNoErr(err error) {
	if err == nil {
		return
	}
	panic(errors.WithStack(err))
}

// 测试辅助：断言 err 必须为空。
func MustNoErrInTest(t *testing.T, err error) {
	if err == nil {
		return
	}
	t.Helper()
	t.Fatalf("%+v", err)
}

// 泛型的三元表达式函数
func IfElse[T any](cond bool, trueValue, falseValue T) T {
	if cond {
		return trueValue
	}
	return falseValue
}

