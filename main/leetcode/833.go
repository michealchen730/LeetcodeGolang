package main

import (
	"sort"
	"strings"
)

type Sort833 struct {
	indexes []int
	sources []string
	targets []string
}

func (s Sort833) Len() int           { return len(s.indexes) }
func (s Sort833) Less(i, j int) bool { return s.indexes[i] < s.indexes[j] }
func (s Sort833) Swap(i, j int) {
	s.indexes[i], s.indexes[j] = s.indexes[j], s.indexes[i]
	s.sources[i], s.sources[j] = s.sources[j], s.sources[i]
	s.targets[i], s.targets[j] = s.targets[j], s.targets[i]
}

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	var res strings.Builder
	s833 := Sort833{
		indexes: indexes,
		sources: sources,
		targets: targets,
	}
	sort.Sort(s833)

	start := 0
	for k, v := range s833.indexes {
		length := len(s833.sources[k])
		if S[v:v+length] == s833.sources[k] {
			res.WriteString(S[start:v])
			res.WriteString(s833.targets[k])
			start = v + length
		}
	}
	res.WriteString(S[start:])

	return res.String()

}
