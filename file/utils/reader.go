
package utils

import "os"

func ReadFile(fileName string) (string, error){
	content, err := os.ReadFile(fileName)
	if err != nil {
	//there is a problem
	return "", err
	}else{
	//Read operation
	return string(content),nil
	}

}
