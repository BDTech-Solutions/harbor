#!/usr/bin/env bash

set -e

echo "⚓ Harbor - WordPress bootstrap"

# 1. Verificar Docker
if ! command -v docker >/dev/null 2>&1; then
  echo "❌ Docker não está instalado."
  echo "Instale o Docker antes de continuar."
  exit 1
fi

# 2. Verificar Docker em execução
if ! docker info >/dev/null 2>&1; then
  echo "❌ Docker não está em execução."
  echo "Inicie o Docker e tente novamente."
  exit 1
fi

# 3. Criar .gitignore se não existir
if [[ ! -f .gitignore ]]; then
  echo "🌿 Criando .gitignore..."
  cat <<EOF > .gitignore
# WordPress core
wp/

# Uploads e cache
wp-content/uploads/
wp-content/cache/

# Docker volumes
docker/mysql_data
docker/php_data

# IDEs
.idea/
*.iml

# Logs
*.log
*.tmp
EOF
fi

# 4. Criar estrutura de pastas se necessário
echo "📂 Criando estrutura de pastas..."
mkdir -p docker/nginx
mkdir -p wp/wp-content/plugins
mkdir -p wp/wp-content/themes
mkdir -p wp/wp-content/uploads

# 5. Criar arquivo .env se não existir
if [[ ! -f .env ]]; then
  echo "🔧 Criando arquivo .env..."
  cat <<EOF > .env
# MySQL
DB_NAME=wordpress
DB_USER=wordpress
DB_PASSWORD=secret
DB_ROOT_PASSWORD=secret

# WordPress
WP_PORT=8080
EOF
fi

# 6. Se docker-compose.yml não existir, copiar o template
if [[ ! -f docker-compose.yml ]]; then
  echo "🐳 Copiando template docker-compose.yml..."
  cp "$HARBOR_ROOT/templates/wordpress/docker-compose.yml" ./docker-compose.yml
fi

# 7. Subir containers
echo "🚀 Iniciando containers..."
docker-compose up -d

echo "✅ Ambiente WordPress iniciado com sucesso."
echo "👉 Abra http://localhost:\$(grep WP_PORT .env | cut -d= -f2) para acessar o site."
