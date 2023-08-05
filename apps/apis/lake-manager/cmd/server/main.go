package main

import (
     // "log"
	"apps/apis/lake-manager/configs"
	"apps/apis/lake-manager/internal/entity"
	"apps/apis/lake-manager/internal/infra/database"
	"apps/apis/lake-manager/internal/infra/webserver/handlers"
     _ "apps/apis/lake-manager/docs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
     "github.com/go-chi/jwtauth"
     httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


// @title           Invest Sense Lake Manager API
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Fabio Caffarello
// @contact.email  fabio.caffarello@gmail.com

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
     if err != nil {
          panic(err)
     }
     db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
     if err != nil {
          panic(err)
     }
     db.AutoMigrate(&entity.User{}, &entity.Product{})

     productDB := database.NewProduct(db)
     productHandler := handlers.NewProductHandler(productDB)

     userDB := database.NewUser(db)
     userHandler := handlers.NewUserHandler(userDB)

     healthzHandler := handlers.NewHealthzHandler()

     r := chi.NewRouter()
     r.Use(middleware.Logger)
     r.Use(middleware.Recoverer)
     r.Use(middleware.WithValue("jwt", configs.TokenAuth))
     r.Use(middleware.WithValue("JwtExperesIn", configs.JWTExpireIn))
     r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

     r.Post("/users", userHandler.Create)
     r.Post("/users/jwt", userHandler.GetJWT)

     r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

     r.Get("/healthz", healthzHandler.Healthz)
	http.ListenAndServe(":8000", r)
}



