# Go + auth Quickstart for Jenkins X           
1. buka cmd
2. route ke path article-be
3. jalankan docker compose up -d
4. jalankan query create di database
CREATE TABLE article.posts (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(200) NOT NULL,
  content TEXT,
  category VARCHAR(100),
  created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  STATUS VARCHAR(100) NOT NULL DEFAULT 'Draft'
);
5. export postman collection ke postman local
6. Jalankan program secara local dengan go run cmd/http/main.go
7. Jika ada error pada go mod, jalankan go mod tidy, lalu jalankan go mod vendor, lalu jalankan program di langkah nomor 6