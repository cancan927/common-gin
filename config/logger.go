package config

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

//InitLogger 初始化zap日志
func InitLogger() *zap.SugaredLogger {
	logMode := zapcore.DebugLevel
	if !viper.GetBool("mode.dev") {
		logMode = zapcore.InfoLevel
	}

	//让日志同时输出到文件和控制台
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(
		getWriterSyncer(),
		zapcore.AddSync(os.Stdout)),
		logMode)
	return zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format("2006-01-02 15:04:05"))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSyncer() zapcore.WriteSyncer {
	stSeparater := string(filepath.Separator)
	stRootDir, _ := os.Getwd()
	stLogFilePath := stRootDir + stSeparater + "log" + stSeparater + time.Now().Format("2006-01-02")
	fmt.Println(stLogFilePath)

	luberjackSyncer := &lumberjack.Logger{
		Filename:   stLogFilePath,
		MaxSize:    viper.GetInt("log.MaxSize"),
		MaxBackups: viper.GetInt("log.MaxBackups"),
		MaxAge:     viper.GetInt("log.MaxAge"),
		Compress:   false,
	}
	return zapcore.AddSync(luberjackSyncer)
}
