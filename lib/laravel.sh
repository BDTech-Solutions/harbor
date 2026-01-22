init_laravel() {
  if [ "$(ls -A .)" ]; then
    echo "O diretório não está vazio."
    exit 1
  fi

  echo "Criando projeto Laravel via Docker (Composer)..."

  docker run --rm \
    -u "$(id -u):$(id -g)" \
    -v "$(pwd):/app" \
    -w /app \
    composer:2 \
    composer create-project laravel/laravel .

  echo "Instalando Laravel Sail..."

  docker run --rm \
    -u "$(id -u):$(id -g)" \
    -v "$(pwd):/app" \
    -w /app \
    laravelsail/php83-composer:latest \
    php artisan sail:install

  echo "Projeto Laravel criado com Sail (não iniciado)."

  HAROBR_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

TEMPLATE="$HAROBR_ROOT/templates/laravel/harbor.sh"
TARGET="./harbor.sh"

cp "$TEMPLATE" "$TARGET"
chmod +x "$TARGET"
}

bootstrap_laravel() {
  echo "🔎 Verificando projeto Laravel..."

  # 1. Verificar se é um projeto Laravel
  if [[ ! -f artisan || ! -f composer.json || ! -f bootstrap/app.php ]]; then
    echo "❌ Este diretório não parece ser um projeto Laravel."
    exit 1
  fi

  # 2. Verificar se harbor.sh já existe
  if [[ -f harbor.sh ]]; then
    echo "ℹ️  harbor.sh já existe neste projeto. Nada a fazer."
    exit 0
  fi

  # 3. Verificar se o template existe
  local template="$HARBOR_ROOT/templates/laravel/harbor.sh"

  if [[ ! -f "$template" ]]; then
    echo "❌ Template não encontrado em:"
    echo "   $template"
    exit 1
  fi

  # 4. Copiar o template para a raiz do projeto
  cp "$template" ./harbor.sh

  # 5. Tornar o arquivo executável
  chmod +x ./harbor.sh

  echo "✅ harbor.sh criado com sucesso."
  echo "👉 Execute ./harbor.sh para iniciar o ambiente Laravel."
}
