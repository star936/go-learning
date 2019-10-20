package main


import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// walkDir 递归地遍历以dir为根目录的整个文件树
// 并在fileSizes上发送每个已找到的文件的大小
func walkDirV2(dir string, fileSizes chan<- int64)  {
	for _, entry := range direntsV2(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			walkDirV2(subDir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents 返回dir目录中的条目
func direntsV2(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprint(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

// print 打印结果
func printV2(nfiles, nbytes int64)  {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	// 遍历文件树
	go func() {
		for _, root := range roots {
			walkDirV2(root, fileSizes)
		}
		close(fileSizes)
	}()

	// 定期输出结果
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Microsecond)
	}
	var nfiles, nbytes int64
	loop:
		for {
			select {
			case size, ok := <-fileSizes:
				if !ok {
					break loop
				}
				nfiles++
				nbytes += size
			case <-tick:
				printV2(nfiles, nbytes)
			}
		}
		printV2(nfiles, nbytes)
}

