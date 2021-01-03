package settingsModels

import "encoding/xml"

type DBModel struct {
	XMLName  xml.Name `xml:"DB"`
	Name     string   `xml:"name,attr"`
	Engine   string   `xml:"Engine"`
	Server   string   `xml:"Server"`
	Port     string   `xml:"Port"`
	User     string   `xml:"User"`
	Password string   `xml:"Password"`
	Database string   `xml:"Database"`
	SslMode  string   `xml:"SslMode"`
}
