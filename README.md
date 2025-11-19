# 🏨 Sistema de Reservas de Hoteles

## 🏗 Arquitectura del Proyecto

El sistema está compuesto por tres servicios independientes, cada uno
con responsabilidades específicas y base de datos dedicada:

  -----------------------------------------------------------------------
  Servicio                Responsabilidad         Base de Datos
  ----------------------- ----------------------- -----------------------
  **User API**            Gestión de usuarios y   MySQL (`mysql_users`)
                          autenticación           

  **Hotels API**          Catálogo, búsqueda y    MongoDB
                          gestión de hoteles      (`mongo_hoteles`)

  **Reservation API**     Lógica transaccional de MySQL
                          reservas y              (`mysql_reservation`)
                          disponibilidad          
  -----------------------------------------------------------------------

------------------------------------------------------------------------

## 🛠 Stack Tecnológico

### Infraestructura

-   **Docker & Docker Compose** - Orquestación de contenedores
-   **MySQL** - Base de datos relacional para transacciones y usuarios
-   **MongoDB** - Base de datos NoSQL para el catálogo de hoteles

## 📋 Pre-requisitos

Antes de comenzar, asegúrate de tener instalado:

-   [Docker Desktop](https://www.docker.com/products/docker-desktop) (o
    Docker Engine + Docker Compose)
-   Git

------------------------------------------------------------------------

## 🚀 Instalación y Ejecución

Sigue estos pasos para levantar el entorno de desarrollo local:

### 1. Configurar el Proyecto (Backend)

``` bash
cd reserva-hoteles/backend
docker compose up --build
```

### 2. Verificar el estado y hacer logs (Backend)

#### User API

``` bash
docker logs -f user_api
```

#### Reservation API

``` bash
docker logs -f reservation_api
```

#### Hotels API

``` bash
docker logs -f hotels_api
```

### 3. Ejecutar la Aplicación Frontend 💻

``` bash
cd ../frontend
npm install
npm install react-router-dom
npm run dev
```

Luego abre: http://localhost:5173

------------------------------------------------------------------------

### 4. Inspeccionar datos manualmente (Bases de Datos)

#### MYSQL

**user_api**

``` bash
docker exec -it mysql_users sh
mysql -u root -p
```

**reservation_api**

``` bash
docker exec -it mysql_reservation sh
mysql -u root -p
```

------------------------------------------------------------------------

### MONGODB

#### Hotels_api

``` bash
docker exec -it mongo_hoteles mongosh
```
