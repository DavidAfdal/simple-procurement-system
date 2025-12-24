# Simple Procurement System

Sistem manajemen pengadaan sederhana yang dibangun dengan arsitektur full-stack modern menggunakan Go untuk backend dan HTML/JavaScript untuk frontend.

## 🛠️ Teknologi yang Digunakan

### Backend
- **Go (Golang)** - Bahasa pemrograman utama untuk REST API
- Framework web Go (Gin)
- Database (MySQL)
- JWT untuk autentikasi

### Frontend
- **HTML5** - Struktur halaman web
- **CSS3** - Styling dan layout
- **Tailwind** - Styling dan layout
- **JavaScript & JQuery** - Interaktivitas dan komunikasi dengan API

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
│   ├── pkg/                    
│   └── go.mod
│
├── frontend/                   
│   ├── assets/                 
│   ├── pages/                  
│   └── index.html              
│
└── README.md
```


## 🚀 Instalasi & Setup

### Prasyarat

Pastikan Anda telah menginstal:
- Go 1.24 
- Database (MySQL)
- Web browser
- Git

### Langkah Instalasi

1. **Clone repository**
   ```bash
   git clone https://github.com/DavidAfdal/simple-procurement-system.git
   cd simple-procurement-system
   ```

2. **Setup Backend**
   ```bash
   cd backend
   
   # Install dependencies
   go mod download
   
   # Copy dan konfigurasi environment variables
   cp .env.example .env
   ```

3. **Konfigurasi .env file**
    Buat file `.env` di folder backend dengan konfigurasi berikut:

    ```env
    ENV=development

    DATABASE_HOST=localhost
    DATABASE_PORT=3306
    DATABASE_USER=root
    DATABASE_PASSWORD=
    DATABASE_DATABASE=purchase_db

    JWT_SECRET_KEY=supersecret
    JWT_EXPIRES_AT=1

    WEBHOOK_URL=https://webhook.site/f0129d12-4d5b-4e35-8ddc-3dcaddf9405d0
    ```
5. **Run Backend**
    ```bash
    # Jalankan migrations
    go run migration/main.go
    
    # Jalankan server
    go run main.go
    ```

4. **Setup Frontend**
   ```bash
   cd ../frontend
   
   # Jika menggunakan live server, jalankan:
   # npx live-server
   
   # Atau buka langsung index.html di browser
   ```

4. **Setup Frontend**
   ```bash
   cd ../frontend
   
   # Jika menggunakan live server, jalankan:
   # npx live-server
   
   # Atau buka langsung index.html di browser
   ```


5. **Akses Aplikasi**
   - Backend API: `http://localhost:8080`

## 🔧 Konfigurasi

Buat file `.env` di folder backend dengan konfigurasi berikut:

```env
# Server Configuration
PORT=8080
ENV=development

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=procurement_db

# JWT Configuration
JWT_SECRET=your_secret_key_here
JWT_EXPIRATION=24h

# CORS Configuration
ALLOWED_ORIGINS=http://localhost:3000
```

## 📚 API Endpoints

### Authentication
- `POST /api/users/login` - User login
- `POST /api/users/logout` - User logout
- `POST /api/users/register` - Register user baru

### Purchasing
- `GET /api/purchasings` - Get semua purchase requests
- `POST /api/purchasings` - Create purchase request baru
- `PUT /api/purchasings/me` - Update purchase request

### Vendors
- `GET /api/vendors` - Get semua vendors
- `GET /api/vendors/:id` - Get vendor by ID
- `POST /api/vendors` - Create vendor baru
- `PUT /api/vendors/:id` - Update vendor
- `DELETE /api/vendors/:id` - Delete vendor

### Approvals
- `GET /api/approvals` - Get pending approvals
- `POST /api/approvals/:id/approve` - Approve request
- `POST /api/approvals/:id/reject` - Reject request

## 👤 Default User

Setelah menjalankan migrations, Anda dapat login dengan:

```
Email: admin@procurement.com
Password: admin123
```

