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
	default:
		return DeviceBrandOther
	}
}

const (
	DeviceBrandApple  DeviceBrand = "apple"
	DeviceBrandVIVO   DeviceBrand = "vivo"
	DeviceBrandHuawei DeviceBrand = "huawei"
	DeviceBrandXiaomi DeviceBrand = "xiaomi"
	DeviceBrandOPPO   DeviceBrand = "oppo"
	DeviceBrandOther  DeviceBrand = "other"
	DeviceBrandOVHM   DeviceBrand = ""      // 通投 Brand
	DeviceBrandHonor  DeviceBrand = "honor" // 荣耀 TODO 需要实现区分荣耀和华为
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
