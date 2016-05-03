package milog 

// package main 

// import (
// 	"fmt"
// 	"os"
// 	"io"

// 	"github.com/nguyendangminh/milog"
// )

// func main() {
// 	logger := milog.New()
// 	logger.Info("Info message")
// 	logger.Error("Error message")

// 	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err != nil {
// 	    fmt.Println("Could not write to file")
// 	}

// 	logger = milog.NewWithWriters(os.Stdout, os.Stdout, file)
// 	logger.Info("Info message")
// 	logger.Error("Error message")

// 	logger = milog.NewWithWriters(os.Stdout, os.Stdout, io.MultiWriter(file, os.Stdout))
// 	logger.Debug("Debug message")
// 	logger.Error("Error message")

// 	logger, err = milog.NewWithFile("log.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		logger.Debug("Debug message")	
// 	}
// }