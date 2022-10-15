package config

import (
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Database struct {
		URI string
	}

	Redis struct {
		URI     string
		Expired int
	}

	Server struct {
		Port string
		Env  string
	}
}

var (
	Cfg    Config
	Logger *zap.SugaredLogger
)

func InitLog() {

	file, _ := os.OpenFile("./src/logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	writer := zapcore.AddSync(file)

	zapInit := zap.NewProductionEncoderConfig()
	zapInit.EncodeTime = zapcore.ISO8601TimeEncoder

	encoder := zapcore.NewJSONEncoder(zapInit)
	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)

	l := zap.New(core)
	Logger = l.Sugar()

}

func ReadConfig() {

	InitLog()

	Config := &Cfg

	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		Logger.Fatal(err)
	}

	Config.Server.Port = viper.GetString("PORT")
	Config.Server.Env = viper.GetString("ENV")
	Config.Database.URI = viper.GetString("POSTGRES_URI")
	Config.Redis.URI = viper.GetString("REDIS_URI")
	Config.Redis.Expired = viper.GetInt("CACHE_TIME")

	Logger.Info("env ready")
}
