package main

import "fmt"

var preRequests = map[string][]string{
	"算法":   {"数据结构", "微积分"},
	"数据结构": {"离散数学"},
	"微积分":  {"线性代数"},
	"编译原理": {
		"数据结构",
		"形式语言",
		"计算机组成",
	},
	"数据库":   {"数据结构"},
	"离散数学":  {"程序导论"},
	"形式语言":  {"离散数学"},
	"计算机网络": {"数据机构", "计算机组成"},
}

func topoSort(graph *map[string][]string) []string {

	seen := make(map[string]bool)
	order := []string{}
	var dfsAll func(nodes []string)
	dfsAll = func(nodes []string) {
		for _, node := range nodes {
			if !seen[node] {
				seen[node] = true
				dfsAll(preRequests[node])
				order = append(order, node)
			}
		}
	}
	keys := []string{}
	for key := range *graph {
		keys = append(keys, key)
	}
	dfsAll(keys)
	return order
}

func main() {
	for i, name := range topoSort(&preRequests) {
		fmt.Printf("%d\t%v\n", i, name)
	}
}
