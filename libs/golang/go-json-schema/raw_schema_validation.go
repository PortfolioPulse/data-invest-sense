package gojsonschema

import (
	"fmt"
	"log"

	"github.com/xeipuuv/gojsonschema"
)

// Define the JSON Schema
// schemaContent := `{
//      "$id": "https://example.com/person.schema.json",
//      "$schema": "https://json-schema.org/draft/2020-12/schema",
//      "title": "Person",
//      "type": "object",
//      "properties": {
//           "firstName": {
//                "type": "string",
//                "description": "The person's first name."
//           },
//           "lastName": {
//                "type": "string",
//                "description": "The person's last name."
//           },
//           "age": {
//                "description": "Age in years which must be equal to or greater than zero.",
//                "type": "integer",
//                "minimum": 0
//           }
//      }
// }`

// func ValidSchema(schemaContent string) bool {

// 	// Parse the JSON Schema
// 	schemaLoader := gojsonschema.NewStringLoader(schemaContent)

// 	// Validate the JSON Schema
// 	result, err := gojsonschema.Validate(schemaLoader)
// 	if err != nil {
// 		log.Fatalf("Error validating schema: %v", err)
// 	}

// 	// Check if the schema is valid
// 	if result.Valid() {
// 		return true
// 	} else {
// 		fmt.Println("Validation errors:")
// 		for _, desc := range result.Errors() {
// 			fmt.Printf("- %s\n", desc)
// 		}
// 		return false
// 	}
// }
