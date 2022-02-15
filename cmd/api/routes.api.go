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
	mux.Post("/api/reset-password", app.ResetPassword)

	// /api/admin is prefix for all the following routes
	mux.Route("/api/admin", func(mux chi.Router) {
		mux.Use(app.Auth)
		mux.Post("/virtual-terminal-succeeded", app.VirtualTerminalPaymentSucceeded)

		mux.Get("/sales", app.AllSales)
		mux.Get("/subscriptions", app.AllSubscriptions)

		mux.Get("/get-sale/{id}", app.GetSale)

		mux.Post("/refund", app.RefundCharge)
		mux.Post("/cancel-subscription", app.CancelSubscription)

		mux.Get("/users", app.AllUsers)
		mux.Get("/users/{id}", app.OneUser)

		mux.Put("/users/{id}", app.EditUser)
		mux.Delete("/users/{id}", app.DeleteUser)

	})

	return mux
}
