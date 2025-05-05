package logs

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"go.uber.org/zap"
)

var (
	logOnce        sync.Once
	logFile        *os.File
	logRequestOnce sync.Once
	logRequestFile *os.File
	Log            *zap.Logger
	LogRequest     *zap.Logger
)

func InitLogger() {
	logOnce.Do(func() {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.StacktraceKey = ""

		date := time.Now().Format("2006-01-02")

		storagesFolder := os.Getenv("STORAGES_FOLDER_ERROR_PATH")
		if _, err := os.Stat(storagesFolder); os.IsNotExist(err) {
			err := os.MkdirAll(storagesFolder, 0755)
			if err != nil {
				panic(err)
			}
		}

		filename := fmt.Sprintf("%s/%s.log", storagesFolder, date)
		var err error
		logFile, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0775)
		if err != nil {
			panic(err)
		}

		config.OutputPaths = []string{logFile.Name()}
		config.ErrorOutputPaths = []string{logFile.Name()}

		Log, err = config.Build(zap.AddCallerSkip(1))
		if err != nil {
			panic(err)
		}
	})
	// Delete log files older than LOG_ERROR_EXP
	storagesFolder := os.Getenv("STORAGES_FOLDER_ERROR_PATH")
	logErrExp, _ := strconv.Atoi(os.Getenv("LOG_ERROR_EXP"))
	_ = filepath.Walk(storagesFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		// Check if the path is a file and not the storagesFolder itself
		if info.Mode().IsRegular() && path != storagesFolder {
			fmt.Println("Checking file:", path)
			fmt.Println("File age (minutes):", time.Since(info.ModTime()).Minutes())

			// Check if the file is older than LOG_ERROR_EXP
			if time.Since(info.ModTime()).Hours() > (float64(logErrExp) * 24) {
				fmt.Println("Removing file:", path)
				err := os.Remove(path)
				if err != nil {
					fmt.Println("Error removing file:", err)
					return err
				}
			}
		}
		return nil
	})
}

func InitLoggerRequest() {
	logRequestOnce.Do(func() {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.CallerKey = ""
		config.EncoderConfig.StacktraceKey = ""

		date := time.Now().Format("2006-01-02")

		storagesFolder := os.Getenv("STORAGES_FOLDER_INFO_PATH")
		if _, err := os.Stat(storagesFolder); os.IsNotExist(err) {
			err := os.MkdirAll(storagesFolder, 0755)
			if err != nil {
				panic(err)
			}
		}

		filename := fmt.Sprintf("%s/%s.log", storagesFolder, date)
		var err error
		logRequestFile, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0775)
		if err != nil {
			panic(err)
		}

		config.OutputPaths = []string{logRequestFile.Name()}
		config.ErrorOutputPaths = []string{logRequestFile.Name()}

		LogRequest, err = config.Build()
		if err != nil {
			panic(err)
		}
	})
	// Delete log files older than LOG_INFO_EXP
	storagesFolder := os.Getenv("STORAGES_FOLDER_INFO_PATH")
	logInfoExp, _ := strconv.Atoi(os.Getenv("LOG_INFO_EXP"))
	_ = filepath.Walk(storagesFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		// Check if the path is a file and not the storagesFolder itself
		if info.Mode().IsRegular() && path != storagesFolder {
			fmt.Println("Checking file:", path)
			fmt.Println("File age (minutes):", time.Since(info.ModTime()).Minutes())

			// Check if the file is older than LOG_ERROR_EXP
			if time.Since(info.ModTime()).Hours() > (float64(logInfoExp) * 24) {
				fmt.Println("Removing file:", path)
				err := os.Remove(path)
				if err != nil {
					fmt.Println("Error removing file:", err)
					return err
				}
			}
		}
		return nil
	})
}

func CloseLogReq() {
	if logRequestFile != nil {
		logRequestFile.Close()
	}
	if LogRequest != nil {
		LogRequest.Sync()
	}
}

func Close() {
	if logFile != nil {
		logFile.Close()
	}
	if Log != nil {
		Log.Sync()
	}
}

func Info(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		Log.Info(v.Error(), fields...)
	case string:
		Log.Info(v, fields...)
	}
}

func Debug(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		Log.Debug(v.Error(), fields...)
	case string:
		Log.Debug(v, fields...)
	}
}

func Error(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		Log.Error(v.Error(), fields...)
	case string:
		Log.Error(v, fields...)
	}
}
