package siso

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func init() {
	AvailableWriters.AddWriter("sql", &SQLWriter{})
}

type SQLWriter struct{}

func (t SQLWriter) WriteEnum(d *DisEnum) error {
	if _, ok := InterestingEnums[d.Name]; ok {
		enumRouter(d)
	}
	return nil
}

func enumRouter(d *DisEnum) {

	if d.Name == "Entity Kind" {
		handleKinds(d)
	}
	if d.Name == "Country" {
		handleCountries(d)
	}
	if d.Name == "Force ID" {
		handleForces(d)
	}
	if d.Name == "DIS-PDU Type" {
		handlePdus(d)
	}
	if strings.Contains(d.Name, "Domain") {
		handleDomains(d)
	}
	if strings.Contains(d.Name, "Category") {
		handleCategories(d)
	}
}

func handleKinds(d *DisEnum) {
	f, _ := os.Create("kinds.sql")
	defer f.Close()
	writeSQLHeader(f)
	f.WriteString("CREATE TABLE \"kinds\" (\n")
	f.WriteString("\"value\" int2,\n")
	f.WriteString("\"description\" varchar(20));\n")
	f.WriteString("ALTER TABLE \"kinds\" ADD PRIMARY KEY (value);\n")
	insertCommand := "INSERT INTO \"kinds\" (\"value\",\"description\")"
	for _, v := range d.Values {
		f.WriteString(fmt.Sprintf("%s VALUES ('%d','%s');\n", insertCommand, v.Value, v.Description))
	}
	writeSQLFooter(f, "kinds")
}

func handleDomains(d *DisEnum) {
	_, err := os.Stat("domains.sql")
	if os.IsNotExist(err) {
		f, _ := os.Create("domains.sql")
		writeSQLHeader(f)
		f.WriteString("CREATE TABLE \"domains\" (\n")
		f.WriteString("\"did\" serial,\n")
		f.WriteString("\"kind\" int2,\n")
		f.WriteString("\"value\" int2,\n")
		f.WriteString("\"description\" varchar(20));\n")
		f.WriteString("ALTER TABLE \"domains\" ADD PRIMARY KEY (did);\n")
		f.Close()
	}
	f, _ := os.OpenFile("domains.sql", os.O_APPEND|os.O_WRONLY, 644)
	defer f.Close()
	kind := d.Applicability[:1]
	insertCommand := "INSERT INTO \"domains\" (\"kind\",\"value\",\"description\")"
	for _, v := range d.Values {
		f.WriteString(fmt.Sprintf("%s VALUES ('%s','%d','%s');\n", insertCommand, kind, v.Value, v.Description))
	}
	writeSQLFooter(f, "domains")
}

func handleCategories(d *DisEnum) {
	_, err := os.Stat("categories.sql")
	if os.IsNotExist(err) {
		f, _ := os.Create("categories.sql")
		writeSQLHeader(f)
		f.WriteString("CREATE TABLE \"categories\" (\n")
		f.WriteString("\"cid\" serial,\n")
		f.WriteString("\"kind\" int2,\n")
		f.WriteString("\"domain\" int2,\n")
		f.WriteString("\"value\" int2,\n")
		f.WriteString("\"description\" varchar(20));\n")
		f.WriteString("ALTER TABLE \"categories\" ADD PRIMARY KEY (cid);\n")
		f.Close()
	}
	f, _ := os.OpenFile("categories.sql", os.O_APPEND|os.O_WRONLY, 644)
	defer f.Close()
	info := strings.Split(d.Applicability, ".")
	kind := info[0]
	domain := info[1]
	if domain == "X" {
		domain = ""
	}
	insertCommand := "INSERT INTO \"categories\" (\"kind\",\"domain\",\"value\",\"description\")"
	for _, v := range d.Values {
		f.WriteString(fmt.Sprintf("%s VALUES ('%s','%s','%d','%s');\n", insertCommand, kind, domain, v.Value, v.Description))
	}
	writeSQLFooter(f, "categories")
}

