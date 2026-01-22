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
    local UP_FLAG=false

    # Verifica se --up foi passado
    for arg in "$@"; do
        [[ "$arg" == "--up" ]] && UP_FLAG=true
    done

    echo "⚓ Harbor - Inicializando projeto WordPress"

    # 1. Cria docker-compose.yml se não existir
    if [[ ! -f docker-compose.yml ]]; then
        echo "📦 Criando docker-compose.yml básico para WordPress..."
        cp "$HARBOR_ROOT/templates/wordpress/docker-compose.yml" .
    else
        echo "✅ docker-compose.yml já existe"
    fi

    # 2. Cria diretórios
    mkdir -p wp/wp-content/plugins wp/wp-content/themes wp/wp-content/uploads

    # 3. Cria .env se não existir
    if [[ ! -f .env ]]; then
        echo "🌿 Criando arquivo .env..."
        cp "$HARBOR_ROOT/templates/wordpress/.env" .env
    else
        echo "✅ .env já existe, mantendo valores atuais."
    fi

    # 4. Cria harbor.sh na raiz do projeto se não existir
    if [[ ! -f bin/harbor.sh ]]; then
        echo "📄 Criando bin/harbor.sh..."
        mkdir -p bin
        cp "$HARBOR_ROOT/templates/wordpress/harbor.sh" bin/harbor.sh
        chmod +x bin/harbor.sh
    fi

    echo "✅ Estrutura inicial do WordPress criada com sucesso!"

    # 5. Se --up foi passado, sobe os containers automaticamente
    if [[ "$UP_FLAG" == true ]]; then
        up_wordpress
    else
        # Pergunta interativa caso --up não seja usado
        read -p "Deseja subir os containers agora? [y/N]: " RESP
        [[ "$RESP" =~ ^[Yy]$ ]] && up_wordpress
    fi

    echo "👉 Use ./bin/harbor.sh para gerenciar o projeto WordPress."
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
  docker compose -f docker-compose.yml up -d

  echo "✅ Containers WordPress iniciados com sucesso."
}

down_wordpress() {
  echo "🛑 Parando containers WordPress..."

  check_docker

  if [[ ! -f docker-compose.yml ]]; then
    echo "❌ docker-compose.yml não encontrado."
    exit 1
  fi

  docker compose down

  echo "✅ Containers WordPress parados com sucesso."
}
