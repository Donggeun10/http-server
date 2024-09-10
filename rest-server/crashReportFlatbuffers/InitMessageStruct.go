package CrashReportFlatbuffers

import "time"

type AppLoad struct {
	GameCode        string
	Os              string
	OsVersion       string
	Model           string
	CrashSDKVersion string
	AppVersion      string
	PackageName     string
	UserKey         string
	DeviceId        string
	Geo             string
	City            string
	ReportDatetime  int64
	MemoryWarning   string
	Carrier         string
	SessionKey      string
	Emulator        bool
	NetworkKind     string
	VendorUserKey   string
	VendorDeviceId  string
	ReceiveDatetime string
}

type AppLoadV3 struct {
	AppLoad
}

func AppLoadV3FromFB(AppLoadLogV3 *AppLoadLogV3) *AppLoadV3 {

	apploadV3 := &AppLoadV3{}
	apploadV3.GameCode = string(AppLoadLogV3.GameCode())
	apploadV3.Os = string(AppLoadLogV3.Os())
	apploadV3.OsVersion = string(AppLoadLogV3.OsVersion())
	apploadV3.Model = string(AppLoadLogV3.Model())
	apploadV3.CrashSDKVersion = string(AppLoadLogV3.CrashSDKVersion())
	apploadV3.AppVersion = string(AppLoadLogV3.AppVersion())
	apploadV3.PackageName = string(AppLoadLogV3.PackageName())
	apploadV3.UserKey = string(AppLoadLogV3.UserKey())
	apploadV3.DeviceId = string(AppLoadLogV3.DeviceId())
	apploadV3.Geo = string(AppLoadLogV3.Geo())
	apploadV3.City = string(AppLoadLogV3.City())
	apploadV3.ReportDatetime = AppLoadLogV3.ReportDatetime()
	apploadV3.MemoryWarning = string(AppLoadLogV3.MemoryWarning())
	apploadV3.Carrier = string(AppLoadLogV3.Carrier())
	apploadV3.SessionKey = string(AppLoadLogV3.SessionKey())
	apploadV3.Emulator = AppLoadLogV3.Emulator()
	apploadV3.NetworkKind = string(AppLoadLogV3.NetworkKind())
	apploadV3.VendorUserKey = string(AppLoadLogV3.VendorUserKey())
	apploadV3.VendorDeviceId = string(AppLoadLogV3.VendorDeviceId())
	apploadV3.ReceiveDatetime = string(time.Now().Format(time.RFC3339))

	return apploadV3

}

type AppLoadV4 struct {
	AppLoad
	buildCode string
}

func AppLoadV4FromFB(AppLoadLogV4 *AppLoadLogV4) *AppLoadV4 {

	apploadV4 := &AppLoadV4{}
	apploadV4.GameCode = string(AppLoadLogV4.GameCode())
	apploadV4.Os = string(AppLoadLogV4.Os())
	apploadV4.OsVersion = string(AppLoadLogV4.OsVersion())
	apploadV4.Model = string(AppLoadLogV4.Model())
	apploadV4.CrashSDKVersion = string(AppLoadLogV4.CrashSDKVersion())
	apploadV4.AppVersion = string(AppLoadLogV4.AppVersion())
	apploadV4.PackageName = string(AppLoadLogV4.PackageName())
	apploadV4.UserKey = string(AppLoadLogV4.UserKey())
	apploadV4.DeviceId = string(AppLoadLogV4.DeviceId())
	apploadV4.Geo = string(AppLoadLogV4.Geo())
	apploadV4.City = string(AppLoadLogV4.City())
	apploadV4.ReportDatetime = AppLoadLogV4.ReportDatetime()
	apploadV4.MemoryWarning = string(AppLoadLogV4.MemoryWarning())
	apploadV4.Carrier = string(AppLoadLogV4.Carrier())
	apploadV4.SessionKey = string(AppLoadLogV4.SessionKey())
	apploadV4.Emulator = AppLoadLogV4.Emulator()
	apploadV4.NetworkKind = string(AppLoadLogV4.NetworkKind())
	apploadV4.VendorUserKey = string(AppLoadLogV4.VendorUserKey())
	apploadV4.VendorDeviceId = string(AppLoadLogV4.VendorDeviceId())
	apploadV4.ReceiveDatetime = string(time.Now().Format(time.RFC3339))
	apploadV4.buildCode = string(AppLoadLogV4.BuildCode())

	return apploadV4
}

type InitMessage struct {
	GameCode   string
	Os         string
	Log        *AppLoadV3
	LogVersion string
}

func InitMessageFromApploadV3(gameCode string, os string, log *AppLoadV3, logVersion string) *InitMessage {
	initMessage := &InitMessage{}
	initMessage.GameCode = gameCode
	initMessage.Os = os
	initMessage.Log = log
	initMessage.LogVersion = logVersion
	return initMessage
}
