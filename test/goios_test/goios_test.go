package goios_test

import (
	"github.com/danielpaulus/go-ios/ios"
	"github.com/danielpaulus/go-ios/ios/deviceinfo"
	"github.com/danielpaulus/go-ios/ios/instruments"
	"log"
	"testing"
)

func TestDeviceInfo(t *testing.T) {
	devices, err := ios.ListDevices()
	if err != nil {
		log.Printf("获取设备列表失败: %v\n", err)
	}

	deviceInfo, err := deviceinfo.NewDeviceInfo(devices.DeviceList[0])
	log.Printf("获取设备列表失败: %+v\n", deviceInfo)

	deviceInfoService, err := instruments.NewDeviceInfoService(devices.DeviceList[0])
	if err != nil {
		log.Printf("获取设备信息失败: %v\n", err)
	}
	defer deviceInfoService.Close()

	hardwareInfo, err := deviceInfoService.HardwareInformation()
	if err != nil {
		log.Printf("获取设备信息失败: %v\n", err)
	}
	log.Printf("获取设备列表失败: %v\n", hardwareInfo)
}
