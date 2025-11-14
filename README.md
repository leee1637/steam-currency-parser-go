```markdown
# 🎮 Steam Currency Parser Bot

**English** | [Русский](#-steam-currency-parser-bot-ru)

A Go-based Telegram bot that fetches and displays relative currency exchange rates used by Steam, calculated from the price of a single in-game item across different currencies.

## ✨ Features

- 💬 Command `/currency` — get current rates.
- 🕐 Automatic updates at **1:00 Moscow time**.
- 🧩 Support for **customizable currencies**.
- 📁 Easy setup with `.env` file.

## 📥 Installation

1. **Install Go** (version 1.25.1 or higher) from [https://go.dev/dl/](https://go.dev/dl).
2. Clone the repository:

   ```bash
   git clone https://github.com/leee1637/steam-currency-parser-go.git
   cd steam-currency-parser-go
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

## ⚙️ Setup

### 1. Get Telegram Bot Token

1. Open Telegram and search for `@BotFather`.
2. Start a chat and send `/newbot`.
3. Follow the instructions to create a bot and get the **token**.
4. Save the token for later.

### 2. Configure `.env` file

Create a file named `.env` in the root of the project:

```env
TELEGRAM_TOKEN=your_bot_token_here
APP_ID=730
MARKET_HASH_NAME=AK-47%20|%20Redline%20(Factory%20New)
```

- `TELEGRAM_TOKEN`: Your bot token from BotFather.
- `APP_ID`: Steam game ID. `730` is CS:GO.
- `MARKET_HASH_NAME`: Name of the item on Steam Market (URL-encoded).

### 3. How to Find `MARKET_HASH_NAME`

1. Go to [Steam Market](https://steamcommunity.com/market/).
2. Find an item (e.g. `AK-47 | Redline (Factory New)`).
3. Click on it.
4. Copy the URL. Example:
   ```
   https://steamcommunity.com/market/listings/730/AK-47%20|%20Redline%20(Factory%20New)
   ```
5. The part after `listings/730/` is your `MARKET_HASH_NAME`: `AK-47%20|%20Redline%20(Factory%20New)`.

### 4. Supported Currencies

The bot fetches rates based on a **fixed list of currencies** in `cmd/main.go`. You can **add or remove** them easily:

| Currency | Code | Symbol |
|----------|------|--------|
| USD      | 1    | $      |
| EUR      | 3    | €      |
| RUB      | 5    | ₽      |
| JPY      | 8    | ¥      |
| CNY      | 23   | ¥      |
| AED      | 32   | د.إ    |

**Currency codes** are based on [Steam's internal codes](https://partner.steamgames.com/doc/store/pricing/currencies).

## 🔧 How It Works

1. The bot fetches the price of a single item (e.g. `AK-47 | Redline`) in USD.
2. It then fetches the same item's price in other currencies.
3. It calculates the relative rate: `price_in_currency / price_in_USD`.
4. The result is sent to Telegram as: `1 USD = X currency`.
5. Automatically updates at 1:00 Moscow time.

## ▶️ Run

```bash
go run cmd/main.go
```

## 📷 Screenshots

(Add your screenshots here)

## 📌 Recommended Items for Stable Rates

For more stable currency rates, use items that change price less frequently:

- `CS:GO Case Key`
- `Glock-18 | Fade (Factory New)`
- `Souvenir AK-47`

Avoid volatile items like cases or skins with high demand.

---

# 🎮 Steam Currency Parser Bot (RU)

Telegram-бот, который показывает относительные курсы валют, используемые Steam, на основе цен на один и тот же предмет.

## ✨ Возможности

- 💬 Команда `/currency` — получить текущие курсы.
- 🕐 Автоматическое обновление в **1:00 по Московскому времени**.
- 🧩 Поддержка **настраиваемых валют**.
- 📁 Простая настройка через `.env`.

## 📥 Установка

1. **Установите Go** (версия 1.25.1 или выше) с [https://go.dev/dl/](https://go.dev/dl/).
2. Клонируйте репозиторий:

   ```bash
   git clone https://github.com/leee1637/steam-currency-parser-go.git
   cd steam-currency-parser-go
   ```

3. Установите зависимости:

   ```bash
   go mod tidy
   ```

## ⚙️ Настройка

### 1. Получите токен Telegram-бота

1. Откройте Telegram и найдите `@BotFather`.
2. Начните чат и отправьте `/newbot`.
3. Следуйте инструкциям, чтобы создать бота и получить **токен**.
4. Сохраните токен.

### 2. Настройте файл `.env`

Создайте файл `.env` в корне проекта:

```env
TELEGRAM_TOKEN=ваш_токен_бота
APP_ID=730
MARKET_HASH_NAME=AK-47%20|%20Redline%20(Factory%20New)
```

- `TELEGRAM_TOKEN`: Токен вашего бота из BotFather.
- `APP_ID`: ID игры в Steam. `730` — это CS:GO.
- `MARKET_HASH_NAME`: Название предмета на рынке Steam (в URL-кодировке).

### 3. Как найти `MARKET_HASH_NAME`

1. Перейдите на [Steam Market](https://steamcommunity.com/market/).
2. Найдите предмет (например, `AK-47 | Redline (Factory New)`).
3. Нажмите на него.
4. Скопируйте URL. Пример:
   ```
   https://steamcommunity.com/market/listings/730/AK-47%20|%20Redline%20(Factory%20New)
   ```
5. Часть после `listings/730/` — это `MARKET_HASH_NAME`: `AK-47%20|%20Redline%20(Factory%20New)`.

### 4. Поддерживаемые валюты

Бот получает курсы на основе **фиксированного списка валют** в `cmd/main.go`. Вы можете **добавить или удалить** их:

| Валюта | Код | Символ |
|--------|-----|--------|
| USD    | 1   | $      |
| EUR    | 3   | €      |
| RUB    | 5   | ₽      |
| JPY    | 8   | ¥      |
| CNY    | 23  | ¥      |
| AED    | 32  | د.إ    |

**Коды валют** основаны на [внутренних кодах Steam](https://partner.steamgames.com/doc/store/pricing/currencies).

## 🔧 Как это работает

1. Бот получает цену на один предмет (например, `AK-47 | Redline`) в USD.
2. Затем он получает цену на тот же предмет в других валютах.
3. Вычисляет относительный курс: `цена_в_валюте / цена_в_USD`.
4. Результат отправляется в Telegram как: `1 USD = X валюта`.
5. Автоматическое обновление в 1:00 по Московскому времени.

## ▶️ Запуск

```bash
go run cmd/main.go
```

## 📷 Скриншоты

(Добавьте сюда свои скриншоты)

## 📌 Рекомендуемые предметы для стабильных курсов


Для более стабильных курсов используйте предметы, цена на которые меняется реже:

- `CS:GO Case Key`
- `Glock-18 | Fade (Factory New)`
- `Souvenir AK-47`

Избегайте волатильных предметов, таких как кейсы или скины с высоким спросом.
```
