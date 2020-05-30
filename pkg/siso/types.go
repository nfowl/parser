package siso

// var AvailableWriters Writers

type Writers struct {
	writers map[string]Writer
}

func (w *Writers) AddWriter(name string, writer Writer) error {
	w.writers[name] = writer
	return nil
}

func (w *Writers) GetWriter(name string) (Writer, error) {
	r := w.writers[name]
	return r, nil
}

var AvailableWriters = &Writers{
	writers: make(map[string]Writer),
}

var InterestingEnums = map[string]bool{
	"DIS-PDU Type":                   true,
	"Force ID":                       true,
	"Entity Kind":                    true,
	"Platform Domain":                true,
	"Platform-Land Category":         true,
	"Platform-Air Category":          true,
	"Platform-Surface Category":      true,
	"Platform-Subsurface Category":   true,
	"Platform-Space Category":        true,
	"Munition Domain":                true,
	"Munition Category":              true,
	"Environmental Subcategory":      true,
	"Radio Category":                 true,
	"Radio Subcategory":              true,
	"Expendable-Air Category":        true,
	"Expendable-Surface Category":    true,
	"Expendable-Subsurface Category": true,
	"Sensor/Emitter Category":        true,
	"Country":                        true,
}

type Writer interface {
	WriteCet(*Cet) error
	WriteEnum(*DisEnum) error
}

type SISOFile struct {
	Enums    []DisEnum `xml:"enum"`
	Entities []Cet     `xml:"cet"`
}

type DisEnum struct {
	Name          string       `xml:"name,attr"`
	UID           int          `xml:"uid,attr"`
	Values        []EnumValues `xml:"enumrow"`
	Applicability string       `xml:"applicability,attr"`
}

type EnumValues struct {
	Value       int    `xml:"value,attr"`
	Description string `xml:"description,attr"`
}

type Cet struct {
	Name     string       `xml:"name,attr"`
	Entities []EntityType `xml:"entity"`
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
