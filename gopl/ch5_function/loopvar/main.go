package main

import "os"

func main() {
	var rmDirs []func()
	tempDirs := []string {"tmp1", "tmp2", "tmp3"}
	for _, dir := range tempDirs{
		//dir := dir
		os.MkdirAll(dir, 0755)
		rmDirs = append(rmDirs, func(){
			os.RemoveAll(dir)
		})
	}

	for _, rmdir := range rmDirs {
		rmdir()
	}
}

