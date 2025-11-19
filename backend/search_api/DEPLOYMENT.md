# Guía de Despliegue - Search API

## Requisitos Previos

- Docker y Docker Compose instalados
- Go 1.25.4+ (para desarrollo local)
- Git

## Despliegue con Docker Compose

### 1. Estructura del Proyecto

```
backend/
├── docker-compose.yml      # Orquestación de contenedores
├── search_api/             # Este microservicio
├── hotels_api/
├── user_api/
└── reservation_api/
```

### 2. Variables de Entorno

Crea el archivo `.env` en `backend/search_api/.env`:

```env
# RabbitMQ
RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
SEARCH_EXCHANGE=search_exchange
SEARCH_QUEUE=search_queue
SEARCH_ROUTING_KEY=search.*

# MongoDB
MONGODB_URI=mongodb://mongo_search:27017/
DATABASE_NAME=search_db
```

### 3. Levantar los Servicios

```bash
# Desde la carpeta backend/
cd backend

# Levantar todos los servicios
docker-compose up -d

# Ver logs del search_api
docker-compose logs -f search_api

# Ver logs del RabbitMQ
docker-compose logs -f rabbitmq
```

### 4. Verificar que Todo Está Funcionando

```bash
# Ver contenedores en ejecución
docker-compose ps

# Probar el health de RabbitMQ
curl -u guest:guest http://localhost:15672/api/healthchecks/node

# Probar el API
curl -X GET http://localhost:8084/search/history \
  -H "Authorization: Bearer your-jwt-token"
```

## Acceso a Herramientas

### RabbitMQ Management
- URL: `http://localhost:15672`
- Usuario: `guest`
- Contraseña: `guest`

### MongoDB Express
- URL: `http://localhost:8081`
- Servidor: `mongo_search`

## Despliegue Local (Sin Docker)

### 1. Instalar Dependencias

```bash
cd search_api
go mod download
go mod tidy
```

### 2. Servicios Externos Necesarios

```bash
# MongoDB
docker run -d -p 27017:27017 --name mongodb mongo:7.0

# RabbitMQ
docker run -d -p 5672:5672 -p 15672:15672 --name rabbitmq rabbitmq:3.13-management
```

### 3. Configurar Variables de Entorno

```bash
export RABBITMQ_URL=amqp://guest:guest@localhost:5672/
export MONGODB_URI=mongodb://localhost:27017/
export DATABASE_NAME=search_db
```

### 4. Ejecutar el Servidor

```bash
go run main.go
```

El servidor estará disponible en `http://localhost:8084`

## Troubleshooting

### RabbitMQ no se conecta
```bash
# Verificar que RabbitMQ está corriendo
docker-compose ps rabbitmq

# Ver logs
docker-compose logs rabbitmq

# Reiniciar
docker-compose restart rabbitmq
```

### MongoDB no se conecta
```bash
# Verificar que MongoDB está corriendo
docker-compose ps mongo_search

# Ver logs
docker-compose logs mongo_search

# Reiniciar
docker-compose restart mongo_search
```

### Build fallan en Docker
```bash
# Eliminar caché
docker-compose build --no-cache search_api

# Rebuildar
docker-compose up -d search_api
```

### Ver logs en tiempo real
```bash
# Todos los logs
docker-compose logs -f

# Solo search_api
docker-compose logs -f search_api

# Últimas 100 líneas
docker-compose logs --tail=100 search_api
```

## Monitoreo

### Verificar Estado de Procesamiento

```bash
# Ver colas en RabbitMQ
curl -u guest:guest http://localhost:15672/api/queues

# Ver búsquedas en MongoDB
docker-compose exec mongo_search mongosh search_db --eval "db.searches.find()"

# Estadísticas de la cola
curl -u guest:guest http://localhost:15672/api/queues/%2F/search_queue
```

## Escalado

### Aumentar Replicas del Search API

Edita `docker-compose.yml`:

```yaml
search_api_1:
  build:
    context: ./search_api
  container_name: search_api_1
  # ... resto de config

search_api_2:
  build:
    context: ./search_api
  container_name: search_api_2
  # ... resto de config
```

### Load Balancer (nginx)

```yaml
nginx:
  image: nginx:alpine
  ports:
    - "80:80"
  volumes:
    - ./nginx.conf:/etc/nginx/nginx.conf
  depends_on:
    - search_api_1
    - search_api_2
```

## Respaldo y Recuperación

### Respaldo de MongoDB

```bash
# Crear backup
docker-compose exec mongo_search mongodump --out=/backup

# Restaurar
docker-compose exec mongo_search mongorestore /backup
```

### Copiar volúmenes

```bash
# Copiar data de mongo_search_data
docker run --rm -v mongo_search_data:/data -v $(pwd):/backup \
  alpine tar czf /backup/mongo_backup.tar.gz -C /data .
```

## Seguridad

### Cambiar Credenciales de RabbitMQ

1. Edita `docker-compose.yml`:
```yaml
rabbitmq:
  environment:
    RABBITMQ_DEFAULT_USER: tu_usuario
    RABBITMQ_DEFAULT_PASS: tu_contraseña_fuerte
```

2. Actualiza `.env`:
```env
RABBITMQ_URL=amqp://tu_usuario:tu_contraseña_fuerte@rabbitmq:5672/
```

### Usar HTTPS (Producción)

Configura un reverse proxy con TLS (nginx, Caddy, etc.)

## Performance Tuning

### MongoDB Indexes

```bash
docker-compose exec mongo_search mongosh search_db --eval "
db.searches.createIndex({ 'user_id': 1, 'timestamp': -1 })
db.searches.createIndex({ '_id': 1 })
"
```

### RabbitMQ Memory

Edita `docker-compose.yml`:
```yaml
rabbitmq:
  environment:
    RABBITMQ_DEFAULT_USER: guest
    RABBITMQ_DEFAULT_PASS: guest
    RABBITMQ_VM_MEMORY_HIGH_WATERMARK: 2GB
```

## Métricas y Logs

### Exportar logs
```bash
docker-compose logs > docker_logs.txt
docker-compose logs search_api > search_api_logs.txt
```

### Usar ELK Stack (Elasticsearch, Logstash, Kibana)

Opcional pero recomendado para producción.

## Rollback de Versión

```bash
# Cambiar a versión anterior en docker-compose.yml
docker-compose down search_api
docker-compose up -d search_api
```

## Checklist de Despliegue

- [ ] Variables de entorno configuradas
- [ ] RabbitMQ corriendo y accesible
- [ ] MongoDB corriendo y accesible
- [ ] Puertos disponibles (8084, 5672, 15672, 27017)
- [ ] Backups configurados
- [ ] Logs configurados
- [ ] Monitoreo en lugar
- [ ] Tests pasando
- [ ] Documentación actualizada
