package duplicate

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"sync"
)

func GetDuplicateFile(pathDir string) ([]string, error) {
	files, err := getAllFiles(pathDir)
	listDuplicate := []string{}
	listOrigin := []fileInfo{}

	if err != nil {
		return listDuplicate, err
	}

	for _, file := range files.list {
		exist := false
		for _, val := range listOrigin {
			if val.fileName == file.fileName && val.hash_md5 == file.hash_md5 && val.hash_sha256 == file.hash_sha256 && !exist {
				listDuplicate = append(listDuplicate, file.path)
				exist = true
			}
		}

		if !exist {
			item := fileInfo{fileName: file.fileName, hash_md5: file.hash_md5, hash_sha256: file.hash_sha256}
			listOrigin = append(listOrigin, item)
		}
	}

	return listDuplicate, nil
}

func RemoveDuplicate(duplicate []string) error {
	for _, item := range duplicate {
		err := os.Remove(item)

		if err != nil {
			return err
		}
	}
	return nil
}

func getAllFiles(path string) (FilesInfo, error) {
	list := FilesInfo{}

	files, err := ioutil.ReadDir(path)

	if err != nil {
		return list, fmt.Errorf("directory %s  does not open. error: %v", path, err)
	}

	var mutex sync.Mutex
	wg := sync.WaitGroup{}

	readDirectory(&wg, &mutex, path, files, &list)

	return list, nil
}

func readDirectory(wg *sync.WaitGroup, mutex *sync.Mutex, path string, files []fs.FileInfo, list *FilesInfo) error {
	for _, file := range files {
		newPath := path + "\\" + file.Name()
		if !file.IsDir() {
			fileByte, err := ioutil.ReadFile(newPath)
			if err != nil {
				return err
			}
			wg.Add(1)
			go addFileInfo(wg, mutex, fileByte, newPath, file.Name(), list)
		} else {

			dir, err := ioutil.ReadDir(newPath)

			if err != nil {
				return fmt.Errorf("directory %s  does not open. err: %v", newPath, err)
			}

			readDirectory(wg, mutex, newPath, dir, list)
		}

	}

	wg.Wait()
	return nil
}

func addFileInfo(wg *sync.WaitGroup, mutex *sync.Mutex, file []byte, path string, fileName string, list *FilesInfo) {
	defer wg.Done()
	h1 := md5.New()
	h1.Write(file)
	hash_md5 := hex.EncodeToString(h1.Sum(nil))

	h2 := sha256.New()
	h2.Write(file)
	hash_sha256 := hex.EncodeToString(h2.Sum(nil))

	mutex.Lock()
	list.AddItem(fileName, path, hash_md5, hash_sha256)
	mutex.Unlock()
}
