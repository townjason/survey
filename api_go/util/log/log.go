package log

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"runtime"
	"time"
)

func Verbose(text string){
	_, fn, line, _ := runtime.Caller(1)
	fmt.Printf("[Verbose] %s ,line %d: %s", fn, line, text)

	//if gin.Mode() == gin.ReleaseMode {
	//	var builder strings.Builder
	//	f, _ := os.OpenFile(config.PathConfig.VerboseLogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//	defer closeFile(f)
	//	log.SetOutput(f)
	//	_, _ = fmt.Fprintf(&builder, "[Verbose] %s ,line %d: %s \n", fn, line, text)
	//	log.Println(builder.String())
	//}
}

func Error(err error){
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		fmt.Printf("[ERROR] %s ,line %d: %s \n", fn, line, err.Error())

		//if gin.Mode() == gin.ReleaseMode {
		//	var builder strings.Builder
		//	f, _ := os.OpenFile(config.PathConfig.ErrorLogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		//	defer closeFile(f)
		//	log.SetOutput(f)
		//	_,_ = fmt.Fprintf(&builder, "[ERROR] %s ,line %d: %s \n", fn, line, err.Error())
		//	log.Println(builder.String())
		//}
	}
}

//func closeFile(f *os.File){
//	if f != nil {
//		_ = f.Close()
//	}
//}

func Sql(text interface{}, values []interface{}){
	_, fn, line, _ := runtime.Caller(1)
	fmt.Printf("%s [SQL] %s ,line %d: %v \n %v \n", time.Now().Format("2006-01-02 15:04:05"), fn, line, text, values)
}

func Clean() {
	if files, err := ioutil.ReadDir("log"); err != nil {
		Error(err)
	} else {
		count := len(files) - 30
		for index, file := range files {
			if index < count {
				Error(exec.Command("rm", "-f", "log/" + file.Name()) .Run())
			}
		}
	}
}