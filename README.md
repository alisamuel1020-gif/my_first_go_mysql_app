# My First Go & MySQL App 🚀

This is my first backend project learning how to connect a **Go** application to a **MySQL** database. 

The project demonstrates how to connect to a database, safely execute parameterized queries (to prevent SQL injection), and load sensitive credentials using environment variables.

## 🛠️ Prerequisites

Before you run this project, make sure you have the following installed:
* [Go](https://go.dev/doc/install) (version 1.20 or higher)
* [MySQL](https://dev.mysql.com/doc/refman/8.0/en/installing.html) (running locally on port `3306`)

## ⚙️ Setup & Installation

**1. Clone the repository:**
```bash
git clone https://github.com/alisamuel1020-gif/my_first_go_mysql_app
cd my_first_go_mysql_app
```

**2. Install dependencies:**
```bash
go get github.com/go-sql-driver/mysql
go get github.com/joho/godotenv
```

**3. Set up the Database:**
Log into your MySQL server and create the database and table

**4. Create the Environment File:**
Create a file named .env in the root directory and add your specific database credentials.
