package day12

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

type DataMap map[string][]string

func MakeMap(r io.Reader) (DataMap, error) {

	s := bufio.NewScanner(r)

	result := DataMap{}

	for s.Scan() {
		parts := strings.Split(s.Text(), "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("unexpected input format found %q", s.Text())
		}

		result[parts[0]] = append(result[parts[0]], parts[1])
		result[parts[1]] = append(result[parts[1]], parts[0])
	}

	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("failed to read all input data: %w", err)
	}

	for k, v := range result {
		sort.Strings(v)
		result[k] = v
	}

	return result, nil
}

func ReadData(fn string) (DataMap, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer f.Close()
	return MakeMap(f)
}

var (
	allCaps = regexp.MustCompile(`^[A-Z]+$`)
)

func canVisitDoor(pos string, path []string) bool {
	if allCaps.MatchString(pos) {
		return true
	}

	for _, v := range path {
		if pos == v {
			return false
		}
	}

	return true
}

func tryDoor(dm DataMap, nextPos string, currPath []string) [][]string {
	// fmt.Printf("tryDoor: nextPos: %s, currPath: %#v\n", nextPos, currPath)

	if nextPos == "end" {
		return [][]string{
			append(currPath, "end"),
		}
	}

	allPaths := [][]string{}

	for _, next := range dm[nextPos] {
		if !canVisitDoor(next, currPath) {
			continue
		}

		cp := make([]string, len(currPath)+1)
		copy(cp, currPath)
		paths := tryDoor(dm, next, append(currPath, nextPos))
		allPaths = append(allPaths, paths...)

	}

	return allPaths
}

func findPath(dm DataMap, currPaths [][]string) [][]string {

	for _, next := range dm["start"] {
		paths := tryDoor(dm, next, []string{"start"})
		currPaths = append(currPaths, paths...)
	}

	return currPaths
}

func FindPaths(dm DataMap) [][]string {
	return findPath(dm, [][]string{})
}
