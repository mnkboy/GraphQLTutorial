package settingsModels

import "encoding/xml"

type DBSettingsModel struct {
	XMLName  xml.Name  `xml:"DBSettings"`
	DataBase []DBModel `xml:"DB"`
}
