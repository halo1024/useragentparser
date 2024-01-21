package useragentparser

import (
	"regexp"
)

type DeviceBrand string

func (d DeviceBrand) String() string {
	return string(d)
}

func ParseBrand(brand string) DeviceBrand {
	switch brand {
	case DeviceBrandApple.String():
		return DeviceBrandApple
	case DeviceBrandVIVO.String():
		return DeviceBrandVIVO
	case DeviceBrandHuawei.String():
		return DeviceBrandHuawei
	case DeviceBrandXiaomi.String():
		return DeviceBrandXiaomi
	case DeviceBrandOPPO.String():
		return DeviceBrandOPPO
	case DeviceBrandOVHM.String():
		return DeviceBrandOVHM
	case DeviceBrandHonor.String():
		return DeviceBrandHonor
	case DeviceBrandGOOGLE.String():
		return DeviceBrandGOOGLE
	case DeviceBrandSAMSUNG.String():
		return DeviceBrandSAMSUNG
	case DeviceBrandMOTOROLA.String():
		return DeviceBrandMOTOROLA
	case DeviceBrandLENOVO.String():
		return DeviceBrandLENOVO
	default:
		return DeviceBrandOther
	}
}

const (
	DeviceBrandApple    DeviceBrand = "apple"
	DeviceBrandVIVO     DeviceBrand = "vivo"
	DeviceBrandHuawei   DeviceBrand = "huawei"
	DeviceBrandXiaomi   DeviceBrand = "xiaomi"
	DeviceBrandOPPO     DeviceBrand = "oppo"
	DeviceBrandHonor    DeviceBrand = "honor" // 荣耀 TODO 需要实现区分荣耀和华为
	DeviceBrandGOOGLE   DeviceBrand = "google"
	DeviceBrandSAMSUNG  DeviceBrand = "samsung"
	DeviceBrandMOTOROLA DeviceBrand = "motorola"
	DeviceBrandLENOVO   DeviceBrand = "lenovo"
	DeviceBrandMEIZU    DeviceBrand = "meizu"
	DeviceBrandOther    DeviceBrand = "other"
	DeviceBrandOVHM     DeviceBrand = "" // 通投 Brand
)

type Device struct {
	Brand DeviceBrand
	Model string
}

type deviceParser struct {
	Reg          *regexp.Regexp
	Expr         string               `yaml:"regex"`
	Brand        DeviceBrand          `yaml:"brand"`
	ModelParsers []*deviceModelParser `yaml:"models"`
}

func (parser *deviceParser) Match(userAgentString string, device *Device) {
	// FindStringSubmatchIndex when find models
	matches := parser.Reg.FindStringIndex(userAgentString)

	if len(matches) == 0 {
		return
	}

	//for _, modelParser := range parser.ModelParsers {
	//	modelMatches := modelParser.Reg.FindStringSubmatchIndex(userAgentString)
	//	if len(modelMatches) > 0 {
	//		device.Model = string(parser.Reg.ExpandString(nil, modelParser.Model, userAgentString, matches))
	//		device.Model = strings.TrimSpace(device.Model)
	//
	//		break
	//	}
	//}
	device.Brand = parser.Brand
}
