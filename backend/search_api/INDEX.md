# 📚 Search API - Índice de Documentación

## 🚀 Para Empezar Rápido

- **[QUICK_START.md](QUICK_START.md)** - ⭐ **EMPIEZA AQUÍ**
  - Inicio en 5 minutos
  - Comandos básicos
  - Solución de problemas rápida

## 📖 Documentación Principal

### API Documentation
- **[README.md](README.md)** - Documentación completa del API
  - Características
  - Stack tecnológico
  - Estructura del proyecto
  - Endpoints detallados
  - Flujo de procesamiento
  - Configuración
  - Compilación y ejecución

### Despliegue
- **[DEPLOYMENT.md](DEPLOYMENT.md)** - Guía completa de despliegue
  - Despliegue con Docker Compose
  - Despliegue local sin Docker
  - Troubleshooting
  - Monitoreo
  - Escalado
  - Respaldo y recuperación
  - Seguridad en producción
  - Performance tuning

### Integración con Otros Servicios
- **[RABBITMQ_INTEGRATION.md](RABBITMQ_INTEGRATION.md)** - Cómo usar desde otros microservicios
  - Ejemplos de código
  - Configuración de RabbitMQ
  - Flujo de comunicación
  - Topics y routing keys
  - Management UI
  - Troubleshooting

## 📋 Resúmenes y Checklists

- **[IMPLEMENTATION_SUMMARY.md](IMPLEMENTATION_SUMMARY.md)** - Resumen ejecutivo
  - Estructura del proyecto
  - Arquitectura
  - Componentes principales
  - Features implementadas
  - Dependencias
  - Próximos pasos

- **[CHECKLIST.md](CHECKLIST.md)** - Checklist completo de implementación
  - Todas las características marcadas
  - Estado de cada componente
  - Verificación de implementación

## 🧪 Testing y Ejemplos

- **[examples.sh](examples.sh)** - Scripts de prueba con cURL
  - Ejemplos básicos
  - Casos de prueba
  - Manejo de errores
  - Tests sin autenticación

- **[Search_API.postman_collection.json](Search_API.postman_collection.json)** - Collection para Postman
  - Importa en Postman
  - Todos los endpoints
  - Variables de entorno

## 🗂️ Estructura del Código

### Carpetas Principales
```
search_api/
├── app/              # Configuración de aplicación
├── bd/               # Base de datos
├── controller/       # Endpoints HTTP
├── domain/          # Modelos de dominio
├── dto/             # Data Transfer Objects
├── middleware/      # Middleware (autenticación)
├── messaging/       # Integración RabbitMQ
├── repository/      # Acceso a datos
└── service/         # Lógica de negocio
```

### Archivos Importantes
- **main.go** - Punto de entrada
- **go.mod** - Dependencias
- **Dockerfile** - Imagen Docker
- **.env** - Variables de entorno

## 🔑 Endpoints del API

Todos requieren `Authorization: Bearer <token>`

| Método | Endpoint | Archivo |
|--------|----------|---------|
| POST | `/search/hotels` | [search_controller.go](controller/search_controller.go) |
| GET | `/search/history` | [search_controller.go](controller/search_controller.go) |
| GET | `/search/history/:id` | [search_controller.go](controller/search_controller.go) |
| DELETE | `/search/history/:id` | [search_controller.go](controller/search_controller.go) |

## 🏗️ Componentes Clave

### Application
- [app_config.go](app/app_config.go) - Logger
- [router.go](app/router.go) - Setup Gin
- [url_mapping.go](app/url_mapping.go) - Rutas

### Data Access
- [mongo.go](bd/mongo.go) - Conexión MongoDB
- [search_repository.go](repository/search_repository.go) - CRUD

### Business Logic
- [search_service.go](service/search_service.go) - Lógica de búsquedas

### RabbitMQ
- [rabbitmq.go](messaging/rabbitmq.go) - Conexión
- [consumer.go](messaging/consumer.go) - Procesamiento
- [publisher.go](messaging/publisher.go) - Publicación

### Security
- [auth.go](middleware/auth.go) - Autenticación JWT

## 🔧 Configuración

### Variables de Entorno (.env)
```env
RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
SEARCH_EXCHANGE=search_exchange
SEARCH_QUEUE=search_queue
SEARCH_ROUTING_KEY=search.*
MONGODB_URI=mongodb://mongo_search:27017/
DATABASE_NAME=search_db
```

### Docker Compose
- RabbitMQ: puerto 5672 (AMQP), 15672 (Management)
- MongoDB: puerto 27017
- Search API: puerto 8084

## 📊 Datos

### Estructura de Búsqueda en MongoDB
```json
{
  "_id": "uuid",
  "user_id": "string",
  "hotel_name": "string",
  "city": "string",
  "check_in": "YYYY-MM-DD",
  "check_out": "YYYY-MM-DD",
  "guests": integer,
  "timestamp": "ISO-8601",
  "status": "pending|processed"
}
```

## 📞 Recursos Externos

- [Go Documentation](https://golang.org/doc)
- [Gin Framework](https://github.com/gin-gonic/gin)
- [RabbitMQ Docs](https://www.rabbitmq.com/documentation.html)
- [MongoDB Docs](https://docs.mongodb.com)
- [JWT Token](https://jwt.io)

## 🎯 Rutas de Navegación

### Si quieres...

**Empezar rápido**
→ [QUICK_START.md](QUICK_START.md)

**Entender la arquitectura**
→ [README.md](README.md) + [IMPLEMENTATION_SUMMARY.md](IMPLEMENTATION_SUMMARY.md)

**Desplegar en producción**
→ [DEPLOYMENT.md](DEPLOYMENT.md)

**Integrar con otro microservicio**
→ [RABBITMQ_INTEGRATION.md](RABBITMQ_INTEGRATION.md)

**Probar los endpoints**
→ [examples.sh](examples.sh) o importar en Postman

**Verificar que todo está implementado**
→ [CHECKLIST.md](CHECKLIST.md)

## 🆘 Solución de Problemas

**Problema** → **Archivo**
- Conexión a RabbitMQ → [DEPLOYMENT.md](DEPLOYMENT.md#troubleshooting)
- Conexión a MongoDB → [DEPLOYMENT.md](DEPLOYMENT.md#troubleshooting)
- Endpoints no funcionan → [README.md](README.md#endpoints)
- Docker no construye → [DEPLOYMENT.md](DEPLOYMENT.md#troubleshooting)

## ✅ Estado Actual

**Versión**: 1.0.0  
**Fecha**: Noviembre 2024  
**Status**: ✅ Completado y listo para usar

## 📝 Licencia

Este proyecto es parte del sistema de reservas de hoteles.

---

**Última actualización**: Noviembre 2024

Usa este índice como guía para navegar por toda la documentación.
