# 🎮 Steam Currency Parser Bot / Бот для анализа валют Steam

<div align="center">

<img width="343" height="152" alt="image" src="https://github.com/user-attachments/assets/957a8229-72fc-4f8a-a93a-c3612820c763" />

**🤖 Smart Telegram Bot | 🔄 Real-time Steam Rates | 🎯 Go Language**

*A powerful Telegram bot written in Go that parses Steam's internal currency exchange rates based on in-game item prices across different regions*

🇺🇸 [English](#-features) | 🇷🇺 [Русский](#-возможности)

</div>

## 🎯 About This Bot

This **Go-based Telegram bot** provides real-time monitoring of Steam's internal currency exchange rates. Unlike traditional financial APIs, it calculates rates directly from Steam Marketplace data by comparing prices of the same in-game item across different currency regions.

### 🔍 How It Works:
- **🎮 In-Game Item Based**: Uses actual Steam market prices of CS:GO/CS2 items
- **💱 Real Exchange Rates**: Calculates actual rates Steam uses for conversions  
- **🔄 Automated Updates**: Fresh data daily at 1:00 Moscow time
- **⚡ Go Performance**: Built with Go for high performance and reliability

### 💡 Perfect For:
- Steam traders and market analysts
- Gamers purchasing items from different regions
- Developers needing Steam currency data
- Financial researchers studying virtual economies


## ✨ Features {#-features}

<div align="center">

| Feature | Description |
|---------|-------------|
| 💬 **Command** | `/currency` - Get current exchange rates |
| 🕐 **Auto Updates** | Daily at **1:00 Moscow time** |
| 🧩 **Customizable** | Support for multiple currencies |
| 📁 **Easy Setup** | Simple `.env` configuration |

</div>

## 🚀 Quick Start

### 1. Prerequisites
- **Go 1.25.1+** - [Download here](https://go.dev/dl/)

### 2. Installation
```bash
git clone https://github.com/leee1637/steam-currency-parser-go.git
cd steam-currency-parser-go
go mod tidy
```

### 3. Bot Setup

#### Get Telegram Bot Token
1. Open Telegram, search for `@BotFather`
2. Send `/newbot` and follow instructions
3. Copy your bot token

#### Configure Environment
Create `.env` file:
```env
TELEGRAM_TOKEN=your_bot_token_here
APP_ID=730
MARKET_HASH_NAME=AK-47%20|%20Redline%20(Factory%20New)
```

#### Find Market Hash Name
1. Go to [Steam Market](https://steamcommunity.com/market/)
2. Find your item (e.g., `AK-47 | Redline (Factory New)`)
3. Copy URL segment after `listings/730/`

**Example URL:**
```
https://steamcommunity.com/market/listings/730/AK-47%20|%20Redline%20(Factory%20New)
```
**MARKET_HASH_NAME:** `AK-47%20|%20Redline%20(Factory%20New)`

## 💰 Supported Currencies

<div align="center">

| Currency | Code | Symbol |
|----------|------|--------|
| USD | 1 | $ |
| EUR | 3 | € |
| RUB | 5 | ₽ |
| JPY | 8 | ¥ |
| CNY | 23 | ¥ |
| AED | 32 | د.إ |

</div>

*Currency codes based on [Steam's documentation](https://partner.steamgames.com/doc/store/pricing/currencies)*

## 🛠 How It Works

<div align="center">


</div>

1. **Fetch USD Price** - Gets item price in US Dollars
2. **Get Other Currencies** - Retrieves same item's price in other currencies
3. **Calculate Rates** - Computes `price_in_currency / price_in_USD`
4. **Display Results** - Sends formatted rates to Telegram
5. **Auto Update** - Daily refresh at 1:00 Moscow time

## 🎯 Recommended Items

For stable currency rates, use these items:

- ✅ `CS:GO Case Key`
- ✅ `Glock-18 | Fade (Factory New)` 
- ✅ `Souvenir AK-47`

**Avoid:** Volatile items like cases or high-demand skins

## 📸 Screenshots

<div align="center">

### Bot Interface
<img width="344" height="191" alt="image" src="https://github.com/user-attachments/assets/1a112099-53d0-441c-a7c3-d0f084571ab6" />

### Rate Display
<img width="681" height="169" alt="image" src="https://github.com/user-attachments/assets/d0f72bc6-3706-41b0-8d6a-50fd72e17bb5" />

</div>

## ▶️ Run the Bot

```bash
go run cmd/main.go
```

---

# 🎮 Steam Currency Parser Bot (RU) {#-steam-currency-parser-bot-ru}

<div align="center">

<img width="343" height="152" alt="image" src="https://github.com/user-attachments/assets/b5c5811a-0b4d-42c7-9cfb-d1c82f25bb66" />

🤖 Умный Telegram-бот | 🔄 Курсы Steam в реальном времени | 🎯 Язык Go

Мощный Telegram-бот на языке Go, который парсит внутренние курсы валют Steam на основе цен внутриигровых предметов в разных регионах

[English](#-features) | [Русский](#-возможности)

</div>

🎯 О боте
Этот Telegram-бот на языке Go предоставляет мониторинг внутренних курсов валют Steam в реальном времени. В отличие от традиционных финансовых API, он рассчитывает курсы напрямую из данных Steam Marketplace, сравнивая цены на один и тот же внутриигровой предмет в разных валютных регионах.

🔍 Как это работает:
🎮 На основе игровых предметов: Использует реальные цены предметов CS:GO/CS2 из маркета Steam

💱 Реальные курсы обмена: Рассчитывает фактические курсы, которые использует Steam для конвертаций

🔄 Автоматическое обновление: Свежие данные ежедневно в 1:00 по Москве

⚡ Производительность Go: Написан на Go для высокой производительности и надежности

💡 Идеально подходит для:
Трейдеров Steam и аналитиков рынка

Геймеров, покупающих предметы из разных регионов

Разработчиков, нуждающихся в данных о валютах Steam

Исследователей, изучающих виртуальные экономики

## ✨ Возможности {#-возможности}

<div align="center">

| Функция | Описание |
|---------|----------|
| 💬 **Команды** | `/currency` - Получить текущие курсы |
| 🕐 **Автообновление** | Ежедневно в **1:00 по Москве** |
| 🧩 **Настройка** | Поддержка множества валют |
| 📁 **Простая установка** | Простая настройка через `.env` |

</div>

## 🚀 Быстрый старт

### 1. Требования
- **Go 1.25.1+** - [Скачать здесь](https://go.dev/dl/)

### 2. Установка
```bash
git clone https://github.com/leee1637/steam-currency-parser-go.git
cd steam-currency-parser-go
go mod tidy
```

### 3. Настройка бота

#### Получение токена бота
1. Откройте Telegram, найдите `@BotFather`
2. Отправьте `/newbot` и следуйте инструкциям
3. Скопируйте токен бота

#### Настройка окружения
Создайте файл `.env`:
```env
TELEGRAM_TOKEN=ваш_токен_бота
APP_ID=730
MARKET_HASH_NAME=AK-47%20|%20Redline%20(Factory%20New)
```

#### Поиск Market Hash Name
1. Перейдите на [Steam Market](https://steamcommunity.com/market/)
2. Найдите предмет (например, `AK-47 | Redline (Factory New)`)
3. Скопируйте часть URL после `listings/730/`

**Пример URL:**
```
https://steamcommunity.com/market/listings/730/AK-47%20|%20Redline%20(Factory%20New)
```
**MARKET_HASH_NAME:** `AK-47%20|%20Redline%20(Factory%20New)`

## 💰 Поддерживаемые валюты

<div align="center">

| Валюта | Код | Символ |
|--------|-----|--------|
| USD | 1 | $ |
| EUR | 3 | € |
| RUB | 5 | ₽ |
| JPY | 8 | ¥ |
| CNY | 23 | ¥ |
| AED | 32 | د.إ |

</div>

*Коды валют на основе [документации Steam](https://partner.steamgames.com/doc/store/pricing/currencies)*

## 🛠 Как это работает

<div align="center">


</div>

1. **Получение цены в USD** - Получает цену предмета в долларах
2. **Получение других валют** - Находит цену того же предмета в других валютах
3. **Расчет курсов** - Вычисляет `цена_в_валюте / цена_в_USD`
4. **Отображение результатов** - Отправляет форматированные курсы в Telegram
5. **Автообновление** - Ежедневное обновление в 1:00 по Москве

## 🎯 Рекомендуемые предметы

Для стабильных курсов используйте:

- ✅ `CS:GO Case Key`
- ✅ `Glock-18 | Fade (Factory New)`
- ✅ `Souvenir AK-47`

**Избегайте:** Волатильные предметы like кейсы или скины с высоким спросом

## 📸 Скриншоты

<div align="center">

### Интерфейс бота
<img width="344" height="191" alt="image" src="https://github.com/user-attachments/assets/6c09e4a8-2cba-4fa3-adbb-de8dfb83bbfc" />

### Отображение курсов
<img width="681" height="169" alt="image" src="https://github.com/user-attachments/assets/32e24529-f6c1-4767-8b30-d9f4d77b46b2" />

</div>

## ▶️ Запуск бота

```bash
go run cmd/main.go
```

---

<div align="center">

**⭐ Если вам понравился проект, не забудьте поставить звезду!**

</div>
