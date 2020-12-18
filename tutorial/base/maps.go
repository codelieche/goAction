/*
Map：
- 创建：make(map[string]int)
- 获取元素：m[key]
- key不存在时，获得Value类型的初始值
- 用value, ok := m[key] 来判断是否存在key，如果ok是true就表示有
- delete删除其中key
- 使用range遍历key，或者遍历key. value对
- 注意，是个hash值的，key是无序的
- 使用len可以获取元素的个数

Map的key
- map的key：go语言，map使用哈希表，必须可以比较相等
- 除了 slice, map, function的内建类型都可以作为key
- Struct类型不包含上述字段，可以作为key
 */

package main

import "fmt"

func main() {
	mapDemo01()

}

func mapDemo01() {
	m := map[string]string {
		"name": "Jim",
		"age": "18",
		"email": "admin@",
	}

	fmt.Println("m = ", m, "\nname: ", m["name"])
	// map[age:18 email:admin@ name:Jim]
	// 用for循环打印map
	fmt.Println("\n=== 遍历map ===")
	for k, v := range m {
		fmt.Println(k, "===>", v)
	}

	fmt.Println("\n===== m2 ====")
	m2 := make(map[string]int)

	m2["jim"] = 18
	m2["tom"] = 19
	for k, v := range m2 {
		fmt.Println(k, "===>", v)
	}
	fmt.Println("取个不存在的值jim2：", m2["jim2"])
	// 删除key
	delete(m2, "tom")
	fmt.Println("m2 delete tom后： m2 = ", m2)


	fmt.Println("\n===== m3 ====")

	var m3 map[string]int
	if m3 == nil {
		fmt.Println("m3 is nil")
	}
}
