package cmd

import (
	"fmt"
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
	env := core.EnvLocal

	log := core.SetupLogger(env)
	slog.SetDefault(log)

	err := core.ParseFlags()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	// Initialize storage for each entity
	orderStore, err := repository.NewJSONStorage(filepath.Join(core.Dir, core.OrderFile))
	if err != nil {
		log.Debug(err.Error())
		os.Exit(1)
	}
	menuStore, err := repository.NewJSONStorage(filepath.Join(core.Dir, core.MenuFile))
	if err != nil {
		log.Debug(err.Error())
		os.Exit(1)
	}
	inventoryStore, err := repository.NewJSONStorage(filepath.Join(core.Dir, core.InventoryFile))
	if err != nil {
		log.Debug(err.Error())
		os.Exit(1)
	}

	// Initialize repositories with specific storage files
	orderRepo := repository.NewOrderRepository(orderStore, log)
	menuRepo := repository.NewMenuRepository(menuStore, log)
	inventoryRepo := repository.NewInventoryRepository(inventoryStore, log)

	// Initialize services
	orderService := service.NewOrderService(orderRepo, menuRepo, inventoryRepo, log)
	menuService := service.NewMenuService(menuRepo, log)
	inventoryService := service.NewInventoryService(inventoryRepo, log)

	// Initialize handlers
	orderHandler := handler.NewOrderHandler(orderService, menuService, inventoryService, log)
	menuHandler := handler.NewMenuHandler(menuService, log)
	inventoryHandler := handler.NewInventoryHandler(inventoryService, log)

	// Initialize router
	mux := handler.Routes(orderHandler, menuHandler, inventoryHandler)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", core.Port),
		Handler: mux,
	}

	log.Info(
		"starting http server",
		slog.String("Env", env),
		slog.String("addr", fmt.Sprintf("http://127.0.0.1:%d", core.Port)),
		slog.String("dir", core.Dir),
	)
	log.Debug(fmt.Sprint(srv.ListenAndServe()))
}
