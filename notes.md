### Userful notes

---
Criterias:
    - If an error occurs during startup (e.g., invalid command-line arguments, failure to bind to a port), the program must exit with a non-zero status code and display a clear, understandable error message. During normal operation, the server must handle errors gracefully, returning appropriate HTTP status codes to the client without crashing.

    - Store data locally in JSON files. Data should be stored in separate JSON files for each entity, such as orders.json, menu_items.json and inventory.json located in a designated data/ directory
