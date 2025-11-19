#!/bin/bash

# Ejemplos de uso del Search API

# Variables
API_URL="http://localhost:8084/search"
JWT_TOKEN="your-jwt-token-here"

# Colores para output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== Search API - Ejemplos de Uso ===${NC}\n"

# 1. Realizar una búsqueda de hoteles
echo -e "${GREEN}1. POST /search/hotels - Realizar búsqueda${NC}"
curl -X POST "$API_URL/hotels" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -d '{
    "hotel_name": "Hotel Paradise",
    "city": "Madrid",
    "check_in": "2024-12-01",
    "check_out": "2024-12-05",
    "guests": 2
  }'
echo -e "\n\n"

# 2. Obtener historial de búsquedas
echo -e "${GREEN}2. GET /search/history - Obtener historial${NC}"
curl -X GET "$API_URL/history" \
  -H "Authorization: Bearer $JWT_TOKEN"
echo -e "\n\n"

# 3. Obtener búsqueda específica
echo -e "${GREEN}3. GET /search/history/:id - Obtener búsqueda específica${NC}"
# Reemplaza SEARCH_ID con un ID real
curl -X GET "$API_URL/history/{SEARCH_ID}" \
  -H "Authorization: Bearer $JWT_TOKEN"
echo -e "\n\n"

# 4. Eliminar búsqueda
echo -e "${GREEN}4. DELETE /search/history/:id - Eliminar búsqueda${NC}"
# Reemplaza SEARCH_ID con un ID real
curl -X DELETE "$API_URL/history/{SEARCH_ID}" \
  -H "Authorization: Bearer $JWT_TOKEN"
echo -e "\n\n"

# ===== Testing con jq (para parsear JSON) =====

echo -e "${BLUE}=== Ejemplos con jq (JSON parsing) ===${NC}\n"

# Búsqueda y guardar el ID para usar después
echo -e "${GREEN}Búsqueda y extracción de ID${NC}"
SEARCH_RESPONSE=$(curl -s -X POST "$API_URL/hotels" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $JWT_TOKEN" \
  -d '{
    "hotel_name": "Hotel Mountain",
    "city": "Barcelona",
    "check_in": "2024-12-10",
    "check_out": "2024-12-15",
    "guests": 4
  }')

echo "Respuesta completa:"
echo "$SEARCH_RESPONSE" | jq '.'

SEARCH_ID=$(echo "$SEARCH_RESPONSE" | jq -r '.search.id')
echo -e "\nID extraído: $SEARCH_ID"

# Usar el ID para consultar después
if [ "$SEARCH_ID" != "null" ] && [ -n "$SEARCH_ID" ]; then
  echo -e "\n${GREEN}Consultando búsqueda específica con ID: $SEARCH_ID${NC}"
  curl -s -X GET "$API_URL/history/$SEARCH_ID" \
    -H "Authorization: Bearer $JWT_TOKEN" | jq '.'
fi

echo -e "\n\n"

# ===== Casos de prueba sin autenticación =====

echo -e "${BLUE}=== Pruebas sin autenticación (Deben fallar) ===${NC}\n"

echo -e "${GREEN}POST /search/hotels sin token${NC}"
curl -X POST "$API_URL/hotels" \
  -H "Content-Type: application/json" \
  -d '{
    "hotel_name": "Test",
    "city": "Madrid",
    "check_in": "2024-12-01",
    "check_out": "2024-12-05",
    "guests": 2
  }' | jq '.'

echo -e "\n\n"

# ===== Casos de prueba con token inválido =====

echo -e "${BLUE}=== Pruebas con token inválido ===${NC}\n"

echo -e "${GREEN}GET /search/history con token inválido${NC}"
curl -X GET "$API_URL/history" \
  -H "Authorization: Bearer invalid-token" | jq '.'

echo -e "\n"
