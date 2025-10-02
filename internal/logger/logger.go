package logger

import "go.uber.org/zap"

var Log *zap.Logger

func Init(prod bool) error {
	var err error
	if prod {
		Log, err = zap.NewProduction()
	} else {
		Log, err = zap.NewDevelopment()
	}
	return err
}

func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}
