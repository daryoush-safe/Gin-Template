package localization

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/go-playground/locales/en_US"
	"github.com/go-playground/locales/fa_IR"
	ut "github.com/go-playground/universal-translator"
)

var translationMap = make(map[string]map[string]string)

func GetTranslator(locale string) ut.Translator {
	universalTranslator := createUniversalTranslator()
	loadAndAddTranslations(universalTranslator)

	translator, found := universalTranslator.GetTranslator(locale)
	if !found {
		translator, _ = universalTranslator.GetTranslator("fa_IR")
	}

	return translator
}

func createUniversalTranslator() *ut.UniversalTranslator {
	en := en_US.New()
	fa := fa_IR.New()
	return ut.New(en, en, fa)
}

func loadAndAddTranslations(universalTranslator *ut.UniversalTranslator) {
	addTranslations("fa_IR", "src/localization/fa.json", universalTranslator)
	addTranslations("en_US", "src/localization/en.json", universalTranslator)
}

func addTranslations(locale, filePath string, universalTranslator *ut.UniversalTranslator) {
	translator, found := universalTranslator.GetTranslator(locale)
	if !found {
		panic(fmt.Sprintf("translator for locale %s not found", locale))
	}

	translations := loadTranslations(filePath)

	for key, translation := range translations {
		translator.Add(key, translation, true)
	}
}

func loadTranslations(filePath string) map[string]string {
	if translations, ok := translationMap[filePath]; ok {
		return translations
	}

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(bytes, &jsonData)
	if err != nil {
		panic(err)
	}

	flattenedTranslations := make(map[string]string)
	flattenMap("", jsonData, flattenedTranslations)

	translationMap[filePath] = flattenedTranslations

	return flattenedTranslations
}

func flattenMap(prefix string, input map[string]interface{}, output map[string]string) {
	for k, v := range input {
		fullKey := k
		if prefix != "" {
			fullKey = prefix + "." + k
		}
		switch value := v.(type) {
		case map[string]interface{}:
			flattenMap(fullKey, value, output)
		case string:
			output[fullKey] = value
		default:
			// Handle other types as needed, e.g., numbers, booleans, etc.
		}
	}
}
