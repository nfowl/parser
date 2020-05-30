package text

import (
	"fmt"
	"os"
	"strings"

	"nfowler.dev/siso-parser/pkg/siso"
)

func init() {
	siso.AvailableWriters.AddWriter("text", &TextWriter{})
}

type TextWriter struct{}

func (t TextWriter) WriteEnum(d *siso.DisEnum) error {
	if _, ok := siso.InterestingEnums[d.Name]; ok {
		fmt.Println(d.Name)
	}
	return nil
}

//WriteCet Writes out the cet structure to a text file.
func (t TextWriter) WriteCet(c *siso.Cet) error {
	if !strings.Contains(c.Name, "Entity") {
		return nil
	}
	f, err := os.Create("cet.sql")
	defer f.Close()
	f.WriteString(fmt.Sprintln("Country Kind Domain Category Subcategory Specific Extra Description"))
	for _, entity := range c.Entities {
		for _, category := range entity.Categories {
			f.WriteString(writeCategory(entity, category.Value, category.Description))
			for _, subcategory := range category.Subcategories {
				f.WriteString(writeSubcategory(entity, category.Value, subcategory.Value, subcategory.Description))
				for _, specific := range subcategory.Specifics {
					f.WriteString(writeSpecific(entity, category.Value, subcategory.Value, specific.Value, specific.Description))
					for _, extra := range specific.Extras {
						f.WriteString(writeExtra(entity, category.Value, subcategory.Value, specific.Value, extra.Value, extra.Description))
					}
				}
			}
		}
	}
	return err
}

func writeCategory(e siso.EntityType, catValue int, desc string) string {
	return writeSubcategory(e, catValue, 0, desc)
}

func writeSubcategory(e siso.EntityType, catValue int, subCatValue int, desc string) string {
	return writeSpecific(e, catValue, subCatValue, 0, desc)
}

func writeSpecific(e siso.EntityType, catValue int, subCatValue int, specificValue int, desc string) string {
	return writeExtra(e, catValue, subCatValue, specificValue, 0, desc)
}

func writeExtra(e siso.EntityType, catValue int, subCatValue int, specificValue int, extraValue int, desc string) string {
	return fmt.Sprintln(
		e.Country,
		e.Kind,
		e.Domain,
		catValue,
		subCatValue,
		specificValue,
		extraValue,
		desc)
}
