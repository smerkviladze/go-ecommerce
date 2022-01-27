package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	mux.Post("/api/payment-intent", app.GetPaymentIntent)

	mux.Get("/api/widget/{id}", app.GetWidgetByID)

	mux.Post("/api/create-customer-and-subscribe-to-plan", app.CreateCustomerAndSubscribeToPlan)

	mux.Post("/api/authenticate", app.CreateAuthToken)
	mux.Post("/api/is-authenticated", app.CheckAuthentication)
	mux.Post("/api/forgot-password", app.SendPasswordResetEmail)

	mux.Route("/api/admin", func(mux chi.Router) { // /api/admin is prefix for all the following routes
		mux.Use(app.Auth)

		mux.Post("/virtual-terminal-succeeded", app.VirtualTerminalPaymentSucceeded)

	})

	return mux
}

// package main

// import (
// 	"net/http"

// 	"github.com/go-chi/chi/v5"
// )

// func (app *application) routes() http.Handler {
// 	mux := chi.NewRouter()

// 	mux.Use(func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
// 			(rw).Header().Set("Access-Control-Allow-Origin", "*")
// 			(rw).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 			(rw).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// 			if r.Method == http.MethodOptions {
// 				return
// 			}
// 			next.ServeHTTP(rw, r)
// 		})
// 	})

// 	mux.Post("/api/payment-intent", app.GetPaymentIntent)
// 	mux.Get("/api/payment-intent", func(rw http.ResponseWriter, r *http.Request) {
// 		rw.Write([]byte(""))
// 	})

// 	return mux
// }
