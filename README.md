# 📈 Real-Time Stock Scraper & Graph Generator 🚀

Yo! Welcome to **Wall Street Vibes for Devs** – a Go-powered web scraping project that:
- Scrapes real-time stock data 📊 from **Google Finance**
- Saves it to a CSV 📁
- Generates a super-clean bar chart 🖼️ of stock prices using **gonum/plot**

> Just Go do it. Literally.

---

## ⚡ Features

🔥 **What it does like a boss**:

- ✅ Scrapes stock price, company name, and price change
- ✅ Outputs to `stocks.csv` in a readable tabular format
- ✅ Generates a sexy bar chart as `stocks_graph.png`
- ✅ Clean error logging & concurrent-safe scraping (mutex magic)
- ✅ Extensible for more tickers anytime!

---

## 🔧 Tech Stack

- [Go (Golang)](https://go.dev/) – the beast
- [Colly](https://github.com/gocolly/colly) – scraping made elegant
- [gonum/plot](https://github.com/gonum/plot) – for stunning visuals
- 🧠 `sync.Mutex` – because concurrency isn't for the weak

---

## 📦 How to Run

### 🚨 Prerequisites
- Go installed (`>=1.18`)
- Internet connection (duh 😎)
- Run inside terminal or VS Code or wherever you code like a hacker

### 🏃 Steps

```bash
git clone https://github.com/your-username/your-stock-scraper.git
cd your-stock-scraper
go run main.go

