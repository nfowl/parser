package sql

import (
	"fmt"
	"os"
	"strings"

	"nfowler.dev/siso-parser/pkg/siso"
)

func init() {
	siso.AvailableWriters.AddWriter("sql", &Writer{})
}

type Writer struct{}

func (t Writer) WriteEnum(d *siso.DisEnum) error {
	if _, ok := siso.InterestingEnums[d.Name]; ok {
		fmt.Println(d.Name)
	}
	return nil
}

//WriteCet Writes out the cet structure to a text file.
func (t Writer) WriteCet(c *siso.Cet) error {
	if !strings.Contains(c.Name, "Entity") {
		return nil
	}
	f, err := os.Create("cet.sql")
	defer f.Close()
	writeSQLHeader(f)
	writeCetTable(f)
	// f.WriteString(fmt.Sprintln("Country Kind Domain Category Subcategory Specific Extra Description"))
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
	writeSQLFooter(f)
	return err
}

func writeCetTable(f *os.File) {
	f.WriteString("CREATE TABLE \"entities\" (\n")
	f.WriteString("eid serial,\n")
	f.WriteString("\"country\" int2,\n")
	f.WriteString("\"kind\" int2,\n")
	f.WriteString("\"domain\" int2,\n")
	f.WriteString("\"category\" int2,\n")
	f.WriteString("\"subcategory\" int2,\n")
	f.WriteString("\"specific\" int2,\n")
	f.WriteString("\"extra\" int2,\n")
	f.WriteString("\"description\" varchar(100));\n")
	f.WriteString("ALTER TABLE \"entities\" ADD PRIMARY KEY (eid);\n")
}

func writeSQLFooter(f *os.File) {
	f.WriteString("COMMIT;\nANALYZE \"countries\";")
}

func writeSQLHeader(f *os.File) {
	f.WriteString("SET CLIENT_ENCODING TO UTF8;\nSET STANDARD_CONFORMING_STRINGS TO ON;\nBEGIN;")
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
	desc = strings.Replace(desc, "\"", "\\\"", -1)
	desc = strings.Replace(desc, "'", "\\'", -1)
	valuesString := fmt.Sprintf(
		"VALUES ('%d','%d','%d','%d','%d','%d','%d','%s');",
		e.Country,
		e.Kind,
		e.Domain,
		catValue,
		subCatValue,
		specificValue,
		extraValue,
		desc)
	tableString := "INSERT INTO \"entities\" (\"country\",\"kind\",\"domain\",\"category\",\"subcategory\",\"specific\",\"extra\",\"description\")"
	return fmt.Sprintf("%s %s\n", tableString, valuesString)
}
