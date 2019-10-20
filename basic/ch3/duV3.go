package main


import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// walkDirV3 递归地遍历以dir为根目录的整个文件树
// 并在fileSizes上发送每个已找到的文件的大小
func walkDirV3(dir string, n *sync.WaitGroup, fileSizes chan<- int64)  {
	defer n.Done()
	for _, entry := range direntsV3(dir) {
		if entry.IsDir() {
			n.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go walkDirV3(subDir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sem 是一个用于限制并发数的计数信号量
var sem = make(chan struct{}, 20)

// direntsV3 返回dir目录中的条目
func direntsV3(dir string) []os.FileInfo {
	sem <- struct{}{}    // 获取令牌
	defer func() { <-sem }()   // 释放令牌
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprint(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

// print 打印结果
func printV3(nfiles, nbytes int64)  {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

var verboseV3 = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDirV3(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// 定期输出结果
	var tick <-chan time.Time
	if *verboseV3 {
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
	printV3(nfiles, nbytes)
}

