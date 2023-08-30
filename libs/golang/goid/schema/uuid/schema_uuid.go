package encryptid

import (
    "crypto/sha256"
    "encoding/json"
    "fmt"

    "github.com/google/uuid"
)

type ID = string

// GenerateSchemaID generates a UUID-based schema ID based on schemaType and properties
func GenerateSchemaID(schemaType string, properties map[string]interface{}) (string, error) {
    serializedSchemaProperties, err := json.Marshal(properties)
    if err != nil {
        return "", fmt.Errorf("error marshaling schema properties: %w", err)
    }

    schemaHash := hashSchema(serializedSchemaProperties)
    schemaID := generateUUIDFromHash(schemaHash, schemaType)

    return ID(schemaID), nil
}

func hashSchema(data []byte) []byte {
    hash := sha256.Sum256(data)
    return hash[:]
}

func generateUUIDFromHash(hash []byte, schemaType string) string {
    combinedData := append(hash, []byte(schemaType)...)
    combinedHash := sha256.Sum256(combinedData)
    return uuid.NewSHA1(uuid.Nil, combinedHash[:]).String()
}
