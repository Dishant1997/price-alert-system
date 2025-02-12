# Price Alert System

The **Price Alert System** is a decentralized application (dApp) designed to monitor asset prices and notify users when their specified price thresholds are crossed. It uses **Chainlink** for fetching asset prices, **MongoDB** for data storage, and is built using **Go**.

---

## Features

- Users can set price alerts for specific assets.
- Alerts are triggered when the price crosses the user-defined threshold (above or below).
- Notifications are sent to users based on their alert preferences.
- Backend implemented with Go, Chainlink, and MongoDB.

---


---

## Prerequisites

1. **Go**: Ensure Go is installed on your system. [Download Go](https://go.dev/doc/install)
2. **MongoDB**: Install and set up MongoDB. Use [MongoDB Compass](https://www.mongodb.com/products/compass) for GUI management.
3. **Environment Variables**: Create a `.env` file in the root directory with the following:
    ```env
    MONGODB_URI=mongodb://localhost:27017
    MONGODB_DATABASE=price_alert_system
    MONGODB_COLLECTION=alerts
    ```

---

## Setup and Installation

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/yourusername/price-alert-system.git
   cd price-alert-system
    ```
2. **Install Dependencies:
    ```bash
    go mod tidy
    ```

3. **Run MongoDB Server: Ensure MongoDB is running locally or on a configured server.

4. **Run the Backend:
    ```bash
    go run cmd/main.go
    ```

5. **Test API Endpoints: Use tools like Postman or cURL to test the API endpoints provided in handlers/price_handler.go.

