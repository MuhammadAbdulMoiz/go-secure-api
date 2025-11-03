# go-secure-api

A clean, modular starter project for building secure REST APIs in **Golang** with **JWT-based authentication**.

## Features

- User registration and login
- Password hashing with **bcrypt**
- JWT token generation and validation
- Protected routes using authentication middleware
- Built with **Gin-Gonic** (web framework)
- Uses **MongoDB** for data storage

---

## 📁 Project Structure

```
go-secure-api/
│   go.mod
│   go.sum
│   main.go
│   README.md
├── controllers/
│     └── userController.go
├── database/
│     └── databaseConnection.go
├── helpers/
│     ├── authHelper.go
│     └── tokenHelper.go
├── middleware/
│     └── authMiddleware.go
├── models/
│     └── user_model.go
└── routes/
		├── authRouter.go
		└── userRouter.go
```

---

## ⚙️ Tools/Technologies

Make sure you have the following installed:

| Tool    | Version             | Description           |
|---------|---------------------|-----------------------|
| Go      | 1.20+               | Programming language  |
| MongoDB | Any stable Version  | Database              |
| Git     | Latest              | Version control       |

---

## 🚀 Getting Started

1. **Clone the repository:**
	```sh
	git clone https://github.com/MuhammadAbdulMoiz/go-secure-api.git
	cd go-secure-api
	```
2. **Install dependencies:**
	```sh
	go mod tidy
	```
3. **Configure MongoDB connection:**
	- Update your MongoDB URI in `database/databaseConnection.go` if needed.
4. **Run the server:**
	```sh
	go run main.go
	```

---




