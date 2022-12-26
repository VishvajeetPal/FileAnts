package utils

import (
	"fmt"
	"os"
)

func SetupFileStructure() error {

	print("Creting DIR\n")
	print("Creting DIR\n")

	_, err := os.Stat("./recordsTemp")
	if err == nil {
		print(" DIR Err\n")
		//		print(err.Error())
		return nil
	}
	if os.IsNotExist(err) {
		err = os.Mkdir("./recordsTemp", 0777)
		if err != nil {
			print(" DIR Err\n")
			print(err.Error())
			return err
		}

		err = os.Mkdir("./recordsTemp/download", 0777)
		if err != nil {
			print(" DIR Err")
			print(err.Error())
			return err
		}

		err = os.Mkdir("./recordsTemp/upload", 0777)
		if err != nil {
			print(" DIR Err")
			print(err.Error())
			return err
		}
	}
	print(" DIR Created\n")
	fi, err := os.Stat("./recordsTemp")
	fmt.Printf("%v", fi)
	return nil

}
