package util

import "m4-im/pkg/setting"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
	Salt1 = setting.AppSetting.Salt1
	Salt2 = setting.AppSetting.Salt2
}
