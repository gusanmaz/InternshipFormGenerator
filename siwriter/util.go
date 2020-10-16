package siwriter

import "os"

func DirExists(dirPath string, createDir bool)bool{
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err){
		if createDir{
			os.Mkdir(dirPath, 0755)
			return true
		}
		return false
	}
	return true
}
