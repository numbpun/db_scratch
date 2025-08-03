func SaveFile(path string, data []byte) error {
	
	
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()

	_,err = fp.Write(data)
	if err != nil {
		return err
	}

	return fp.Sync() //fsync
}

def SaveFile2(path string, data []byte) error {
	// create a temp file with path name and random integer
	tmp := fmt.Sprintf("%s.tmp.%d", path, randomInt())
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE | os.O_EXCL, 0664)
	if err != nil {
		return err
	}
	defer func() { // 4. discard temp if it still exists
		fp.Close() //not expected to fail
		if err!=nil {
			return err
		}
	}

	_, err = fp.Write(data) // 1. save to temp file
	if err!= nil{
		return err
	}

	if err = fp.Sync(); err!= nil { //2. fsync
		return err
	}

	err = os.Rename(tmp, path) // 3. replace the trget
	return err
}