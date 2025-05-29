package utils

import "caisse-app/app/logger"

func Errnotnil(err error) {
	if err != nil {
		logger.Error(err.Error())
	}
}
