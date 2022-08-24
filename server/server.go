package server

import (
	"ecommerceapi/lib"
	"ecommerceapi/server/routes/products"
	"ecommerceapi/server/routes/users"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

func getValidator() *lib.RequestValidator {
	// Create a new translator with english as the default language
	en_locale := en.New()
	universal_translator := ut.New(en_locale, en_locale)
	translator, _ := universal_translator.GetTranslator("en")

	// Create a request validator instance
	val := &lib.RequestValidator{Validator: validator.New(), Translator: &translator}

	// Register the english translations for the validator
	en_translations.RegisterDefaultTranslations(val.Validator, translator)

	// Use the json tag name for validation errors instead of the struct field name
	val.Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return val
}

func initRoutes(app *echo.Echo, db *bun.DB) {
	products.InitRoutes(app, db)
	users.InitRoutes(app, db)
}

func Init(port *string, db *bun.DB) *echo.Echo {
	// Create an instance of Echo
	server := echo.New()

	// Attach the custom validator to the echo instance
	server.Validator = getValidator()

	// Initialize the routes
	initRoutes(server, db)

	// Start the server
	server.Logger.Fatal(server.Start(":" + *port))

	return server
}
