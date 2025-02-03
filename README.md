# API Tracker

The **API Tracker** is a powerful and efficient application designed to monitor and manage API usage, ensuring that clients adhere to usage limits and optimizing database storage. It implements a rate limiting mechanism to control the number of requests made by clients within a specified time frame, preventing overuse or abuse of the API. In addition, the application logs all API requests and provides a periodic cleanup process to maintain optimal performance and prevent excessive data accumulation.

### Key Features:
- **API Usage Tracking**: The application logs every API request made by clients, allowing for real-time tracking of usage statistics.
- **Rate Limiting**: It enforces limits on the number of requests a client can make within a certain period, such as per minute or hour, to prevent API abuse.
- **Log Cleanup**: The application regularly cleans up old usage logs based on customizable criteria, helping to keep database storage efficient and performance high.

### Technologies:
- **Go**: The core programming language for building the application.
- **Gin**: A fast and lightweight web framework used to create the API endpoints.
- **GORM**: An ORM library used for interacting with the MySQL database to store and retrieve API usage logs.
- **Cron**: A time-based job scheduler that handles the periodic tasks like log cleanup.
- **MySQL**: A relational database used to store client API usage logs.
- **Postman**: A tool used for testing and documenting the API endpoints, simplifying API exploration.

### How to Use:
1. Clone the repository and navigate to the project directory.
2. Install the necessary dependencies with `go mod tidy`.
3. Run the application with `go run main.go`, and it will start tracking API usage while enforcing rate limits.
4. Customize the configuration settings for rate limiting and log cleanup in the `config` file.

### API Testing:
You can test and explore the API endpoints using the Postman collection (`Api Rate tracker.postman_collection.json`) provided in the repository. The collection contains all the necessary requests and example payloads to make it easy to interact with the application.

The **API Tracker** is ideal for any application where API usage needs to be monitored and controlled, ensuring fair access while maintaining performance and efficient database management.
