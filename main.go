package main

import (
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	r := chi.NewRouter()

	uri, err := url.Parse("http://localhost:1323/api/v3")
	if err != nil {
		panic(err)
	}

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:1323/swagger/doc.json"),
		httpSwagger.BeforeScript(`const UrlMutatorPlugin = (system) => ({
  rootInjects: {
    setScheme: (scheme) => {
      const jsonSpec = system.getState().toJSON().spec.json;
      const schemes = Array.isArray(scheme) ? scheme : [scheme];
      const newJsonSpec = Object.assign({}, jsonSpec, { schemes });

      return system.specActions.updateJsonSpec(newJsonSpec);
    },
    setHost: (host) => {
      const jsonSpec = system.getState().toJSON().spec.json;
      const newJsonSpec = Object.assign({}, jsonSpec, { host });

      return system.specActions.updateJsonSpec(newJsonSpec);
    },
    setBasePath: (basePath) => {
      const jsonSpec = system.getState().toJSON().spec.json;
      const newJsonSpec = Object.assign({}, jsonSpec, { basePath });

      return system.specActions.updateJsonSpec(newJsonSpec);
    }
  }
});`),
		httpSwagger.Plugins([]string{"UrlMutatorPlugin"}),
		httpSwagger.UIConfig(map[string]string{
			"onComplete": fmt.Sprintf(`() => {
    window.ui.setScheme('%s');
    window.ui.setHost('%s');
    window.ui.setBasePath('%s');
  }`, uri.Scheme, uri.Host, uri.Path),
		}),
	))

	http.ListenAndServe(":1323", r)
}
