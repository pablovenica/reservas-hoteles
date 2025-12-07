# Usamos Node 22 en versión Alpine (ligera y estable)
FROM node:22-alpine

WORKDIR /usr/local/app

# Copiamos primero los archivos de dependencias para aprovechar la caché
COPY package*.json ./

# Instalamos dependencias
RUN npm install

# --- AQUÍ ESTABA EL ERROR ---
# Copiamos todo el código fuente del proyecto al contenedor
COPY . .

# Exponemos el puerto de Vite
EXPOSE 5173

# Arrancamos el servidor de desarrollo
CMD ["npm", "run", "dev"]