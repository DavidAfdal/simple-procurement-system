package builder

import (
	"github.com/DavidAfdal/purchasing-systeam/internal/http/handler"
	"github.com/DavidAfdal/purchasing-systeam/internal/http/router"
	"github.com/DavidAfdal/purchasing-systeam/internal/repositories"
	"github.com/DavidAfdal/purchasing-systeam/internal/services"
	"github.com/DavidAfdal/purchasing-systeam/pkg/route"
	"github.com/DavidAfdal/purchasing-systeam/pkg/token"
	"gorm.io/gorm"
)

func BuildAppPublicRoutes(db *gorm.DB, tokenUseCase token.TokenUseCase) []*route.Route {
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo, tokenUseCase)
	userHandler := handler.NewUserHandler(userService)

	supplierRepo := repositories.NewSupplierRepo(db)
	supplierService := services.NewSupplierService(supplierRepo)
	supplierHandler := handler.NewSupplierHandler(supplierService)

	itemRepo := repositories.NewItemRepo(db)
	itemService := services.NewItemService(itemRepo)
	itemHandler := handler.NewItemHandler(itemService)

	handler := handler.NewHandler(userHandler, supplierHandler, itemHandler, nil)
	return router.AppPublicRoute(handler)
}

func BuildAppPrivateRoutes(db *gorm.DB, tokenUseCase token.TokenUseCase) []*route.Route {

	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo, tokenUseCase)
	userHandler := handler.NewUserHandler(userService)

	supplierRepo := repositories.NewSupplierRepo(db)
	supplierService := services.NewSupplierService(supplierRepo)
	supplierHandler := handler.NewSupplierHandler(supplierService)

	itemRepo := repositories.NewItemRepo(db)
	itemService := services.NewItemService(itemRepo)
	itemHandler := handler.NewItemHandler(itemService)

	purchasingRepo := repositories.NewPurchasingRepo(db)
	purchasingDetailRepo := repositories.NewPurchasingDetailRepo()
	purchasingService := services.NewPurchasingService(db, purchasingRepo, itemRepo, purchasingDetailRepo)
	purchasingHandler := handler.NewPurchasingHandler(purchasingService)

	handler := handler.NewHandler(userHandler, supplierHandler, itemHandler, purchasingHandler)
	return router.AppPrivateRoute(handler)
}
