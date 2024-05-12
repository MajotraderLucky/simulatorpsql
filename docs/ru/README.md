# Симулятор psql-gRPC

## Обзор
Симулятор psql-gRPC — это программный инструмент, разработанный для синхронизации данных из CSV-файлов в базу данных PostgreSQL и предоставления интерфейса для взаимодействия с данными через gRPC. Приложение обрабатывает конвертацию данных из `.csv` в записи базы данных PostgreSQL, которые затем доступны через интерфейс gRPC.

## Архитектура
Архитектура системы включает в себя несколько ключевых компонентов:
- **Ввод данных CSV**: Исходные данные в формате `.csv`.
- **Конвертер данных**: Преобразует данные из файлов `.csv` в формат PostgreSQL.
- **База данных PostgreSQL (PgDB)**: Хранит конвертированные данные.
- **gRPC Интерфейс**: Обеспечивает извлечение данных и взаимодействие через gRPC-сервисы.
- **Сервис синхронизации**: Гарантирует консистентность и актуальность данных в системе.

## Функциональность
- **Конвертация данных**: Автоматическое преобразование из `.csv` в PostgreSQL.
- **Извлечение данных**: Облегчает доступ к данным через интерфейс gRPC.
- **Синхронизация данных**: Поддерживает консистентность данных между компонентами.

## Настройка и установка
Подробные инструкции по настройке и запуску симулятора в вашем локальном окружении.

### Предварительные требования
- PostgreSQL
- Go 1.21.4

### Установка
1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/MajotraderLucky/simulatorpsql.git
   