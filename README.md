
## 🏗 Arquitectura del Proyecto

El sistema se divide en tres servicios principales, cada uno con su propia responsabilidad y base de datos:

| Servicio | Responsabilidad | Base de Datos |
| :--- | :--- | :--- |
| **User API** | Gestión de usuarios y autenticación. | MySQL (`mysql_users`) |
| **Hotels API** | Catálogo, búsqueda y gestión de información de hoteles. | MongoDB (`mongo_hoteles`) |
| **Reservation API** | Lógica transaccional de reservas y disponibilidad. | MySQL (`mysql_reservation`) |

---

## 🛠 Tecnologías Utilizadas

* **Docker & Docker Compose**: Orquestación de contenedores.
* **MySQL**: Base de datos relacional para transacciones y usuarios.
* **MongoDB**: Base de datos NoSQL para el catálogo de hoteles.
* **[Tu Lenguaje/Framework]**: (Ej: Node.js / Java Spring Boot / Python).

---

## 📋 Pre-requisitos

Antes de comenzar, asegúrate de tener instalado:
* [Docker Desktop](https://www.docker.com/products/docker-desktop) (o Docker Engine + Compose).
* Git.

---

## 🚀 Instalación y Puesta en Marcha

Sigue estos pasos para levantar el entorno de desarrollo local.

### 1. Levantar los Contenedores
Ubícate en la raíz del proyecto y navega a la carpeta del backend para iniciar los servicios.

# Navegar al directorio del backend
cd ./reserva-hoteles/backend

# Construir y levantar los servicios
docker compose up --build

### 2. Verificar el estado y hacer logs
# User API
docker logs -f user_api

# Reservation API
docker logs -f reservation_api

# Hotels API
docker logs -f hotels_api

### 3. Para poder inspeccionar datos manualmente

## MYSQL

# user_api
docker exec -it mysql_users sh
# Luego: mysql -u root -p

# reservation_api 

docker exec -it mysql_reservation sh
# Luego: mysql -u root -p

## MONGODB
# Hotels_api
docker exec -it mongo_hoteles mongosh
