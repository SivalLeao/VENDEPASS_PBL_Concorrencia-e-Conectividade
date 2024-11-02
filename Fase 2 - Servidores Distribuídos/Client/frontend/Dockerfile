# Usando a imagem base do Node.js
FROM node:20-slim AS base

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos package.json e package-lock.json antes de copiar todo o código (melhor para cache)
COPY package*.json ./

# Instala as dependências de produção e build em camadas separadas
FROM base AS prod-deps
RUN npm install --only=production --frozen-lockfile

# Instala todas as dependências e faz o build
FROM base AS build
RUN npm install --frozen-lockfile
COPY . .
RUN npm run build

# Imagem final onde as dependências de produção e o build são copiados
FROM base

# Copia as dependências de produção e o diretório de build
COPY --from=prod-deps /app/node_modules /app/node_modules
COPY --from=build /app/dist /app/dist

# Copia o diretório de imagens para o local correto
COPY ./public /app/image

# Expõe a porta que o servidor frontend vai usar
EXPOSE 3000

# Comando para iniciar a aplicação
CMD [ "npm", "run", "serve" ]
