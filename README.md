# Hot-Coffee: Coffee Shop Management System

A robust backend system for managing coffee shop operations, built with pure Go and focusing on clean architecture principles.

## Project Overview

Hot-Coffee is a REST API-based management system that handles:
- Order processing and tracking
- Inventory management
- Menu item administration
- Sales reporting and analytics

### Technical Highlights

- **Pure Go Implementation**: Built using only Go standard library, demonstrating deep language expertise
- **Three-Layer Architecture**:
  - Presentation Layer (Handlers)
  - Business Logic Layer (Services)
  - Data Access Layer (Repositories)
- **Optimized Performance**: Implemented O(1) removal operations
- **Persistent Storage**: Custom JSON-based data management system
- **Comprehensive Logging**: Integrated slog package for system monitoring

### Key Features

- RESTful API endpoints for orders, menu items, and inventory management
- Real-time inventory tracking and updates
- Automated ingredient deduction upon order processing
- Sales and popularity reporting
- Structured error handling with appropriate HTTP status codes

### System Architecture

 .
├──  cmd
│   └──  main.go
├──  go.mod
├──  internal
│   ├──  core
│   │   ├──  consts.go
│   │   ├──  flag.go
│   │   └──  slog.go
│   ├──  handler
│   │   ├──  handler.go
│   │   ├──  inventory.go
│   │   ├──  menu.go
│   │   ├──  order.go
│   │   ├──  report.go
│   │   └──  routes.go
│   ├──  repository
│   │   ├──  inventory.go
│   │   ├──  json_store.go
│   │   ├──  menu.go
│   │   ├──  order.go
│   │   └──  report.go
│   └──  service
│       ├──  inventory.go
│       ├──  menu.go
│       ├──  order.go
│       └──  report.go
├──  main.go
├──  Makefile
├──  models
│   ├──  inventory.go
│   ├──  menu.go
│   ├──  order.go
│   └──  report.go
└──  README.md
```

### API Endpoints

#### Orders
- `POST /orders` - Create new order
- `GET /orders` - Retrieve all orders
- `GET /orders/{id}` - Retrieve specific order
- `PUT /orders/{id}` - Update order
- `DELETE /orders/{id}` - Delete order
- `POST /orders/{id}/close` - Close order

#### Menu Items
- `POST /menu` - Add menu item
- `GET /menu` - Retrieve all menu items
- `GET /menu/{id}` - Retrieve specific menu item
- `PUT /menu/{id}` - Update menu item
- `DELETE /menu/{id}` - Delete menu item

#### Inventory
- `POST /inventory` - Add inventory item
- `GET /inventory` - Retrieve all inventory items
- `GET /inventory/{id}` - Retrieve specific inventory item
- `PUT /inventory/{id}` - Update inventory item
- `DELETE /inventory/{id}` - Delete inventory item

#### Reports
- `GET /reports/total-sales` - Get total sales
- `GET /reports/popular-items` - Get popular items

### Technical Implementation

- **Data Persistence**: Custom JSON file-based storage system
- **Error Handling**: Comprehensive error handling with appropriate HTTP status codes
- **Logging**: Structured logging using Go's slog package
- **Performance**: Optimized data operations with O(1) complexity for removals

### Usage

```bash
./hot-coffee [--port <N>] [--dir <S>]
./hot-coffee --help
```

Options:
- `--port N`: Specify port number
- `--dir S`: Set data directory path
- `--help`: Show help information

### Development Highlights

- Completed in 5 days with focus on code quality and architecture
- Implemented without external dependencies
- Structured for maintainability and scalability
- Robust error handling and logging system