package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

const (
	valuesFile  = "https://sourceware.org/git/?p=lvm2.git;a=blob_plain;f=lib/report/values.h;hb=HEAD"
	columnsFile = "https://sourceware.org/git/?p=lvm2.git;a=blob_plain;f=lib/report/columns.h;hb=HEAD"
)

func main() {
	pkgName := flag.String("pkg", "collector", "package name")
	output := flag.String("o", "", "output")
	flag.Parse()

	err := run(*pkgName, *output)
	if err != nil {
		log.Fatal(err)
	}
}

func run(pkgName, output string) error {
	values, columns, err := downloadFiles()
	if err != nil {
		return err
	}

	fields, err := parseFields(columns)
	if err != nil {
		return err
	}

	enums, err := parseEnums(values)
	if err != nil {
		return err
	}
	var enumsKey []string
	for key := range enums {
		enumsKey = append(enumsKey, key)
	}
	sort.Strings(enumsKey)

	var w bytes.Buffer

	fmt.Fprintf(&w, "// Code generated by fields_generator.go. DO NOT EDIT.\n")
	fmt.Fprintf(&w, "package %s\n", pkgName)

	fmt.Fprintf(&w, `
type FieldType string

const (
	FieldTypeSTR      FieldType = "STR"
	FieldTypeNUM      FieldType = "NUM"
	FieldTypeBIN      FieldType = "BIN"
	FieldTypeSIZ      FieldType = "SIZ"
	FieldTypePCT      FieldType = "PCT"
	FieldTypeTIM      FieldType = "TIM"
	FieldTypeSNUM     FieldType = "SNUM"
	FieldTypeSTR_LIST FieldType = "STR_LIST"
)

type Field struct {
	Type        FieldType
	Name        string
	Description string
}

`)
	fmt.Fprintf(&w, "var lvmFields = map[string]Field{\n")
	for _, field := range fields {
		fmt.Fprintf(&w, `	"%s": Field{
		Type:        FieldType%s,
		Name:        "%s",
		Description: "%s",
	},
`,
			field.ID,
			field.Type,
			field.Name,
			field.Description,
		)
	}

	fmt.Fprintf(&w, "}\n\n")

	//

	fmt.Fprintf(&w, "var lvmEnums = map[string][]string{\n")
	for _, key := range enumsKey {
		fmt.Fprintf(&w, `	"%s": []string{`, key)
		for i, v := range enums[key] {
			fmt.Fprintf(&w, `"%s"`, v)
			if i != len(enums[key])-1 {
				fmt.Fprint(&w, `, `)
			}
		}
		fmt.Fprintf(&w, "},\n")
	}
	fmt.Fprintf(&w, "}\n\n")

	//

	if output == "" {
		fmt.Println(w.String())
		return nil
	}
	return ioutil.WriteFile(output, w.Bytes(), 0644)
}

type Enum map[string][]string

func parseEnums(b []byte) (Enum, error) {
	enums := make(Enum)
	s := bufio.NewScanner(bytes.NewReader(b))
	for s.Scan() {
		line := s.Text()
		if !strings.HasPrefix(line, "FIELD_RESERVED_VALUE(NAMED, ") {
			continue
		}
		line = strings.TrimSuffix(strings.TrimPrefix(line, "FIELD_RESERVED_VALUE(NAMED, "), ")")
		split := strings.Split(line, ", ")

		enums[strings.TrimSpace(split[0])] = append(enums[strings.TrimSpace(split[0])], strings.Trim(strings.TrimSpace(split[3]), `"`))
	}
	err := s.Err()
	if err != nil {
		return nil, fmt.Errorf("scan: %w", err)
	}

	return enums, nil
}

type Field struct {
	Type        string
	ID          string
	Name        string
	Description string
}

func parseFields(b []byte) ([]Field, error) {
	var fields []Field
	s := bufio.NewScanner(bytes.NewReader(b))
	for s.Scan() {
		line := s.Text()
		if !strings.HasPrefix(line, "FIELD(") {
			continue
		}
		line = strings.TrimSuffix(strings.TrimPrefix(line, "FIELD("), ")")
		split := strings.SplitN(line, ", ", 9)

		fields = append(fields, Field{
			Type:        strings.TrimSpace(split[2]),
			ID:          strings.TrimSpace(split[7]),
			Name:        strings.Trim(strings.TrimSpace(split[3]), `"`),
			Description: strings.Trim(strings.TrimSpace(strings.Split(split[8], `",`)[0]), `"`),
		})
	}
	err := s.Err()
	if err != nil {
		return nil, fmt.Errorf("scan: %w", err)
	}

	return fields, nil
}

func downloadFiles() (values, columns []byte, err error) {
	{
		respV, err := http.Get(valuesFile)
		if err != nil {
			return nil, nil, fmt.Errorf("values: %w", err)
		}
		defer respV.Body.Close()

		values, err = ioutil.ReadAll(respV.Body)
		if err != nil {
			return nil, nil, err
		}
	}

	{
		respC, err := http.Get(columnsFile)
		if err != nil {
			return nil, nil, fmt.Errorf("columns: %w", err)
		}
		defer respC.Body.Close()

		columns, err = ioutil.ReadAll(respC.Body)
		if err != nil {
			return nil, nil, err
		}
	}

	return values, columns, nil
}
