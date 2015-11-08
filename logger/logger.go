package logger

import (
	"EnCloud/osext"
	"os"
	"time"
)

func Log(message string) {

	wd, _ := osext.ExecutableFolder()

	f, err := os.OpenFile(wd+"/LogFile.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(time.Now().String() + " : " + message + "\n"); err != nil {
		panic(err)
	}
}
