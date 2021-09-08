/* init执行顺序
1. 同一个文件按从上到下执行
2. 同一个包的init按文件字符串从小到大执行
3. 按引入顺序执行
4. 如果包有依赖其他包，先执行依赖包的init
*/
package main

import (
	"init/test2"

	"init/test1"
)

func main() {
	test1.Test()
	test2.Test()
}
