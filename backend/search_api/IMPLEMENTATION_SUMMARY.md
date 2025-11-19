# Search API - Resumen de Implementación

## ✅ Completado

Se ha implementado un **microservicio completo de búsqueda de hoteles** con integración de **RabbitMQ** para procesamiento asincrónico.

## 📁 Estructura del Proyecto

```
search_api/
├── app/
│   ├── app_config.go      # Configuración del logger
│   ├── router.go          # Setup del router Gin
│   └── url_mapping.go     # Mapeo de rutas y middleware
│
├── bd/
│   └── mongo.go           # Conexión y configuración de MongoDB
│
├── controller/
│   └── search_controller.go  # Endpoints del API
│
├── domain/
│   └── search.go          # Modelos de dominio
│
├── dto/
│   └── search_dto.go      # Data Transfer Objects
│
├── middleware/
│   └── auth.go            # Autenticación JWT
│
├── messaging/
│   ├── rabbitmq.go        # Conexión RabbitMQ
│   ├── consumer.go        # Consumer de mensajes
│   └── publisher.go       # Publisher de mensajes
│
├── repository/
│   └── search_repository.go  # Acceso a datos MongoDB
│
├── service/
│   └── search_service.go  # Lógica de negocio
│
├── main.go                # Punto de entrada
├── go.mod                 # Dependencias Go
├── Dockerfile             # Imagen Docker
├── .env                   # Variables de entorno
├── README.md              # Documentación del API
├── RABBITMQ_INTEGRATION.md # Guía de integración con otros servicios
├── DEPLOYMENT.md          # Guía de despliegue
├── examples.sh            # Ejemplos de cURL
└── Search_API.postman_collection.json  # Collection de Postman
```

## 🏗️ Arquitectura

### Componentes Principales

1. **API REST (Gin)** - Puerto 8084
   - Endpoints protegidos con JWT
   - CORS configurado para frontend

2. **RabbitMQ** - Puerto 5672
   - Exchange: `search_exchange` (topic)
   - Queue: `search_queue`
   - Routing key: `search.*`

3. **MongoDB** - Puerto 27017
   - Base de datos: `search_db`
   - Colección: `searches`

### Flujo de Datos

```
Usuario → API REST → Service → RabbitMQ → Consumer → MongoDB
                         ↓
                    Respuesta inmediata
```

## 🔌 Endpoints

Todos requieren `Authorization: Bearer <token>` en el header.

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| POST | `/search/hotels` | Realizar búsqueda |
| GET | `/search/history` | Obtener historial |
| GET | `/search/history/:id` | Obtener búsqueda específica |
| DELETE | `/search/history/:id` | Eliminar búsqueda |

## 🛠️ Dependencias Go

```
github.com/gin-contrib/cors v1.7.6
github.com/gin-gonic/gin v1.11.0
github.com/golang-jwt/jwt/v5 v5.3.0
github.com/rabbitmq/amqp091-go v1.10.1
github.com/sirupsen/logrus v1.9.3
go.mongodb.org/mongo-driver v1.17.6
```

## 🚀 Quick Start

### Con Docker Compose

```bash
cd backend
docker-compose up -d search_api
```

### Localmente

```bash
cd search_api
go mod download
export RABBITMQ_URL=amqp://guest:guest@localhost:5672/
export MONGODB_URI=mongodb://localhost:27017/
go run main.go
```

## 📋 Características

✅ Búsqueda de hoteles con múltiples filtros  
✅ Historial de búsquedas por usuario  
✅ Procesamiento asincrónico con RabbitMQ  
✅ Almacenamiento persistente en MongoDB  
✅ Autenticación JWT  
✅ CORS habilitado  
✅ Logging con Logrus  
✅ Dockerizado  
✅ Documentación completa  

## 📊 Bases de Datos

### MongoDB - Colección `searches`

```json
{
  "_id": "uuid-string",
  "user_id": "user-123",
  "hotel_name": "Hotel Paradise",
  "city": "Madrid",
  "check_in": "2024-12-01",
  "check_out": "2024-12-05",
  "guests": 2,
  "timestamp": "2024-11-19T10:30:00Z",
  "status": "processed"
}
```

## 🔄 Integración con Otros Microservicios

Otros servicios pueden publicar mensajes de búsqueda en RabbitMQ:

```go
// En otro microservicio
messaging.PublishMessage("search.created", searchData)
```

Ver `RABBITMQ_INTEGRATION.md` para detalles.

## 📝 Documentación

- **README.md** - API Documentation
- **RABBITMQ_INTEGRATION.md** - Integración con RabbitMQ
- **DEPLOYMENT.md** - Guía de despliegue en producción
- **examples.sh** - Ejemplos de cURL
- **Search_API.postman_collection.json** - Collection para Postman

## 🧪 Testing

### Con cURL

```bash
curl -X POST http://localhost:8084/search/hotels \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your-jwt-token" \
  -d '{
    "hotel_name": "Hotel Example",
    "city": "Madrid",
    "check_in": "2024-12-01",
    "check_out": "2024-12-05",
    "guests": 2
  }'
```

### Con Postman

Importa `Search_API.postman_collection.json` en Postman.

## 📈 Monitoreo

- **RabbitMQ Management**: http://localhost:15672
- **MongoDB Express**: http://localhost:8081
- **Logs**: `docker-compose logs -f search_api`

## 🔐 Seguridad

- Autenticación JWT en todos los endpoints
- CORS limitado a `http://localhost:5173`
- Validación de entrada en DTOs
- Manejo seguro de conexiones a bases de datos

## 🎯 Próximos Pasos

1. Actualizar credenciales de RabbitMQ en producción
2. Configurar TLS/SSL para HTTPS
3. Implementar Dead Letter Queue para mensajes fallidos
4. Agregar índices en MongoDB para optimización
5. Implementar caché con Redis (opcional)
6. Agregar metrics con Prometheus (opcional)
7. Configurar ELK Stack para logging centralizado (opcional)

## 📞 Soporte

Ver documentación en:
- `README.md` - Uso del API
- `DEPLOYMENT.md` - Problemas comunes
- `RABBITMQ_INTEGRATION.md` - Integración

---

**Versión**: 1.0.0  
**Fecha**: Noviembre 2024  
**Status**: ✅ Listo para usar
