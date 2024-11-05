package cmd

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ab-dauletkhan/hot-coffee/internal/core"
	"github.com/ab-dauletkhan/hot-coffee/internal/handler"
	"github.com/ab-dauletkhan/hot-coffee/internal/repository"
	"github.com/ab-dauletkhan/hot-coffee/internal/service"
)

func Start() {
	err := core.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}

	log := core.SetupLogger(core.Env)
	slog.SetDefault(log)

	log.Info("application started",
		"version", "1.0.0",
		"environment", core.Env,
	)

	// Initialize storage for each entity
	inventoryStorage, err := repository.NewJSONStorage(filepath.Join(core.Dir, core.InventoryFile))
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	menuStorage, err := repository.NewJSONStorage(filepath.Join(core.Dir, core.MenuFile))
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	orderStorage, err := repository.NewJSONStorage(filepath.Join(core.Dir, core.OrderFile))
	if err != nil {
		log.Debug(err.Error())
		os.Exit(1)
	}

	// Initialize repositories with specific storage files
	inventoryRepo := repository.NewInventoryRepository(inventoryStorage, log)
	menuRepo := repository.NewMenuRepository(menuStorage, log)
	orderRepo := repository.NewOrderRepository(orderStorage, log)

	// Initialize services
	inventoryService := service.NewInventoryService(inventoryRepo, log)
	menuService := service.NewMenuService(menuRepo, inventoryService, log)
	orderService := service.NewOrderService(orderRepo, menuService, inventoryService, log)

	// Initialize handlers
	inventoryHandler := handler.NewInventoryHandler(inventoryService, log)
	menuHandler := handler.NewMenuHandler(menuService, log)
	orderHandler := handler.NewOrderHandler(orderService, menuService, inventoryService, log)

	// Initialize router
	mux := handler.Routes(orderHandler, menuHandler, inventoryHandler)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", core.Port),
		Handler: mux,
	}

	log.Info(
		"starting http server",
		slog.String("Env", core.Env),
		slog.String("addr", fmt.Sprintf("http://127.0.0.1:%d", core.Port)),
		slog.String("dir", core.Dir),
	)
	log.Debug(fmt.Sprint(srv.ListenAndServe()))
}
