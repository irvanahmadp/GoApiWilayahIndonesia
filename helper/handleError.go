package helper

import (
	"log"
	"net/http"
	"os"
	"time"
)

func HandleError(w http.ResponseWriter, errText error) {
	//create log error
	now := time.Now()
	date := now.Format("2006-01-02")
	filePath := "log/error/"
	fileName := date + ".log"
	fullFilePath := filePath + fileName

	file, err := os.OpenFile(fullFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("error opening file: ", err.Error())

		//membuat folder klau blm ada
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			err = os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				log.Println("error creating directory: ", err.Error())
			} else{
				log.Println("directory created")
				file, _ = os.OpenFile(fullFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
				defer file.Close()
			}
		}
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println(errText.Error())

	//show response
	var data interface{}
	HandleResponse(w, 500, data)
}