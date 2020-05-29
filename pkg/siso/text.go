package siso

import (
	"fmt"
	"strings"
)

func init() {
	AvailableWriters.addWriter("text", &TextWriter{})
}

type TextWriter struct {
	Ayo string
}

func (t TextWriter) WriteEnum(d *DisEnum) error {
	// fmt.Println(d)
	return nil
}

func (t TextWriter) WriteCet(c *Cet) error {
	// fmt.Println(c.Name)
	if !strings.Contains(c.Name, "Entity") {
		return nil
	}
	for _, entity := range c.Entities {
		for _, category := range entity.Categories {
			writeCategory(entity, category.Value, category.Description)
			for _, subcategory := range category.Subcategories {
				writeSubcategory(entity, category.Value, subcategory.Value, subcategory.Description)
				for _, specific := range subcategory.Specifics {
					writeSpecific(entity, category.Value, subcategory.Value, specific.Value, specific.Description)
					for _, extra := range specific.Extras {
						writeExtra(entity, category.Value, subcategory.Value, specific.Value, extra.Value, extra.Description)
					}
				}
			}
		}
	}
	return nil
}

func writeCategory(e EntityType, catValue int, desc string) {
	writeSubcategory(e, catValue, 0, desc)
}

func writeSubcategory(e EntityType, catValue int, subCatValue int, desc string) {
	writeSpecific(e, catValue, subCatValue, 0, desc)
}

func writeSpecific(e EntityType, catValue int, subCatValue int, specificValue int, desc string) {
	writeExtra(e, catValue, subCatValue, specificValue, 0, desc)
}

func writeExtra(e EntityType, catValue int, subCatValue int, specificValue int, extraValue int, desc string) {
	fmt.Println(e.Country, e.Domain, e.Kind, catValue, subCatValue, specificValue, extraValue, desc)
}
