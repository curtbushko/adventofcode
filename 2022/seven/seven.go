package seven

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/afero"
)

const (
	minDirsize     = 100000
	totalHDSpace   = 70000000
	AppNeededSpace = 30000000
)

var (
	Fs       = afero.NewMemMapFs()
	Afs      = &afero.Afero{Fs: Fs}
	RootPath = "/root"
)

func NewFs(rootPath string) {
	_ = Fs.MkdirAll(rootPath, 0755)
}

func DeleteFs(rootPath string) {
	_ = Fs.RemoveAll(RootPath)
}

func CreateFs(input string) {
	currentDir := RootPath
	NewFs(RootPath)
	for _, line := range strings.Split(input, "\n") {
		f := strings.Fields(line)
		switch f[0] {
		case "$":
			currentDir = Command(f, currentDir)
		case "dir":
			Dir(f[1], currentDir)
		default:
			File(f, currentDir)
		}
	}
}

func Command(cmd []string, currentDir string) string {
	switch cmd[1] {
	case "cd":
		switch cmd[2] {
		case "/":
			currentDir = RootPath
		case "..":
			currentDir = path.Dir(currentDir) // this will trim the last element of a path
		default:
			currentDir = path.Join(currentDir, cmd[2])
		}
	}
	return currentDir
}

func Dir(dir string, currentDir string) {
	currentDir = path.Join(currentDir, dir)
	_ = Fs.Mkdir(currentDir, 0755)
	fmt.Println("Created directory: ", currentDir)
}

func File(f []string, currentDir string) {
	size, _ := strconv.Atoi(f[0])
	name := path.Join(currentDir, f[1])
	contents := make([]byte, size)

	_ = afero.WriteFile(Fs, name, contents, 0644)
	stats, _ := Fs.Stat(name)
	fmt.Println("Created file: ", name, stats.Size())
}

func CalculateTotals(dir string) (total, smallest int) {
	dirs := make(map[string]int)

	err := Afs.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		isDir, _ := Afs.IsDir(path)
		if isDir {
			return nil
		}

		fullPath := filepath.Dir(path) // get the base path
		dirPaths := strings.Split(fullPath, "/")
		stats, _ := Fs.Stat(path)
		size := int(stats.Size())

		// Add the filesize to all the parent directories
		// Create a unique key that is make up of the path
		var key string
		for _, v := range dirPaths {
			key = key + v
			dirTotal := dirs[key]
			dirs[key] = dirTotal + size
		}

		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	// Figure out the totals
	for k, v := range dirs {
		if v < minDirsize {
			fmt.Println("Directory:", k, "Total:", v)
			total = total + v
		}
	}
	fmt.Println("Total:", total)

	// Figure out the smallest directory to remove to free up enough space
	// for the app to run
	sizes := make([]int, 0, len(dirs))
	for _, size := range dirs {
		sizes = append(sizes, size)
	}
	sort.Ints(sizes)
	availableSpace := totalHDSpace - dirs["root"]
	for _, s := range sizes {
		newSpace := availableSpace + s
		if newSpace > AppNeededSpace {
			smallest = s
			break
		}
	}
	fmt.Println("Smallest:", smallest)
	return total, smallest
}
