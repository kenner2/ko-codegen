package main

import (
	"fmt"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"ko-codegen/properties"
	"ko-codegen/templates"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const (
	outputDir        = "out"
	schemaDir        = "json-schema"
	schemaExtPattern = "*.schema.json"
)

func main() {
	fmt.Println("|-----------------------|")
	fmt.Println("| OpenKO Code Generator |")
	fmt.Println("|-----------------------|")

	// Read and compile all schema files in json-schema
	validSchemas, err := loadSchemas()
	if err != nil {
		fmt.Println(err)
		return
	}

	// generate c++ source for all the schemas
	err = generateCpp(validSchemas)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("OpenKO Code Generator completed successfully")
}

// loadSchemas reads and compiles all json-schema/*.schema.json files
// then returns a slice of all valid schemas
func loadSchemas() (validSchemas []jsonschema.Schema, err error) {
	fmt.Println("creating json-schema compiler with ExtractAnnotations=true")
	compiler := jsonschema.NewCompiler()
	compiler.ExtractAnnotations = true

	fmt.Println("reading schema file names from: " + schemaDir)
	fileNames, err := getSchemaFileNames()
	if err != nil {
		err = fmt.Errorf("failed to read schema file names: %w", err)
		return validSchemas, err
	}
	fmt.Println(fmt.Sprintf("found %d schema files", len(fileNames)))

	for i := range fileNames {
		fmt.Print(fmt.Sprintf("compiling schema file: %s", fileNames[i]))
		sch, err := compiler.Compile(fileNames[i])
		if err != nil {
			err = fmt.Errorf("failed to compile schema file: %w", err)
			return validSchemas, err
		}

		fmt.Println(" ...done")
		validSchemas = append(validSchemas, *sch)
	}

	return validSchemas, nil
}

// generateCpp generates c++ code files for each schema passed in, and writes the result to the output directory.
func generateCpp(validSchemas []jsonschema.Schema) (err error) {
	// the identifier is used to correctly assign types to each property based on constraints
	// and other properties
	identifier := properties.CppIdentifier{}

	for i := range validSchemas {
		fmt.Print(fmt.Sprintf("generating c++ for: %s", validSchemas[i].Title))

		// the template is an interface implementation that allows us to
		// structure and generate a code file
		template := templates.CppTemplate{}
		tableName, err := getTableNameFromId(validSchemas[i].Location)
		template.SetClassName(tableName)
		template.SetClassDesc(validSchemas[i].Description)

		// always include string since we're always generating a GetTableName func
		template.AddInclude("string")

		// Generate a static GetTableName() func
		tblNameDef := templates.MethodDefinition{
			ReturnType:  "static std::string",
			Name:        "GetTableName",
			Body:        fmt.Sprintf("return \"%s\";", strings.ToUpper(tableName)),
			Description: "Returns the database table name",
		}
		template.AddMethod(tblNameDef)

		// Setup properties and property-based funcs
		for key, _ := range validSchemas[i].Properties {
			// translate the json-schema type to a C++ type
			// i.e., a number type with a min:max of 0:255 will return uint8_t
			isOptional := !slices.Contains(validSchemas[i].Required, key)
			_type, tErr := identifier.GetType(key, *validSchemas[i].Properties[key], isOptional)
			if tErr != nil {
				err = fmt.Errorf("%s:%s failed to get type for property: %w", tableName, key, tErr)
				return err
			}

			if strings.HasSuffix(_type, "_t") {
				template.AddInclude("cstdint")
			}

			if strings.HasPrefix(_type, "Nullable") {
				template.AddInclude("cgHelpers/nullable.h")
				template.AddUsing("cgHelpers::Nullable")
			}

			// Setup property definition and add to template
			propDef := templates.PropertyDefinition{
				Type:        _type,
				Name:        key,
				Description: validSchemas[i].Properties[key].Description,
			}
			template.AddProperty(propDef)

			// Setup static column-name funcs
			colNameDef := templates.MethodDefinition{
				ReturnType:  "static std::string",
				Name:        "CN" + "_" + key,
				Body:        fmt.Sprintf("return \"%s\";", key),
				Description: "Returns the database column name",
			}
			template.AddMethod(colNameDef)
		}

		// generate template
		templateStr, tErr := template.Generate()
		if tErr != nil {
			err = fmt.Errorf("%s failed to generate c++ source: %w", tableName, tErr)
			return err
		}

		// write the template to a file
		outFile := filepath.Join(outputDir, template.GetFileName())
		if fErr := writeToFile(outFile, templateStr); fErr != nil {
			err = fmt.Errorf("failed to write file %s: %w", outFile, fErr)
			return err
		}
		fmt.Println(fmt.Sprintf("... written to: %s", outFile))
	}

	fmt.Println("c++ code generated successfully")
	return nil
}

func writeToFile(filename string, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func getSchemaFileNames() (fileNames []string, err error) {
	return filepath.Glob(filepath.Join(schemaDir, schemaExtPattern))
}

func getTableNameFromId(id string) (string, error) {
	// id format:   "$id": "https://github.com/srmeier/KnightOnline/json-schema/account_char.schema.json",
	slashSplit := strings.Split(id, "/")
	if len(slashSplit) == 0 {
		return "", fmt.Errorf("failed to get table name from id: %s", id)
	}
	dotSplit := strings.Split(slashSplit[len(slashSplit)-1], ".")
	if len(dotSplit) == 0 {
		return "", fmt.Errorf("failed to get table name from id: %s", id)
	}

	return dotSplit[0], nil
}
