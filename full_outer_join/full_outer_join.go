package full_outer_join

import (
	"log"
	"os"
	"sort"
	"strings"
)

func FullOuterJoin(f1Path, f2Path, resultPath string) {
	f1, err := os.ReadFile(f1Path)
	if err != nil {
		log.Fatalf("couldn't open f1: %v", err)
	}
	f2, err := os.ReadFile(f2Path)
	if err != nil {
		log.Fatalf("couldn't open f2: %v", err)
	}
	m := map[string]bool{}
	read(m, f1, f2)
	lines := make([]string, 0, len(m))
	for str, is := range m {
		if is {
			lines = append(lines, str)
		}
	}
	sort.Strings(lines)
	result := strings.Builder{}
	for i, line := range lines {
		result.Write([]byte(line))
		if i < len(lines)-1 {
			result.WriteByte('\n')
		}
	}
	if err := os.WriteFile(resultPath, []byte(result.String()), os.ModePerm); err != nil {
		log.Fatalf("could not write result: %s", err)
	}
}

func read(m map[string]bool, files ...[]byte)  {
	for _, file := range files {
		s := strings.Split(string(file), "\n")
		for _, v := range s {
			if _, is := m[v]; is {
				m[v] = false
			} else {
				m[v] = true
			}
		}
	}
}
