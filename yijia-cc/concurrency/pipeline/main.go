package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

func main() {
	// Calculate the MD5 sum of all files under the specified directory,
	// then print the results sorted by path name.
	if len(os.Args) != 2 {
		fmt.Println("input directory must be specified")
	}

	hashes, err := md5AllFiles(os.Args[1])
	if err != nil {
		panic(err)
	}

	paths := make([]string, 0)
	for path := range hashes {
		paths = append(paths, path)
	}

	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%s %x\n", path, hashes[path])
	}
}

type result struct {
	path string
	hash [md5.Size]byte
	err  error
}

func md5AllFiles(dir string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

	paths, errCh := walkFiles(done, dir)

	output := make(chan result)
	wg := sync.WaitGroup{}

	const numWorkers = 15
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			readAndHash(done, paths, output)
		}()
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	hashes := make(map[string][md5.Size]byte)
	for res := range output {
		if res.err != nil {
			return nil, res.err
		}
		hashes[res.path] = res.hash
	}
	if err := <-errCh; err != nil {
		return nil, err
	}
	return hashes, nil
}

func walkFiles(done <-chan struct{}, root string) (chan string, chan error) {
	paths := make(chan string)
	errCh := make(chan error, 1)
	go func() {
		defer close(paths)
		err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				// Ignore if not file available to read
				return nil
			}
			select {
			case paths <- path:
			case <-done:
				return errors.New("walk cancelled")
			}
			return nil
		})
		errCh <- err
	}()
	return paths, errCh
}

func readAndHash(done chan struct{}, paths chan string, output chan result) {
	for path := range paths {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			output <- result{err: err}
		}

		hash := md5.Sum(data)
		select {
		case output <- result{path: path, hash: hash}:
		case <- done:
			return
		}
	}
}
