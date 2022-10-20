package helper

import (
	"context"
	"os"
)

func GetRequestId(c context.Context) string {
	request_id, _ := c.Value("request_id").(string)
	return request_id
}

func IsProductionEnv() bool {
	return os.Getenv("ENV") == "PROD"
}
