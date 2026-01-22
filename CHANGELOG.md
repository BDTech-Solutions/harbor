# Changelog

Todas as mudanças notáveis neste projeto serão documentadas neste arquivo.

O formato segue [Keep a Changelog](https://keepachangelog.com/).

---

## [Unreleased]
- Preparar suporte para novas stacks: Python, Go, PHP puro
- Adicionar comandos adicionais para Laravel e WordPress
- Melhorias na experiência do usuário e mensagens do CLI

## [0.2.0] - 2026-01-21
### Added
- Suporte inicial para WordPress
- Comando `harbor wordpress init` para bootstrap completo
- Estrutura de pastas e containers Docker pronta para WordPress

## [0.1.0] - 2026-01-20
### Added
- Bootstrap de projetos Laravel via `harbor laravel init` e `harbor laravel bootstrap`
- Criação automática do `harbor.sh` para projetos Laravel clonados
- Integração com Docker para Composer e Sail
- Mensagens de verificação de ambiente (Docker ativo, Sail instalado)
