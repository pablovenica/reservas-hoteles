
# 🏨 Sistema de Reservas de Hoteles

## 🏗 Arquitectura del Proyecto

El sistema está compuesto por tres servicios independientes, cada uno con responsabilidades específicas y base de datos dedicada:

| Servicio | Responsabilidad | Base de Datos |
| :--- | :--- | :--- |
| **User API** | Gestión de usuarios y autenticación | MySQL (`mysql_users`) | 
| **Hotels API** | Catálogo, búsqueda y gestión de hoteles | MongoDB (`mongo_hoteles`) |
| **Reservation API** | Lógica transaccional de reservas y disponibilidad | MySQL (`mysql_reservation`) |

---

## 🛠 Stack Tecnológico

### Infraestructura
- **Docker & Docker Compose** - Orquestación de contenedores
- **MySQL** - Base de datos relacional para transacciones y usuarios
- **MongoDB** - Base de datos NoSQL para el catálogo de hoteles

## 📋 Pre-requisitos

Antes de comenzar, asegúrate de tener instalado:

- [Docker Desktop](https://www.docker.com/products/docker-desktop) (o Docker Engine + Docker Compose)
- Git

---

## 🚀 Instalación y Ejecución

Sigue estos pasos para levantar el entorno de desarrollo local:

### 1.Configurar el Proyecto
```bash
cd reserva-hoteles/backend
docker compose up --build
```

### 2. Verificar el estado y hacer logs
#### User API
docker logs -f user_api

#### Reservation API
docker logs -f reservation_api

#### Hotels API
docker logs -f hotels_api

### 3. Para poder inspeccionar datos manualmente

### MYSQL

#### user_api
docker exec -it mysql_users sh
 Luego: mysql -u root -p

 reservation_api 

docker exec -it mysql_reservation sh
 Luego: mysql -u root -p

### MONGODB
#### Hotels_api
docker exec -it mongo_hoteles mongosh
