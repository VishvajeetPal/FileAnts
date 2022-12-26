package utils

import (
	"fmt"
	"os"
)

func SetupFileStructure() error {

	print("Creting DIR\n")
	print("Creting DIR\n")

	_, err := os.Stat("/app/recordsTemp")
	if err == nil {
		print("DIR Was present\n")
		return nil
	}
	if os.IsNotExist(err) {
		err = os.Mkdir("/app/recordsTemp", 0777)
		if err != nil {
			print(" DIR Err\n")
			print(err.Error())
			return err
		}

		err = os.Mkdir("/app/recordsTemp/download", 0777)
		if err != nil {
			print(" DIR Err")
			print(err.Error())
			return err
		}

		err = os.Mkdir("/app/recordsTemp/upload", 0777)
		if err != nil {
			print(" DIR Err")
			print(err.Error())
			return err
		}
	}
	print(" DIR Created\n")
	fi, err := os.Stat("/app/recordsTemp")
	fmt.Printf("%v", fi)
	return nil

}
