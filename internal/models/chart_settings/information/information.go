package models_chart_settings_information

type Information struct {
	PriceOpenName     InformationData `gorm:"embedded;embeddedPrefix:price_open_name_" json:"priceOpenName"`
	PriceOpenValue    InformationData `gorm:"embedded;embeddedPrefix:price_open_value_" json:"priceOpenValue"`
	PriceHighName     InformationData `gorm:"embedded;embeddedPrefix:price_high_name_" json:"priceHighName"`
	PriceHighValue    InformationData `gorm:"embedded;embeddedPrefix:price_high_value_" json:"priceHighValue"`
	PriceLowName      InformationData `gorm:"embedded;embeddedPrefix:price_low_name_" json:"priceLowName"`
	PriceLowValue     InformationData `gorm:"embedded;embeddedPrefix:price_low_value_" json:"priceLowValue"`
	PriceCloseName    InformationData `gorm:"embedded;embeddedPrefix:price_close_name_" json:"priceCloseName"`
	PriceCloseValue   InformationData `gorm:"embedded;embeddedPrefix:price_close_value_" json:"priceCloseValue"`
	VolumeName        InformationData `gorm:"embedded;embeddedPrefix:volume_name_" json:"volumeName"`
	VolumeValue       InformationData `gorm:"embedded;embeddedPrefix:volume_value_" json:"volumeValue"`
	TradesName        InformationData `gorm:"embedded;embeddedPrefix:trades_name_" json:"tradesName"`
	TradesValue       InformationData `gorm:"embedded;embeddedPrefix:trades_value_" json:"tradesValue"`
	HorizontalPadding int             `json:"horizontalPadding"`
	VerticalPadding   int             `json:"verticalPadding"`
	HorizontalSpacing int             `json:"horizontalSpacing"`
	VerticalSpacing   int             `json:"verticalSpacing"`
	NameSpacing       int             `json:"nameSpacing"`
}

func LoadDefaultInformation() Information {
	return Information{
		PriceOpenName:     LoadDefaultInformationDataColor("открытие", "#000000ff"),
		PriceOpenValue:    LoadDefaultInformationData("", "#000000ff", "#008f01ff", "#d90400ff"),
		PriceHighName:     LoadDefaultInformationDataColor("максимум", "#000000ff"),
		PriceHighValue:    LoadDefaultInformationData("", "#000000ff", "#008f01ff", "#d90400ff"),
		PriceLowName:      LoadDefaultInformationDataColor("минимум", "#000000ff"),
		PriceLowValue:     LoadDefaultInformationData("", "#000000ff", "#008f01ff", "#d90400ff"),
		PriceCloseName:    LoadDefaultInformationDataColor("закрытие", "#000000ff"),
		PriceCloseValue:   LoadDefaultInformationData("", "#000000ff", "#008f01ff", "#d90400ff"),
		VolumeName:        LoadDefaultInformationDataColor("объем", "#000000ff"),
		VolumeValue:       LoadDefaultInformationData("", "#000000ff", "#008f01ff", "#d90400ff"),
		TradesName:        LoadDefaultInformationDataColor("количество сделок", "#000000ff"),
		TradesValue:       LoadDefaultInformationData("", "#000000ff", "#008f01ff", "#d90400ff"),
		HorizontalPadding: 10,
		VerticalPadding:   10,
		HorizontalSpacing: 20,
		VerticalSpacing:   15,
		NameSpacing:       5,
	}
}
