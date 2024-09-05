package localization

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/fa_IR"
	ut "github.com/go-playground/universal-translator"
)

var (
	translationMap = make(map[string]map[string]string)
	translator     ut.Translator
)

func Register(request *http.Request) {
	var universalTranslator *ut.UniversalTranslator
	en := en.New()
	universalTranslator = ut.New(en, fa_IR.New())

	farsi, _ := universalTranslator.GetTranslator("fa_IR")
	farsiMap, _ := loadTranslations("localization/fa.json")
	for key, translation := range farsiMap {
		farsi.Add(key, translation, true)
	}

	english, _ := universalTranslator.GetTranslator("en")
	englishMap, _ := loadTranslations("localization/en.json")
	for key, translation := range englishMap {
		english.Add(key, translation, true)
	}

	locale := GetLocale(request)
	translator, _ = universalTranslator.GetTranslator(locale)
}

func GetTranslator() ut.Translator {
	return translator
}

func GetLocale(request *http.Request) string {
	return request.Header.Get("Accept-Language")
}

func loadTranslations(filePath string) (map[string]string, error) {
	if translations, ok := translationMap[filePath]; ok {
		return translations, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(bytes, &jsonData)
	if err != nil {
		return nil, err
	}

	flattenedTranslations := make(map[string]string)
	flattenMap("", jsonData, flattenedTranslations)

	translationMap[filePath] = flattenedTranslations

	return flattenedTranslations, nil
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
