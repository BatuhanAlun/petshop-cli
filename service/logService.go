package service

import "petshop/pkg"

func AddLog(userId int, userRole, logMessage string) error {
	return pkg.AddLog(userId, userRole, logMessage)
}
