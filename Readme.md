# Simple Procurement System

Sistem manajemen pengadaan sederhana yang dibangun dengan arsitektur full-stack modern menggunakan Go untuk backend dan HTML/JavaScript untuk frontend.

## 🛠️ Teknologi yang Digunakan

### Backend
- Go (Golang)
- Gin Web Framework
- MySQL
- JWT Authentication

### Frontend
- HTML5
- CSS3
- Tailwind CSS
- JavaScript & jQuery

---

## 📁 Struktur Proyek

```
simple-procurement-system/
├── backend/
│   ├── cmd/
│   │   ├── migration/
│   │   └── app/
│   ├── internal/
│   │   ├── http/
│   │   │   ├── handler/
│   │   │   └── router/
│   │   ├── middleware/
│   │   ├── models/
│   │   ├── service/
│   │   └── repository/
│   ├── database/
│   │   ├── migrations/
│   │   └── seeders/
│   ├── config/
│   └── go.mod
│
├── frontend/
│   ├── assets/
│   ├── pages/
│   └── index.html
│
└── README.md
```

---

## 🚀 Instalasi & Setup

### Prasyarat
- Go 1.24
- MySQL
- Git
- Visual Studio Code
- Postman (API testing)


### Clone Repository
```bash
git clone https://github.com/DavidAfdal/simple-procurement-system.git
cd simple-procurement-system
```

### Setup Backend
```bash
cd backend
go mod download
cp .env.example .env
```

### Konfigurasi .env

File `.env` digunakan untuk menyimpan konfigurasi environment aplikasi backend.

#### MySQL Configuration
- `DATABASE_HOST`  
  Alamat server MySQL (umumnya `localhost`)

- `DATABASE_PORT`  
  Port MySQL, default `3306`

- `DATABASE_USER`  
  Username untuk mengakses MySQL

- `DATABASE_PASSWORD`  
  Password MySQL (kosongkan jika tidak ada)

- `DATABASE_DATABASE`  
  Nama database yang digunakan oleh aplikasi

#### JWT Configuration
- `JWT_SECRET_KEY`  
  Secret key yang digunakan untuk **menandatangani dan memverifikasi JWT token**.  
  Wajib diisi dan harus bersifat rahasia.

- `JWT_EXPIRES_AT`  
  Waktu kedaluwarsa token JWT (dalam jam).  
  Contoh: `1` berarti token berlaku selama 1 jam.

#### Webhook Configuration
- `WEBHOOK_URL`  
  URL tujuan webhook untuk menerima callback atau notifikasi dari backend. Url bisa didapatkan pada **https://webhook.site** .

#### Contoh `.env`
```env
ENV=development

DATABASE_HOST=localhost
DATABASE_PORT=3306
DATABASE_USER=root
DATABASE_PASSWORD=
DATABASE_DATABASE=purchase_db

JWT_SECRET_KEY=supersecret
JWT_EXPIRES_AT=1

WEBHOOK_URL=https://webhook.site/xxxxxx
```

### Run Backend
```bash
go run cmd/migration/main.go
go run cmd/app/main.go
```

### Setup Frontend
Frontend dijalankan menggunakan **extension Live Server** di Visual Studio Code.

#### Langkah-langkah:
1. Buka folder `frontend` menggunakan **Visual Studio Code**
2. Pastikan extension **Live Server** sudah terinstall
3. Klik kanan pada file `index.html`
4. Pilih **Open with Live Server**
5. Frontend akan terbuka otomatis di browser

---

## 📚 API Endpoints

### Users
- POST /users/register
- POST /users/login
- POST /users/logout

### Purchasings
- POST /purchasings
- GET /purchasings
- GET /purchasings/me

### Suppliers
- GET /suppliers
- GET /suppliers/:supplier_id
- POST /suppliers
- PUT /suppliers/:supplier_id
- DELETE /suppliers/:supplier_id

### Items
- GET /items
- GET /items/:item_id
- POST /items
- PUT /items/:item_id
- DELETE /items/:item_id

---

## 📮 Cara Import Postman Collection

Ikuti langkah-langkah berikut untuk menggunakan API melalui Postman:

### 1. Buka Postman
Pastikan aplikasi Postman sudah terinstal di komputer Anda.

### 2. Import Collection
- Klik tombol **Import**
- Pilih tab **File**
- Klik **Upload Files**
- Arahkan ke file berikut:
  ```
  backend/postman/purchasing-system-api.postman_collection.json
  ```
- Klik **Import**

### 3. Set Environment Variable
Buat environment baru di Postman dengan variabel berikut:

| Variable | Value |
|--------|-------|
| base_url | http://localhost:8080 |
| access_token | (diisi otomatis setelah login) |

### 4. Login untuk Mendapatkan Token
- Jalankan endpoint:
  ```
  POST /users/login
  ```
- Token akan otomatis tersimpan ke variabel `access_token`

### 5. Akses Endpoint Terproteksi
Untuk endpoint yang membutuhkan autentikasi, Postman akan otomatis mengirimkan header:
```
Authorization: Bearer {{access_token}}
```



