package utils

import "os"

func SetupFileStructure() error {

	_, err := os.Stat("./recordsTemp")
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = os.Mkdir("./recordsTemp", 0777)
		if err != nil {
			return err
		}

		err = os.Mkdir("./recordsTemp/download", 0777)
		if err != nil {
			return err
		}

		err = os.Mkdir("./recordsTemp/upload", 0777)
		if err != nil {
			return err
		}
	}

	return nil

}
