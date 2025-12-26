<div align="center">
   <h1>📋 Ticket Management System API</h1>
  
  <p>
    <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
    <img src="https://img.shields.io/badge/Gorilla_Mux-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
    <img src="https://img.shields.io/badge/MySQL-4479A1?style=for-the-badge&logo=mysql&logoColor=white" />
    <img src="https://img.shields.io/badge/Bcrypt-003A8F?style=for-the-badge&logo=letsencrypt&logoColor=white" />
    <img src="https://img.shields.io/badge/UUID-FF6C37?style=for-the-badge&logo=uuid&logoColor=white" />
  </p>

  <p>
    <strong>RESTful API untuk sistem manajemen tiket acara</strong><br />
  </p>

  <p>
    <img src="https://img.shields.io/badge/Status-Active-success" />
    <img src="https://img.shields.io/badge/Version-1.0.0-blue" />
  </p>
</div>

---

## 🌟 Gambaran Umum

> Backend service terstruktur untuk mengelola sistem pemesanan tiket acara, dengan fokus pada **concurrency control**, **transaction management**, dan **clean architecture**

Ticket Management System API adalah backend service yang menyediakan fitur pengelolaan sistem tiket untuk berbagai acara, mulai dari registrasi pengguna, manajemen event, booking tiket, hingga validasi penggunaan tiket.

Aplikasi ini dirancang dengan pendekatan **RESTful API**, menerapkan **clean architecture Go**, serta memperhatikan aspek **keamanan data, transaction safety, dan race condition handling**.

---

## ✨ Fitur Utama

### 🚀 Kenapa Project Ini? 
- Cocok untuk **tugas kuliah**, **portofolio backend**, dan **latihan sistem skala menengah**
- Menerapkan **database transaction & row-level locking** untuk mencegah race condition
- Struktur kode mengikuti **clean architecture** dengan separation of concerns
- **Repository pattern** untuk abstraksi data layer
- **Validation layer** terpisah untuk input validation

### 🔹 Fungsionalitas Inti
- 👥 **User Management** (Customer & Organizer role)
- 🎫 **Event Management** (CRUD dengan role-based access)
- 📝 **Booking System** (dengan automatic seat management)
- 🎟️ **Ticket Generation** (unique ticket code per booking)
- ✅ **Ticket Validation** (mark as used)
- 🔐 **Password Hashing** dengan bcrypt

### 🔹 Sorotan Teknis
- 🎯 **Database Transaction** dengan row-level locking (`FOR UPDATE`)
- 🔒 **Concurrency Control** untuk mencegah overbooking
- ✅ **Input Validation** di layer validation terpisah
- 🏗️ **Clean Architecture** - Handler → Service → Repository → Database
- 🛡️ **Password Security** dengan bcrypt hashing
- 🔄 **Automatic Status Management** (available/unavailable based on seats)
- 🚦 **Consistent Error Handling** dengan standard response format

---

