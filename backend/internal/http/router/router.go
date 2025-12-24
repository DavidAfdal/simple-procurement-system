package router

import (
	"github.com/DavidAfdal/purchasing-systeam/internal/http/handler"
	"github.com/DavidAfdal/purchasing-systeam/pkg/route"
)

const (
	Admin = "Admin"
	Staff = "Staff"
)

var (
	allRoles  = []string{Admin, Staff}
	onlyAdmin = []string{Admin}
	onlyStaff = []string{Staff}
)

func AppPublicRoute(handler *handler.Handler) []*route.Route {
	userHandler := handler.UserHandler
	supplierHandler := handler.SupplierHandler
	itemHandler := handler.ItemHandler

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
		{
			Method:  "GET",
			Path:    "/items",
			Handler: itemHandler.GetItems,
		},
		{
			Method:  "GET",
			Path:    "/items/:id",
			Handler: itemHandler.GetItemByID,
		},
	}
}

func AppPrivateRoute(handler *handler.Handler) []*route.Route {
	userHandler := handler.UserHandler
	supplierHandler := handler.SupplierHandler
	itemHandler := handler.ItemHandler
	purchasingHandler := handler.PurchasingHandler

	return []*route.Route{
		{
			Method:  "POST",
			Path:    "/users/logout",
			Handler: userHandler.Logout,
			Roles:   allRoles,
		},
		{
			Method:  "POST",
			Path:    "/suppliers",
			Handler: supplierHandler.CreateSupplier,
			Roles:   allRoles,
		},
		{
			Method:  "PUT",
			Path:    "/suppliers/:id",
			Handler: supplierHandler.UpdateSupplier,
			Roles:   allRoles,
		},
		{
			Method:  "DELETE",
			Path:    "/suppliers/:id",
			Handler: supplierHandler.DeleteSupplier,
			Roles:   allRoles,
		},
		{
			Method:  "POST",
			Path:    "/items",
			Handler: itemHandler.CreateItem,
			Roles:   allRoles,
		},
		{
			Method:  "PUT",
			Path:    "/items/:id",
			Handler: itemHandler.UpdateItem,
			Roles:   allRoles,
		},
		{
			Method:  "DELETE",
			Path:    "/items/:id",
			Handler: itemHandler.DeleteItem,
			Roles:   allRoles,
		},
		{
			Method:  "POST",
			Path:    "/purchasings",
			Handler: purchasingHandler.CreatePurchasing,
			Roles:   allRoles,
		},
		{
			Method:  "GET",
			Path:    "/purchasings",
			Handler: purchasingHandler.GetPurchasing,
			Roles:   allRoles,
		},
		{
			Method:  "GET",
			Path:    "/purchasings/me ",
			Handler: purchasingHandler.GetMyPurchasings,
			Roles:   allRoles,
		},
	}
}
