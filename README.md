# GOLANG BY KOMGRIP TEAM
## _e-permit-api-service_

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https)

## Migration
### คำสั่งในการติดตั้ง migrate
```sh
brew install golang-migrate
```
### สร้าง/ลบ table
```sh
migrate create -ext sql -dir database/migrations -seq create_${table}_table
```
### file .up.sql
```sql
CREATE TABLE IF NOT EXISTS users(
   user_id serial PRIMARY KEY,
   username VARCHAR (50) UNIQUE NOT NULL,
   password VARCHAR (50) NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL
);
```
### file .down.sql
```sql
DROP TABLE IF EXISTS users;
```

### เพิ่ม/ลบ column
```sh
migrate create -ext sql -dir database/migrations -seq create_${column}_${table}_table
```
### file .up.sql
```sql
ALTER TABLE users ADD COLUMN `mood` int;
```

### file .down.sql
```sql
Alter table users drop column `mood`
```
### run migrate
```sh
migrate -database 'mysql://root:1qazxsw2@tcp(localhost:3306)/bma-e-permit-api-service' -path databases/migrations up
```
```sh
migrate -database 'mysql://root:1qazxsw2@tcp(localhost:3306)/bma-e-permit-api-service' -path databases/migrations down
```
```sh
migrate -database 'mysql://root:1qazxsw2@tcp(localhost:3306)/bma-e-permit-api-service' -path databases/migrations force {version}
```

<br>
<br>
<br>

# ✨ คู่มือการใช้ - Golang Seeder CLI - Help you seed data easily.
## 🔧 Installation

```sh
go install github.com/komgrip/seeder@latest
```

It is always installed in `~/go/bin`

<br/>

## 📝 Get Started
### 🏗️ **Create seed file**

```sh
seeder create -dir example/seeds/masterdata example_table1
```

```go
//คำสั่งในการสร้าง seed file
seeder create 

//กำหนด directory path ที่ต้องการ ในกรณีที่มี directory path อยู่แล้ว
//จะสร้าง seed file ในนั้น แต่ถ้ายัง command จะสร้าง directory path และ seed file
-dir example/seeds/masterdata

//ชื่อตารางที่เราต้องการจะทำการ seed data
example_table1
```
seed file จะถูก gen ชื่อเป็น `seed_{ชื่อตารางที่เราใส่เข้ามา}_table.sql`

ในกรณีตามตัวอย่างข้างต้นชื่อ seed file จะเป็น `seed_example_table1_table.sql`

<br/>

### 🏗️ **Create multiple seed files**
```sh
seeder create -dir example/seeds/masterdata example_table1 example_table2
```

<br/>

### 📥 **Seed data all files**
```sh
seeder seed -database 'postgres://username:password@localhost:5432/example-db?sslmode=disable' -path databases/seeds/masterdata
```
```go
//คำสั่งใน seed data
seeder seed

//database ที่ต้องการสำหรับ insert ข้อมูล
-database 'postgres://username:password@localhost:5432/example-db?sslmode=disable'
```
<br/>

### 📥 **Seed data specific file**
```sh
seeder seed -database 'postgres://username:password@localhost:5432/example-db?sslmode=disable' -path databases/seeds/masterdata seed_example_table1_table.sql
```
```go
//file ที่ต้องการจะ seed
seed_example_table1_table.sql
```
