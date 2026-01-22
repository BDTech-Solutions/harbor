#!/usr/bin/env bash

check_docker() {
  # Verifica se Docker está instalado
  if ! command -v docker >/dev/null 2>&1; then
    echo "❌ Docker não encontrado."
    exit 1
  fi

  # Verifica se Docker está rodando
  if ! docker info >/dev/null 2>&1; then
    echo "❌ Docker não está em execução."
    exit 1
  fi
}

init_wordpress() {
  echo "⚓ Harbor - Inicializando projeto WordPress"

  # 1. Cria docker-compose.yml se não existir
  if [[ -f docker-compose.yml ]]; then
    echo "✅ docker-compose.yml já existe"
  else
    echo "📦 Criando docker-compose.yml básico para WordPress..."
    cp "$HARBOR_ROOT/templates/wordpress/docker-compose.yml" .
  fi

  # 2. Cria diretórios wp-content/plugins e wp-content/themes
  mkdir -p wp/wp-content/plugins wp/wp-content/themes

  # 3. Cria arquivo .env
  if [[ ! -f .env ]]; then
    echo "🌿 Criando arquivo .env..."
    cp "$HARBOR_ROOT/templates/wordpress/.env" .env
  fi

  # 4. Cria bin/harbor.sh
  if [[ ! -f bin/harbor.sh ]]; then
    echo "📄 Criando bin/harbor.sh..."
    mkdir -p bin
    cp "$HARBOR_ROOT/templates/wordpress/harbor.sh" bin/harbor.sh
    chmod +x bin/harbor.sh
  fi

  echo "✅ Estrutura inicial do WordPress criada com sucesso!"
  echo "Use ./bin/harbor.sh para subir containers e gerenciar o projeto."
}

up_wordpress() {
  echo "🚀 Subindo containers WordPress..."

  check_docker

  # Verifica docker-compose.yml
  if [[ ! -f docker-compose.yml ]]; then
    echo "❌ docker-compose.yml não encontrado."
    exit 1
  fi

  # Sobe os containers
  docker-compose up -d

  echo "✅ Containers WordPress iniciados com sucesso."
}

down_wordpress() {
  echo "🛑 Parando containers WordPress..."

  check_docker

  if [[ ! -f docker-compose.yml ]]; then
    echo "❌ docker-compose.yml não encontrado."
    exit 1
  fi

  docker-compose down

  echo "✅ Containers WordPress parados com sucesso."
}