func handleCountries(d *DisEnum) {
	f, _ := os.Create("countries.sql")
	defer f.Close()
	writeSQLHeader(f)
	f.WriteString("CREATE TABLE \"countries\" (\n")
	f.WriteString("\"id\" int2,\n")
	f.WriteString("\"name\" varchar(50),\n")
	f.WriteString("\"iso3\" varchar(3));\n")
	f.WriteString("ALTER TABLE \"countries\" ADD PRIMARY KEY (id);\n")
	insertCommand := "INSERT INTO \"countries\" (\"id\",\"name\",\"iso3\")"
	for _, v := range d.Values {
		re := regexp.MustCompile(`\([A-Z]{3}\)`)
		if index := re.FindStringIndex(v.Description); index != nil {
			// fmt.Printf("%s, %v\n", v.Description, index)
			name := v.Description[:index[0]]
			name = strings.ReplaceAll(name, "\"", "\"\"")
			name = strings.ReplaceAll(name, "'", "''")
			iso3 := strings.ReplaceAll(strings.ReplaceAll(v.Description[index[0]:index[1]], "(", ""), ")", "")
			f.WriteString(fmt.Sprintf("%s VALUES ('%d','%s','%s');\n", insertCommand, v.Value, name, iso3))
		}
	}
	writeSQLFooter(f, "countries")
}

func handleForces(d *DisEnum) {
	f, _ := os.Create("forces.sql")
	defer f.Close()
	writeSQLHeader(f)
	f.WriteString("CREATE TABLE \"forces\" (\n")
	f.WriteString("\"value\" int2,\n")
	f.WriteString("\"description\" varchar(20));\n")
	f.WriteString("ALTER TABLE \"forces\" ADD PRIMARY KEY (value);\n")
	insertCommand := "INSERT INTO \"forces\" (\"value\",\"description\")"
	for _, v := range d.Values {
		f.WriteString(fmt.Sprintf("%s VALUES ('%d','%s');\n", insertCommand, v.Value, v.Description))
	}
	writeSQLFooter(f, "forces")
}

func handlePdus(d *DisEnum) {
	f, _ := os.Create("pdus.sql")
	defer f.Close()
	writeSQLHeader(f)
	f.WriteString("CREATE TABLE \"pdus\" (\n")
	f.WriteString("\"value\" int2,\n")
	f.WriteString("\"description\" varchar(30));\n")
	f.WriteString("ALTER TABLE \"pdus\" ADD PRIMARY KEY (value);\n")
	insertCommand := "INSERT INTO \"pdus\" (\"value\",\"description\")"
	for _, v := range d.Values {
		f.WriteString(fmt.Sprintf("%s VALUES ('%d','%s');\n", insertCommand, v.Value, v.Description))
	}
	writeSQLFooter(f, "pdus")
}

//WriteCet Writes out the cet structure to a text file.
func (t SQLWriter) WriteCet(c *Cet) error {
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
	writeSQLFooter(f, "entities")
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

func writeSQLFooter(f *os.File, name string) {
	f.WriteString(fmt.Sprintf("COMMIT;\nANALYZE \"%s\";\n", name))
}

func writeSQLHeader(f *os.File) {
	f.WriteString("SET CLIENT_ENCODING TO UTF8;\nSET STANDARD_CONFORMING_STRINGS TO ON;\nBEGIN;\n")
}

func writeCategory(e EntityType, catValue int, desc string) string {
	return writeSubcategory(e, catValue, 0, desc)
}

func writeSubcategory(e EntityType, catValue int, subCatValue int, desc string) string {
	return writeSpecific(e, catValue, subCatValue, 0, desc)
}

func writeSpecific(e EntityType, catValue int, subCatValue int, specificValue int, desc string) string {
	return writeExtra(e, catValue, subCatValue, specificValue, 0, desc)
}

func writeExtra(e EntityType, catValue int, subCatValue int, specificValue int, extraValue int, desc string) string {
	desc = strings.ReplaceAll(desc, "\"", "\"\"")
	desc = strings.ReplaceAll(desc, "'", "''")
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
