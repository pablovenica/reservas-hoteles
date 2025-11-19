# 🚀 Search API - Quick Start Guide

## Inicio Rápido en 5 Minutos

### Opción 1: Con Docker Compose (Recomendado)

```bash
# 1. Navega a la carpeta backend
cd backend

# 2. Levanta todos los servicios
docker-compose up -d

# 3. Verifica que todo esté corriendo
docker-compose ps

# 4. Ve los logs
docker-compose logs -f search_api
```

✅ El Search API estará disponible en `http://localhost:8084`

### Opción 2: Ejecución Local

```bash
# 1. Descarga dependencias
cd search_api
go mod download

# 2. Levanta MongoDB y RabbitMQ en Docker
docker run -d -p 27017:27017 --name mongodb mongo:7.0
docker run -d -p 5672:5672 -p 15672:15672 --name rabbitmq rabbitmq:3.13-management

# 3. Configura variables de entorno
export RABBITMQ_URL=amqp://guest:guest@localhost:5672/
export MONGODB_URI=mongodb://localhost:27017/

# 4. Ejecuta el servidor
go run main.go
```

## Primer Uso

### 1. Obtén un JWT Token

Para los ejemplos, necesitas un token JWT. Si no tienes uno:

```bash
# Usa este token de ejemplo (reemplaza en los comandos)
JWT_TOKEN="tu-jwt-token-aqui"
```

### 2. Realiza una búsqueda

```bash
curl -X POST http://localhost:8084/search/hotels \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -d '{
    "hotel_name": "Hotel Paradise",
    "city": "Madrid",
    "check_in": "2024-12-01",
    "check_out": "2024-12-05",
    "guests": 2
  }'
```

### 3. Obtén el historial

```bash
curl -X GET http://localhost:8084/search/history \
  -H "Authorization: Bearer $JWT_TOKEN"
```

## Herramientas de Administración

| Herramienta | URL | Usuario | Contraseña |
|-------------|-----|---------|-----------|
| RabbitMQ Management | http://localhost:15672 | guest | guest |
| MongoDB Express | http://localhost:8081 | - | - |

## Carpetas Importantes

```
search_api/
├── README.md                    # Documentación del API
├── DEPLOYMENT.md                # Guía de despliegue
├── RABBITMQ_INTEGRATION.md     # Cómo integrar con otros servicios
├── IMPLEMENTATION_SUMMARY.md   # Resumen de implementación
└── examples.sh                  # Ejemplos de cURL
```

## Solución de Problemas

### ❌ "No se puede conectar a RabbitMQ"

```bash
# Verifica que RabbitMQ está corriendo
docker-compose ps rabbitmq

# Reinicia RabbitMQ
docker-compose restart rabbitmq
```

### ❌ "MongoDB no responde"

```bash
# Verifica MongoDB
docker-compose ps mongo_search

# Reinicia MongoDB
docker-compose restart mongo_search
```

### ❌ "Error: Puerto 8084 ya está en uso"

```bash
# Encuentra el proceso
lsof -i :8084

# Mata el proceso (en Windows: taskkill /PID <PID> /F)
kill -9 <PID>
```

## Archivos Clave

| Archivo | Descripción |
|---------|-----------|
| `main.go` | Punto de entrada |
| `app/router.go` | Configuración de rutas |
| `service/search_service.go` | Lógica de negocio |
| `messaging/consumer.go` | Procesamiento asincrónico |
| `repository/search_repository.go` | Acceso a datos |

## Variables de Entorno

```env
# .env
RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
MONGODB_URI=mongodb://mongo_search:27017/
DATABASE_NAME=search_db
SEARCH_EXCHANGE=search_exchange
SEARCH_QUEUE=search_queue
```

## Estructura de la Búsqueda

### Request
```json
{
  "hotel_name": "string",
  "city": "string",
  "check_in": "YYYY-MM-DD",
  "check_out": "YYYY-MM-DD",
  "guests": integer
}
```

### Response
```json
{
  "message": "Búsqueda iniciada",
  "search": {
    "id": "uuid",
    "user_id": "user-id",
    "hotel_name": "string",
    "city": "string",
    "check_in": "YYYY-MM-DD",
    "check_out": "YYYY-MM-DD",
    "guests": integer,
    "timestamp": "ISO-8601",
    "status": "pending|processed"
  }
}
```

## Detener los Servicios

```bash
# Parar todos los servicios
docker-compose down

# Parar y eliminar volúmenes
docker-compose down -v

# Solo el search_api
docker-compose stop search_api
```

## Logs

```bash
# Ver logs en tiempo real
docker-compose logs -f search_api

# Últimas 100 líneas
docker-compose logs --tail=100 search_api

# Guardar logs en archivo
docker-compose logs search_api > search_api.log
```

## Pasos Siguientes

1. ✅ Lee `README.md` para detalles del API
2. ✅ Revisa `DEPLOYMENT.md` para despliegue en producción
3. ✅ Consulta `RABBITMQ_INTEGRATION.md` si usas otros microservicios
4. ✅ Importa `Search_API.postman_collection.json` en Postman

## Links Útiles

- 📖 [Go Documentation](https://golang.org/doc)
- 📚 [Gin Web Framework](https://github.com/gin-gonic/gin)
- 🐰 [RabbitMQ Documentation](https://www.rabbitmq.com/documentation.html)
- 🍃 [MongoDB Documentation](https://docs.mongodb.com)

---

**¿Necesitas ayuda?** Revisa los archivos de documentación en la carpeta.

**¡Listo para usar! 🎉**
