package siso

type SISOFile struct {
	Enums    []DisEnum `xml:"enum"`
	Entities Cet       `xml:"cet"`
}

type DisEnum struct {
	Name   string       `xml:"name,attr"`
	UID    int          `xml:"uid,attr"`
	Values []EnumValues `xml:"enumrow"`
}

type EnumValues struct {
	Value       int    `xml:"value,attr"`
	Description string `xml:"description,attr"`
}

type Cet struct {
	Name     string       `xml:"name,attr"`
	Entities []EntityType `xml:"entity`
}

type Category struct {
	Description   string        `xml:"description,attr"`
	Value         int           `xml:"value,attr"`
	Subcategories []Subcategory `xml:"subcategory"`
}

type Subcategory struct {
	Description string     `xml:"description,attr"`
	Value       int        `xml:"value,attr"`
	Specifics   []Specific `xml:"specific"`
}

type Specific struct {
	Description string  `xml:"description,attr"`
	Value       int     `xml:"value,attr"`
	Extras      []Extra `xml:"extra"`
}

type Extra struct {
	Description string `xml:"description,attr"`
	Value       int    `xml:"value,attr"`
}

type EntityType struct {
	Domain     int        `xml:"domain,attr"`
	Kind       int        `xml:"kind,attr"`
	Country    int        `xml:"country,attr"`
	Categories []Category `xml:"category"`
}
