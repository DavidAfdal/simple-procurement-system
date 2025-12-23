package router

import (
	"github.com/DavidAfdal/purchasing-systeam/internal/http/handler"
	"github.com/DavidAfdal/purchasing-systeam/pkg/route"
)

func AppPublicRoute(handler *handler.Handler) []*route.Route {
	userHandler := handler.UserHandler
	supplierHandler := handler.SupplierHandler

	return []*route.Route{
		{
			Method:  "POST",
			Path:    "/users/register",
			Handler: userHandler.Register,
		},
		{
			Method:  "POST",
			Path:    "/users/login",
			Handler: userHandler.Login,
		},
		{
			Method:  "GET",
			Path:    "/suppliers",
			Handler: supplierHandler.GetSuppliers,
		},
		{
			Method:  "GET",
			Path:    "/suppliers/:id",
			Handler: supplierHandler.GetSupplier,
		},
	}
}

func AppPrivateRoute(handler *handler.Handler) []*route.Route {

	// userHandler := handler.UserHandler
	// supplierHandler := handler.SupplierHandler
	purchasingHandler := handler.PurchasingHandler

	return []*route.Route{
		{
			Method:  "POST",
			Path:    "/purchasings",
			Handler: purchasingHandler.CreatePurchasing,
		},
	}
}
