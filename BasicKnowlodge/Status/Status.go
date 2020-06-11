//package status定义了后面常用的一些量
package status

import "errors"

//类型别名，大写变量名表示可以导出
type Status = error

var (
	Error Status = errors.New("失败")
	OK    Status = errors.New("成功")
)
