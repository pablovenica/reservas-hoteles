# Search API

Microservicio de búsqueda de hoteles integrado con RabbitMQ para procesamiento asincrónico de búsquedas.

## Características

- **Búsqueda de hoteles**: Buscar hoteles por ciudad, fechas y cantidad de huéspedes
- **Historial de búsquedas**: Ver el historial de búsquedas del usuario
- **Procesamiento asincrónico**: Las búsquedas se procesan a través de RabbitMQ
- **Almacenamiento en MongoDB**: Persistencia de datos de búsquedas
- **Autenticación JWT**: Protección de endpoints

## Stack Tecnológico

- **Go 1.25.4**
- **Gin** (Web Framework)
- **MongoDB** (Base de datos)
- **RabbitMQ** (Message Broker)
- **JWT** (Autenticación)
- **Docker** (Containerización)

## Estructura del Proyecto

```
search_api/
├── app/                    # Configuración de la aplicación
│   ├── app_config.go      # Configuración principal
│   ├── router.go          # Setup del router
│   └── url_mapping.go     # Mapeo de rutas
├── bd/                    # Configuración de base de datos
│   └── mongo.go           # Conexión con MongoDB
├── controller/            # Controladores
│   └── search_controller.go
├── domain/               # Modelos de dominio
│   └── search.go
├── dto/                  # Data Transfer Objects
│   └── search_dto.go
├── middleware/           # Middleware
│   └── auth.go          # Autenticación JWT
├── messaging/           # Integración con RabbitMQ
│   ├── rabbitmq.go      # Configuración de RabbitMQ
│   └── consumer.go      # Consumer de mensajes
├── repository/          # Acceso a datos
│   └── search_repository.go
├── service/            # Lógica de negocio
│   └── search_service.go
├── main.go            # Punto de entrada
├── go.mod             # Dependencias
├── Dockerfile         # Imagen Docker
└── .env               # Variables de entorno
```

## Variables de Entorno

```env
RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
SEARCH_EXCHANGE=search_exchange
SEARCH_QUEUE=search_queue
SEARCH_ROUTING_KEY=search.*
MONGODB_URI=mongodb://mongo_search:27017/
DATABASE_NAME=search_db
```

## Endpoints

Todos los endpoints requieren autenticación JWT en el header `Authorization: Bearer <token>`

### POST /search/hotels
Realizar una búsqueda de hoteles

**Request:**
```json
{
  "hotel_name": "Hotel Example",
  "city": "Madrid",
  "check_in": "2024-12-01",
  "check_out": "2024-12-05",
  "guests": 2
}
```

**Response:**
```json
{
  "message": "Búsqueda iniciada",
  "search": {
    "id": "uuid-string",
    "user_id": "user-id",
    "hotel_name": "Hotel Example",
    "city": "Madrid",
    "check_in": "2024-12-01",
    "check_out": "2024-12-05",
    "guests": 2,
    "timestamp": "2024-11-19T10:30:00Z",
    "status": "pending"
  }
}
```

### GET /search/history
Obtener el historial de búsquedas del usuario

**Response:**
```json
{
  "searches": [
    {
      "id": "uuid-string",
      "user_id": "user-id",
      "hotel_name": "Hotel Example",
      "city": "Madrid",
      "timestamp": "2024-11-19T10:30:00Z"
    }
  ]
}
```

### GET /search/history/:id
Obtener una búsqueda específica

**Response:**
```json
{
  "id": "uuid-string",
  "user_id": "user-id",
  "hotel_name": "Hotel Example",
  "city": "Madrid",
  "check_in": "2024-12-01",
  "check_out": "2024-12-05",
  "guests": 2,
  "timestamp": "2024-11-19T10:30:00Z",
  "status": "processed"
}
```

### DELETE /search/history/:id
Eliminar una búsqueda

**Response:**
```json
{
  "message": "Búsqueda eliminada"
}
```

## Flujo de Procesamiento

1. El usuario realiza una búsqueda (POST /search/hotels)
2. Se crea un registro con estado "pending" en MongoDB
3. El mensaje se publica en el exchange RabbitMQ
4. El consumer recibe el mensaje de la cola
5. Se actualiza el estado del registro a "processed"
6. El usuario puede consultar el historial y detalles de sus búsquedas

## Configuración de RabbitMQ

- **Exchange**: `search_exchange` (tipo: topic)
- **Queue**: `search_queue`
- **Routing Key**: `search.*`

## Compilación y Ejecución

### Con Docker Compose
```bash
docker-compose up -d search_api
```

### Localmente
```bash
go mod download
go run main.go
```

## Variables de Sistema

- **Puerto**: 8084
- **Base de datos**: MongoDB (search_db)
- **Message Broker**: RabbitMQ

## Dependencias Go

- github.com/gin-contrib/cors v1.7.6
- github.com/gin-gonic/gin v1.11.0
- github.com/golang-jwt/jwt/v5 v5.3.0
- github.com/rabbitmq/amqp091-go v1.10.1
- github.com/sirupsen/logrus v1.9.3
- go.mongodb.org/mongo-driver v1.17.6
