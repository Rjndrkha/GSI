# Personal Finance Management API - PT GSI Technical Test

## 🛠️ Tech Stack

- **Language**: Go
- **Framework**: Fiber v2
- **ORM**: GORM
- **Database**: PostgreSQL
- **Libraries**: JWT, Excelize (Excel Generator), Bcrypt, Google UUID, Godotenv.

---

## 📌 API Endpoints

### Authentication

- `POST /api/auth/login` - Mendapatkan JWT Token
- `GET /api/auth/profile` - Mengambil profil user (Auth required)

### Pockets

- `POST /api/pockets` - Membuat pocket baru
- `GET /api/pockets` - List semua pocket user
- `GET /api/pockets/total-balance` - Akumulasi saldo semua pocket

### Transactions

- `POST /api/incomes` - Menambah saldo pocket
- `POST /api/expenses` - Mengurangi saldo pocket (dengan validasi saldo cukup)

### Reports

- `POST /api/pockets/:id/create-report` - Memicu background job pembuatan Excel
- `GET /reports/:id` - Download/Stream file Excel (Public)

## 🚀 Quick Start

1. **Clone Repository**

   ```bash
   git clone https://github.com/Rjndrkha/GSI.git
   cd GSI
   ```

2. **Install Dependencies & Run**
   ```bash
   go mod tidy
   go run cmd/api/main.go
   ```

## 📖 Deskripsi Fitur

API untuk manajemen keuangan pribadi yang dibangun menggunakan **Golang** dengan framework **Fiber**. Proyek ini menerapkan **Clean Architecture** untuk memastikan kode mudah dipelihara dan diuji.

- **Authentication**: Login menggunakan JWT
- **Pocket Management**: Membuat dan melihat daftar kantong
- **Transaction Logic**: Pencatatan Income & Expense dengan Database Transaction (ACID)
- **Reporting**: Generate laporan Excel di Background Process
- **File Streaming**: Endpoint publik untuk mengunduh laporan