## 📋 Daftar Isi
1. [Prasyarat](#-1-prasyarat)
2. [Instalasi](#-2-instalasi)
3. [Setup Database](#️-3-setup-database)
4. [Konfigurasi Environment](#️-4-konfigurasi-environment)
5. [Menjalankan Aplikasi](#-5-menjalankan-aplikasi)
6. [Struktur Proyek](#-6-struktur-proyek)
7. [API Endpoints](#-7-api-endpoints)
8. [Testing dengan Postman](#-8-testing-dengan-postman)
9. [Arsitektur & Keamanan](#️-9-arsitektur--keamanan)
10. [Fitur Keamanan](#️-10-fitur-keamanan)
11. [Business Logic](#-11-business-logic)
12. [Keterbatasan & Rencana Pengembangan](#️-12-keterbatasan--rencana-pengembangan)
13. [Troubleshooting](#️-13-troubleshooting)
14. [Penutup](#-14-penutup)

---

## 🔧 1. Prasyarat

Pastikan environment pengembangan telah memenuhi kebutuhan berikut:
- **Go** >= 1.23.5
- **MySQL** >= 8.0
- **Git** (untuk clone repository)
- **Postman** (untuk testing API)

---

## 📦 2. Instalasi

Clone repository dan install seluruh dependency:

```bash
git clone https://github.com/WahyuPratama222/Ticket-Api-Golang.git
cd Ticket-Api-Golang
go mod download
```

---

## 🗄️ 3. Setup Database

### 🔹 Buat Database

```sql
CREATE DATABASE ticket_system;
USE ticket_system;
```

### 🔹 Jalankan Migration

```bash
# Run migration
go run cmd/migrations/migrate_main.go
```

Migration akan membuat 4 tabel:
1. **user** - Data pengguna (customer & organizer)
2. **event** - Data acara
3. **booking** - Data pemesanan tiket
4. **ticket** - Data tiket individual

### 🔹 Jalankan Seeder (Optional)

```bash
# Run seeder untuk data dummy
go run cmd/seeders/seeder_main.go
```

Seeder akan membuat:
- 3 user (1 organizer, 2 customer)
- 2 event dummy

---

## ⚙️ 4. Konfigurasi Environment

Buat file `.env` di root project dengan template berikut:

```env
# Server Configuration
PORT=8080

# Database Configuration
DB_USER=root
DB_PASSWORD=your_mysql_password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=ticket_system
```

> ⚠️ **Penting**: Ganti `DB_PASSWORD` dengan password MySQL Anda!

### Penjelasan Environment Variables

| Variable | Deskripsi | Default | Required |
|----------|-----------|---------|----------|
| `PORT` | Port server API | 8080 | No |
| `DB_USER` | Username MySQL | root | Yes |
| `DB_PASSWORD` | Password MySQL | - | Yes |
| `DB_HOST` | Host MySQL | 127.0.0.1 | Yes |
| `DB_PORT` | Port MySQL | 3306 | Yes |
| `DB_NAME` | Nama database | ticket_system | Yes |

---

## 🚀 5. Menjalankan Aplikasi

### Menjalankan Server

```bash
go run ./cmd/api
```

Server akan berjalan di `http://localhost:8080`

### Default User untuk Testing (dari seeder)

| Name | Email | Password | Role |
|------|-------|----------|------|
| Dedi Mulyados | mulmulmul@mail.com | mullllya123 | organizer |
| Pak Jkw | jwkwkw@mail.com | jwkwkww123 | customer |
| Bahlilul | bahlilu@mail.com | bahabahha123 | customer |

---

## 📁 6. Struktur Proyek

```
Ticket-Api-Golang/
├── cmd/
│   ├── api/
│   │   ├── main.go                # Entry point aplikasi
│   │   └── routes.go              # Route definitions
│   ├── migrations/
│   │   └── migrate_main.go        # Migration runner
│   └── seeders/
│       └── seeder_main.go         # Seeder runner
├── handlers/                       # HTTP handlers (presentation layer)
│   ├── user_handler.go
│   ├── event_handler.go
│   ├── booking_handler.go
│   └── ticket_handler.go
├── services/                       # Business logic layer
│   ├── user_service.go
│   ├── event_service.go
│   ├── booking_service.go
│   └── ticket_service.go
├── repositories/                   # Data access layer
│   ├── user_repository.go
│   ├── event_repository.go
│   ├── booking_repository.go
│   └── ticket_repository.go
├── validations/                    # Input validation layer
│   ├── user_validator.go
│   ├── event_validator.go
│   ├── booking_validator.go
│   └── ticket_validator.go
├── models/                         # Data models
│   ├── user.go
│   ├── event.go
│   ├── booking.go
│   └── ticket.go
├── migrations/                     # Database migrations
│   ├── 01_user.go
│   ├── 02_event.go
│   ├── 03_booking.go
│   ├── 04_ticket.go
│   └── migrate.go
├── seeders/                        # Database seeders
│   ├── 01_user.go
│   ├── 02_event.go
│   └── seeder.go
├── pkg/
│   └── db/
│       └── connection.go          # Database connection
├── utils/
│   └── response.go                # Standard response helpers
├── postman/
│   └── Ticket System API Test.postman_collection.json
├── .env                           # Environment variables
├── .gitignore
├── go.mod                         # Go module dependencies
├── go.sum
└── README.md
```

### Penjelasan Layer Arsitektur

1. **Handlers** → Menangani HTTP request/response
2. **Services** → Business logic dan transaction management
3. **Repositories** → Database operations dan queries
4. **Validations** → Input validation logic
5. **Models** → Data structure definitions
6. **Utils** → Helper functions

---

## 🔌 7. API Endpoints

> Seluruh endpoint mengikuti prinsip **RESTful API** dan menggunakan format response JSON.

### 👥 User Management

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/users/register` | Register user baru |
| GET | `/users` | Ambil semua user |
| GET | `/users/{id}` | Ambil user berdasarkan ID |
| PUT | `/users/{id}` | Update user |
| DELETE | `/users/{id}` | Hapus user |

**Contoh Request Body POST /users/register:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123",
  "role": "customer"
}
```

**Response Success:**
```json
{
  "success": true,
  "message": "user registered successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "role": "customer",
    "created_at": "2025-01-15T10:30:00Z",
    "updated_at": "2025-01-15T10:30:00Z"
  }
}
```

---

### 🎫 Event Management

| Method | Endpoint | Deskripsi | Notes |
|--------|----------|-----------|-------|
| POST | `/events` | Buat event baru | Hanya organizer |
| GET | `/events` | Ambil semua event | - |
| GET | `/events/{id}` | Ambil event berdasarkan ID | - |
| PUT | `/events/{id}` | Update event | Hanya organizer |
| DELETE | `/events/{id}` | Hapus event | Hanya organizer |

**Contoh Request Body POST /events:**
```json
{
  "organizer_id": 1,
  "title": "Konser Rock 2025",
  "location": "Jakarta Stadium",
  "capacity": 100,
  "price": 500000,
  "date": "2025-12-31T19:00:00Z"
}
```

**Business Rules:**
- Hanya user dengan role `organizer` yang bisa create/update/delete event
- `available_seat` otomatis di-set sama dengan `capacity` saat create
- `status` otomatis menjadi `unavailable` jika `available_seat` = 0
- Event tidak bisa dihapus jika sudah ada booking dengan status `success`

---

### 📝 Booking Management

| Method | Endpoint | Deskripsi | Notes |
|--------|----------|-----------|-------|
| POST | `/bookings` | Buat booking baru | Auto-generate tickets |
| GET | `/bookings` | Ambil semua booking | - |
| GET | `/bookings/{id}` | Ambil detail booking + tickets | - |

**Contoh Request Body POST /bookings:**
```json
{
  "customer_id": 2,
  "event_id": 1,
  "quantity": 3,
  "holder_names": ["Alice", "Bob", "Charlie"]
}
```

**Proses di Backend:**
1. **Lock row event** dengan `FOR UPDATE` (mencegah race condition)
2. Validasi event status = `available`
3. Validasi `available_seat` >= `quantity`
4. Kurangi `available_seat` sebesar `quantity`
5. Update status event jadi `unavailable` jika seat habis
6. Buat record booking dengan status `pending`
7. Generate unique ticket code untuk setiap quantity
8. Buat record ticket untuk setiap holder
9. Update booking status jadi `success`
10. Commit transaction

**holder_names (optional):**
- Jika tidak dikirim atau kosong, akan auto-generate: "Ticket 1", "Ticket 2", dst
- Jika dikirim, jumlah harus sama dengan `quantity`

---

### 🎟️ Ticket Management

| Method | Endpoint | Deskripsi | Notes |
|--------|----------|-----------|-------|
| GET | `/tickets` | Ambil semua ticket | - |
| GET | `/tickets/{id}` | Ambil ticket berdasarkan ID | - |
| PUT | `/tickets/{id}/use` | Mark ticket sebagai used | One-time use |

**Business Rules:**
- Ticket hanya bisa di-use satu kali
- Status berubah dari `unused` → `used`
- Ticket yang sudah `used` tidak bisa di-use lagi

---

## 🧪 8. Testing dengan Postman

### 📥 Download & Import Collection

1. **Download Postman**
   - Kunjungi [https://www.postman.com/downloads/](https://www.postman.com/downloads/)
   - Download sesuai OS Anda (Windows/Mac/Linux)
   - Install aplikasi Postman

2. **Import Collection**
   - Buka Postman
   - Klik tombol **"Import"** di pojok kiri atas
   - Pilih file `postman/Ticket System API Test.postman_collection.json`
   - Klik **"Import"**

3. **Setup Environment (Optional)**
   - Klik icon gear (⚙️) di pojok kanan atas
   - Klik **"Add"** untuk membuat environment baru
   - Nama: `Ticket API Local`
   - Tambahkan variable:
     - `base_url` = `http://localhost:8080`
   - Save environment
   - Pilih environment dari dropdown di pojok kanan atas

### 🧪 Testing Flow

#### 1. User Management
```
1. POST /users/register - Buat user baru (customer & organizer)
2. GET /users - Lihat semua user
3. GET /users/{id} - Lihat detail user
4. PUT /users/{id} - Update user
5. DELETE /users/{id} - Hapus user (jika tidak ada booking/event)
```

#### 2. Event Management
```
1. POST /events - Buat event (gunakan organizer_id dari user organizer)
2. GET /events - Lihat semua event
3. GET /events/{id} - Lihat detail event
4. PUT /events/{id} - Update event (capacity, price, status, dll)
5. DELETE /events/{id} - Hapus event (jika belum ada booking)
```

#### 3. Booking Flow
```
1. GET /events - Pilih event yang tersedia
2. POST /bookings - Buat booking (customer_id + event_id + quantity)
3. GET /bookings/{id} - Lihat detail booking + tickets
4. GET /tickets - Lihat semua ticket yang tergenerate
```

#### 4. Ticket Validation
```
1. GET /tickets/{id} - Cek status ticket (unused/used)
2. PUT /tickets/{id}/use - Mark ticket sebagai used
3. GET /tickets/{id} - Verify status berubah jadi "used"
```

### 📝 Tips Testing

- **Test Concurrency:** Buat 2+ booking simultan ke event yang sama untuk test race condition handling
- **Test Capacity:** Coba booking dengan quantity > available_seat (should fail)
- **Test Role:** Coba buat event dengan customer_id (should fail - only organizer)
- **Test Validation:** Coba kirim data invalid (email salah, password < 8 karakter, dll)

### 🎯 Expected Responses

**Success Response:**
```json
{
  "success": true,
  "message": "operation success message",
  "data": { ... }
}
```

**Error Response:**
```json
{
  "success": false,
  "error": "error message description"
}
```

---

## 🏗️ 9. Arsitektur & Keamanan

### Flow Arsitektur

```
Client Request
      ↓
   Router (Gorilla Mux)
      ↓
   Handler (HTTP handling)
      ↓
   Service (Business Logic)
      ↓
   Validator (Input validation)
      ↓
   Repository (Database operations)
      ↓
   MySQL Database
      ↓
   Response ke Client
```

**Catatan penting:**
- Handler hanya handle HTTP request/response
- Service berisi semua business logic & transaction management
- Validator memastikan data valid sebelum masuk ke repository
- Repository hanya focus pada database operations

---

### Prinsip Desain

1. **Clean Architecture**
   - Handler: HTTP layer
   - Service: Business logic layer
   - Repository: Data access layer
   - Validation: Input validation layer
   - Model: Data structure

2. **Separation of Concerns**
   - Setiap layer punya tanggung jawab spesifik
   - Tidak ada business logic di handler
   - Tidak ada database query di service (kecuali transaction management)

3. **Database Transaction Safety**
   - Transaction untuk operasi kritikal (booking)
   - Row-level locking untuk race condition handling
   - Rollback mechanism untuk data consistency

4. **Error Handling**
   - Consistent error response format
   - Specific error messages untuk debugging
   - HTTP status codes yang sesuai

---

## 🛡️ 10. Fitur Keamanan

### 1. Password Hashing
- Menggunakan **bcrypt** dengan default cost (10)
- Password tidak pernah disimpan dalam bentuk plain text
- Hashing dilakukan di service layer sebelum save ke database

```go
hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
```

### 2. SQL Injection Prevention
- Menggunakan **prepared statements** untuk semua queries
- Parameter binding untuk user input
- Tidak ada string concatenation untuk SQL queries

```go
query := `INSERT INTO user (name, email, password, role) VALUES (?, ?, ?, ?)`
db.DB.Exec(query, user.Name, user.Email, user.Password, user.Role)
```

### 3. Input Validation
- Validasi di layer terpisah (validations/)
- Email format validation dengan regex
- Password minimum 8 karakter
- Enum validation untuk field tertentu (role, status)

### 4. Race Condition Prevention
- Row-level locking dengan `FOR UPDATE`
- Transaction isolation untuk booking process
- Atomic operations untuk seat updates

```go
query := `SELECT id_event, available_seat, price, status FROM event WHERE id_event=? FOR UPDATE`
```

### 5. Business Rule Enforcement
- Role-based access control (organizer untuk event management)
- Cascade delete prevention (user dengan booking/event tidak bisa dihapus)
- Capacity validation (tidak bisa booking lebih dari available seats)

---

## 💼 11. Business Logic

### Manajemen Seat Event

#### Saat Booking Dibuat
```go
// 1. Lock row event (prevent race condition)
event, err := repo.GetEventWithLock(tx, eventID)

// 2. Validate seat availability
if event.AvailableSeat < booking.Quantity {
    return errors.New("not enough seats available")
}

// 3. Update available seats
newAvailableSeat := event.AvailableSeat - booking.Quantity
newStatus := "available"
if newAvailableSeat == 0 {
    newStatus = "unavailable"
}

repo.UpdateEventSeats(tx, eventID, newAvailableSeat, newStatus)
```

### Ticket Generation

Setiap booking akan generate sejumlah ticket sesuai `quantity`:

```go
for i := 0; i < booking.Quantity; i++ {
    ticketCode := uuid.New().String()[:8]  // Unique 8-char code
    
    holderName := "Ticket 1"  // Default
    if i < len(booking.HolderNames) {
        holderName = booking.HolderNames[i]  // Custom holder name
    }
    
    ticket := models.Ticket{
        BookingID:  booking.ID,
        HolderName: holderName,
        TicketCode: ticketCode,
        Status:     "unused",
    }
    repo.CreateTicket(tx, &ticket)
}
```

### Status Management

**User:**
- `role`: "customer" | "organizer"
- Customer: Bisa booking tiket
- Organizer: Bisa manage event

**Event:**
- `status`: "available" | "unavailable"
- Auto-update ke "unavailable" jika `available_seat` = 0

**Booking:**
- `status`: "pending" | "success" | "failed"
- Default: "pending"
- Success: Jika semua ticket berhasil digenerate
- Failed: Jika ada error saat generate ticket (akan rollback)

**Ticket:**
- `status`: "unused" | "used"
- Default: "unused"
- One-time use: Setelah "used", tidak bisa diubah lagi

---

## ⚠️ 12. Keterbatasan & Rencana Pengembangan

### ❌ Keterbatasan Saat Ini

1. **Tidak Ada Authentication/Authorization**  
   Belum ada JWT atau session management. Role-based access control hanya di level validation, tidak ada token authentication.

2. **Tidak Ada Payment Gateway Integration**  
   Booking langsung success tanpa proses pembayaran.

3. **Tidak Ada Email Notification**  
   Tidak ada notifikasi email untuk booking confirmation atau ticket delivery.

4. **Tidak Ada Pagination**  
   Semua endpoint GET mengembalikan seluruh data tanpa pagination.

5. **Tidak Ada Rate Limiting**  
   API belum dilindungi dari brute force atau spam requests.

6. **Tidak Ada Soft Delete**  
   Delete operation adalah hard delete, tidak bisa recover.

7. **Tidak Ada Audit Trail**  
   Tidak ada logging untuk track who did what when.

8. **Tidak Ada File Upload**  
   Event tidak bisa upload poster/image.

9. **Booking Tidak Bisa Cancel**  
   Setelah booking success, tidak ada mekanisme cancel atau refund.

### 🚧 Rencana Pengembangan Selanjutnya

#### Short Term (1-3 bulan)
- ✅ Implementasi **JWT Authentication**
  - Login endpoint dengan token generation
  - Protected routes dengan middleware
- ✅ Tambah **Pagination & Filtering** di GET endpoints
- ✅ Implementasi **Rate Limiting** dengan middleware
- ✅ Tambah **Soft Delete** mechanism
- ✅ Setup **Logging** dengan structured logging

#### Mid Term (3-6 bulan)
- ✅ Implementasi **Payment Gateway** (Midtrans/Xendit)
- ✅ Tambah **Email Notification** service
- ✅ Implement **Booking Cancellation** dengan refund flow
- ✅ Tambah **File Upload** untuk event images
- ✅ Setup **Unit Testing** dengan testify
- ✅ Implement **API Versioning** (/api/v1, /api/v2)

#### Long Term (6-12 bulan)
- ✅ Setup **Docker** containerization
- ✅ Implement **Caching Layer** dengan Redis
- ✅ Setup **CI/CD Pipeline** dengan GitHub Actions
- ✅ Implement **Microservices Architecture**
- ✅ Tambah **WebSocket** untuk real-time notifications
- ✅ Setup **Monitoring & Alerting** (Prometheus, Grafana)
- ✅ Implement **GraphQL** sebagai alternatif REST

---

## 🛠️ 13. Troubleshooting

### Problem: Error "Access denied for user 'root'@'localhost'"

**Solution:**
```bash
# Login ke MySQL
mysql -u root -p

# Update password
ALTER USER 'root'@'localhost' IDENTIFIED BY 'your_new_password';
FLUSH PRIVILEGES;

# Update .env file dengan password baru
```

---

### Problem: Error "connect: connection refused"

**Solution:**
```bash
# Check apakah MySQL service berjalan
# Windows
net start MySQL80

# macOS/Linux
sudo systemctl start mysql
# atau
sudo service mysql start
```

---

### Problem: Error "database not found"

**Solution:**
```sql
-- Buat database manual
CREATE DATABASE ticket_system;
USE ticket_system;

-- Jalankan migration
go run cmd/migrations/migrate_main.go
```

---

### Problem: Port 8080 already in use

**Solution:**
```bash
# Windows
netstat -ano | findstr :8080
taskkill /PID <PID> /F

# macOS/Linux
lsof -ti:8080 | xargs kill -9

# Atau ganti port di .env
PORT=8081
```

---

### Problem: Migration Error "table already exists"

**Solution:**
```sql
-- Drop semua table dan run migration lagi
DROP TABLE IF EXISTS ticket;
DROP TABLE IF EXISTS booking;
DROP TABLE IF EXISTS event;
DROP TABLE IF EXISTS user;

-- Run migration kembali
go run cmd/migrations/migrate_main.go
```

---

### Problem: Seeder Error "duplicate entry"

**Solution:**
Seeder menggunakan `ON DUPLICATE KEY UPDATE` jadi aman dijalankan multiple times. Jika tetap error, truncate table dulu:

```sql
TRUNCATE TABLE ticket;
TRUNCATE TABLE booking;
TRUNCATE TABLE event;
TRUNCATE TABLE user;

-- Run seeder kembali
go run cmd/seeders/seeder_main.go
```

---

## 📧 14. Penutup

Dokumentasi API ini disusun untuk memberikan gambaran yang jelas mengenai struktur endpoint, flow business logic, serta implementation details pada tahap pengembangan saat ini.

Project ini fokus pada penerapan **clean architecture**, **transaction safety**, dan **race condition handling** untuk sistem booking yang robust.

Dokumentasi ini diharapkan dapat menjadi referensi teknis yang akurat terhadap implementasi sistem saat ini, sekaligus menjadi dasar pengembangan lanjutan pada iterasi berikutnya.

— Wahyu Pratama

---
