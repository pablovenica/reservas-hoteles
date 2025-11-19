# Guía de Integración: Search API con RabbitMQ

## Para Otros Microservicios

Si quieres que otros microservicios publiquen mensajes de búsqueda en el Search API a través de RabbitMQ, sigue esta guía.

## Dependencias Go a Agregar

```bash
go get github.com/rabbitmq/amqp091-go
```

## Ejemplo de Integración en Otro Servicio

### 1. Crear paquete de RabbitMQ en tu servicio

```go
package messaging

import (
    "context"
    "encoding/json"
    "os"
    
    amqp "github.com/rabbitmq/amqp091-go"
    log "github.com/sirupsen/logrus"
)

var (
    Connection *amqp.Connection
    Channel    *amqp.Channel
)

func InitRabbitMQ() error {
    var err error
    rabbitmqURL := os.Getenv("RABBITMQ_URL")
    if rabbitmqURL == "" {
        rabbitmqURL = "amqp://guest:guest@localhost:5672/"
    }

    Connection, err = amqp.Dial(rabbitmqURL)
    if err != nil {
        return err
    }

    Channel, err = Connection.Channel()
    if err != nil {
        return err
    }

    return nil
}

func PublishSearchMessage(userID, hotelName, city, checkIn, checkOut string, guests int) error {
    message := map[string]interface{}{
        "user_id":    userID,
        "hotel_name": hotelName,
        "city":       city,
        "check_in":   checkIn,
        "check_out":  checkOut,
        "guests":     guests,
        "timestamp":  time.Now().String(),
    }

    body, err := json.Marshal(message)
    if err != nil {
        return err
    }

    return Channel.PublishWithContext(
        context.Background(),
        "search_exchange", // exchange name
        "search.created",  // routing key
        false,             // mandatory
        false,             // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    )
}

func CloseConnection() {
    if Channel != nil {
        Channel.Close()
    }
    if Connection != nil {
        Connection.Close()
    }
}
```

### 2. Usar en tu controlador

```go
import (
    "your_api/messaging"
)

func InitiateSearch(ctx *gin.Context) {
    // ... validar datos ...

    userID := ctx.GetString("user_id")
    
    err := messaging.PublishSearchMessage(
        userID,
        hotelName,
        city,
        checkIn,
        checkOut,
        guests,
    )

    if err != nil {
        log.Errorf("Error publishing search: %v", err)
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": "Error al procesar búsqueda",
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "Búsqueda enviada a procesar",
    })
}
```

## Docker Compose

Asegúrate que tu docker-compose.yml incluya RabbitMQ:

```yaml
rabbitmq:
  image: rabbitmq:3.13-management
  container_name: rabbitmq
  environment:
    RABBITMQ_DEFAULT_USER: guest
    RABBITMQ_DEFAULT_PASS: guest
  ports:
    - "5672:5672"      # AMQP port
    - "15672:15672"    # Management UI
  networks:
    - backend-net
```

## Variables de Entorno

En tu `.env`:

```env
RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
```

## Flujo de Comunicación

1. **Servicio A** publica un mensaje `search.created` en el exchange `search_exchange`
2. **RabbitMQ** enruta el mensaje a `search_queue` (basado en el routing key `search.*`)
3. **Search API** consumer recibe el mensaje
4. **Search API** procesa y guarda en MongoDB
5. El usuario puede consultar el historial en el Search API

## Topics/Routing Keys Disponibles

- `search.created` - Nueva búsqueda creada
- `search.*` - Coincide con cualquier evento de búsqueda

## Management UI

Accede a RabbitMQ Management:
- URL: `http://localhost:15672`
- Usuario: `guest`
- Contraseña: `guest`

## Troubleshooting

### Mensaje no se procesa
- Verifica que RabbitMQ esté corriendo
- Revisa logs del consumer: `docker logs search_api`
- Asegúrate que el routing key sea `search.*`

### Conexión rechazada
- Verifica las credenciales en `.env`
- Asegúrate que RabbitMQ esté en la misma red Docker
- Revisa puertos expuestos

### Dead Letter Queue
Implementar DLQ para mensajes que fallan:

```go
err = Channel.QueueDeclare(
    "search_queue_dlq",
    true,
    false,
    false,
    false,
    amqp.Table{
        "x-message-ttl":    86400000,
        "x-max-length":     10000,
    },
)
```
