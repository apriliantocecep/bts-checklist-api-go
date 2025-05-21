# 📝 Check List API

A simple and secure RESTful API for managing checklists and tasks, built using the Go programming language. This project demonstrates how to implement JWT-based authentication, relational data modeling with GORM, and a clean routing structure using the Fiber web framework.

## 🚀 Tech Stack

- **Go** (Golang)
- **Fiber** – Express-style web framework for Go
- **GORM** – ORM for Go
- **PostgreSQL** – Relational database
- **JWT** – Authentication and authorization

## 🔐 Features

- User registration and login with JWT
- Authentication middleware
- CRUD operations for:
    - ✅ Checklists
    - 📝 Items in checklists
- Toggle task completion status
- Clean and modular project structure

## 📦 API Endpoints

### Auth
- `POST /login` – User login
- `POST /register` – User registration

### Checklist (Protected)
- `GET /checklist` – Get all checklists
- `POST /checklist` – Create a new checklist
- `DELETE /checklist/:checklistId` – Delete checklist

### Item (Protected)
- `GET /checklist/:checklistId/item` – Get items in a checklist
- `POST /checklist/:checklistId/item` – Add item to a checklist
- `GET /checklist/:checklistId/item/:checklistItemId` – Get item detail
- `PUT /checklist/:checklistId/item/:checklistItemId` – Toggle item status
- `DELETE /checklist/:checklistId/item/:checklistItemId` – Delete item
- `PUT /checklist/:checklistId/item/rename/:checklistItemId` – Rename item

## ⚙️ Getting Started

1. Clone the repo
2. Setup `.env` file with your DB and JWT settings
3. Run migrations (auto or manually)
4. Start the server

```bash
go run main.go
```

## 📬 API Testing with Postman
You can test all available endpoints using the included Postman collection:

### 📥 Download Postman Collection:
BTS.id Checklist API.postman_collection.json
