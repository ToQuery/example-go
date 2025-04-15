package gidevice

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	giDevice "github.com/electricbubble/gidevice"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"testing"
)

type DeviceDetail struct {
	DeviceName                string `json:"DeviceName,omitempty"`
	DeviceColor               string `json:"DeviceColor,omitempty"`
	DeviceClass               string `json:"DeviceClass,omitempty"`
	ProductVersion            string `json:"ProductVersion,omitempty"`
	ProductType               string `json:"ProductType,omitempty"`
	ProductName               string `json:"ProductName,omitempty"`
	ModelNumber               string `json:"ModelNumber,omitempty"`
	SerialNumber              string `json:"SerialNumber,omitempty"`
	SIMStatus                 string `json:"SIMStatus,omitempty"`
	PhoneNumber               string `json:"PhoneNumber,omitempty"`
	CPUArchitecture           string `json:"CPUArchitecture,omitempty"`
	ProtocolVersion           string `json:"ProtocolVersion,omitempty"`
	RegionInfo                string `json:"RegionInfo,omitempty"`
	TelephonyCapability       bool   `json:"TelephonyCapability,omitempty"`
	TimeZone                  string `json:"TimeZone,omitempty"`
	UniqueDeviceID            string `json:"UniqueDeviceID,omitempty"`
	WiFiAddress               string `json:"WiFiAddress,omitempty"`
	WirelessBoardSerialNumber string `json:"WirelessBoardSerialNumber,omitempty"`
	BluetoothAddress          string `json:"BluetoothAddress,omitempty"`
	BuildVersion              string `json:"BuildVersion,omitempty"`
}

func TestDevices(t *testing.T) {
	usbmux, err := giDevice.NewUsbmux()
	if err != nil {
		log.Fatalln(err)
	}

	devices, err := usbmux.Devices()
	if err != nil {
		log.Fatal(err)
	}

	for _, dev := range devices {
		log.Println(dev.Properties().SerialNumber, dev.Properties().ProductID, dev.Properties().DeviceID)
	}
}

func TestDeviceDetail(t *testing.T) {
	usbmux, err := giDevice.NewUsbmux()
	if err != nil {
		log.Fatal(err)
	}

	devices, err := usbmux.Devices()
	if err != nil {
		log.Fatal(err)
	}

	if len(devices) == 0 {
		log.Fatal("No Device")
	}

	d := devices[0]

	detail, err1 := d.GetValue("", "")
	if err1 != nil {
		fmt.Errorf("get %s device detail fail : %w", d.Properties().SerialNumber, err1)
	}

	data, _ := json.Marshal(detail)
	d1 := &DeviceDetail{}
	json.Unmarshal(data, d1)
	fmt.Println(d1)
}

func TestDeviceImages(t *testing.T) {
	usbmux, err := giDevice.NewUsbmux()
	if err != nil {
		log.Fatal(err)
	}

	devices, err := usbmux.Devices()
	if err != nil {
		log.Fatal(err)
	}

	if len(devices) == 0 {
		log.Fatal("No Device")
	}

	d := devices[0]

	imageSignatures, err := d.Images()
	if err != nil {
		log.Fatalln(err)
	}

	for i, imgSign := range imageSignatures {
		log.Printf("[%d] %s\n", i+1, base64.StdEncoding.EncodeToString(imgSign))
	}

	dmgPath := "/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/DeviceSupport/14.4/DeveloperDiskImage.dmg"
	signaturePath := "/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/DeviceSupport/14.4/DeveloperDiskImage.dmg.signature"

	err = d.MountDeveloperDiskImage(dmgPath, signaturePath)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestDeviceApp(t *testing.T) {
	usbmux, err := giDevice.NewUsbmux()
	if err != nil {
		log.Fatalln(err)
	}

	devices, err := usbmux.Devices()
	if err != nil {
		log.Fatalln(err)
	}

	if len(devices) == 0 {
		log.Fatalln("No Device")
	}

	d := devices[0]

	bundleID := "com.apple.Preferences"
	pid, err := d.AppLaunch(bundleID)
	if err != nil {
		log.Fatalln(err)
	}

	err = d.AppKill(pid)
	if err != nil {
		log.Fatalln(err)
	}

	runningProcesses, err := d.AppRunningProcesses()
	if err != nil {
		log.Fatalln(err)
	}

	for _, process := range runningProcesses {
		if process.IsApplication {
			log.Printf("%4d\t%-24s\t%-36s\t%s\n", process.Pid, process.Name, filepath.Base(process.RealAppName), process.StartDate)
		}
	}
}

func TestDeviceScreenshot(t *testing.T) {
	usbmux, err := giDevice.NewUsbmux()
	if err != nil {
		log.Fatalln(err)
	}

	devices, err := usbmux.Devices()
	if err != nil {
		log.Fatalln(err)
	}

	if len(devices) == 0 {
		log.Fatalln("No Device")
	}

	d := devices[0]

	raw, err := d.Screenshot()
	if err != nil {
		log.Fatalln(err)
	}

	img, format, err := image.Decode(raw)
	if err != nil {
		log.Fatalln(err)
	}
	userHomeDir, _ := os.UserHomeDir()
	file, err := os.Create(userHomeDir + "/Downloads/s1." + format)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() { _ = file.Close() }()
	switch format {
	case "png":
		err = png.Encode(file, img)
	case "jpeg":
		err = jpeg.Encode(file, img, nil)
	}
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(file.Name())
}
