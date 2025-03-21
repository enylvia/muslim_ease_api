# Muslim Ease API

**Muslim Ease API** is a modern and open-source RESTful API that provides a wide range of Islamic features for developers. From Al-Qur'an data and prayer times to daily duas — this API helps you build meaningful Islamic applications with ease and flexibility.

---

## ✨ Features

- 📖 Surah and Ayah information from the Qur'an
- 🕌 Accurate prayer times (coming soon)
- 🧭 Nearby mosque locator (coming soon)
- 📿 Daily duas and hadith (planned)
- 🔐 Optional JWT or API Key authentication
- 📚 Interactive Swagger-based documentation

---

## 📚 API Documentation

Swagger UI available at:

```
http://localhost:8080/swagger/index.html
```

---

## 🚀 Getting Started

### Requirements

- Go 1.20 or higher
- PostgreSQL (or your preferred DB)
- (Optional) Docker

### Installation

```bash
git clone https://github.com/yourusername/muslim-ease-api.git
cd muslim-ease-api
go mod tidy
```

### Running the API

```bash
go run main.go
```

Or with Docker:

```bash
docker-compose up --build
```

---

## 📁 Project Structure

```text
├── app/            → App initialization (DB, env)
├── handler/        → HTTP handlers (controllers)
├── service/        → Business logic
├── repository/     → DB access layer
├── model/          → Structs (request, response, DB)
├── middleware/     → Auth, rate limiter, etc.
├── docs/           → Auto-generated Swagger files
├── logs/           → API request logs
├── main.go         → Entry point
```

---

## 🛠 Tech Stack

- [Go (Golang)](https://golang.org/)
- [Gin Gonic](https://github.com/gin-gonic/gin)
- [GORM ORM](https://gorm.io/)
- [Swaggo](https://github.com/swaggo/swag) for Swagger docs
- [CORS Middleware](https://github.com/rs/cors)
- PostgreSQL, Docker

---

## 🙌 Contributing

Pull requests are welcome!  
Feel free to fork the repo and submit your improvements.

---

## 📄 License

Licensed under the [MIT License](LICENSE).

---

## 🙏 Support

If you find this project helpful, please give it a ⭐ on GitHub or share it with others.

> _“The best of people are those that bring the most benefit to the rest of mankind.”_ – Prophet Muhammad ﷺ