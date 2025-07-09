# 📈 stock-scraper 🚀

[![Made with Go](https://img.shields.io/badge/Made%20with-Go-blue)](https://golang.org)
[![Colly Web Scraper](https://img.shields.io/badge/Powered%20by-Colly-00bfff)](https://github.com/gocolly/colly)
[![Graph by Gonum](https://img.shields.io/badge/Graph%20with-Gonum-44cc11)](https://gonum.org)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

> ⚡ A savage Go-powered scraper that literally goes to Google Finance, grabs real-time stock data like a Wall Street ninja, and plots a damn beautiful bar graph.  
> 💾 Saves CSV too because spreadsheets = ✨corporate comfort✨.

---

## 🧠 What It Does

This Go project scrapes **live stock market data** from [Google Finance](https://www.google.com/finance) using the beast 🕷️ `Colly` and visualizes it using `gonum/plot`.

It:
- 🕸️ Visits stock pages on Google Finance
- 🧠 Extracts **company name**, **stock price**, and **price change**
- 📦 Dumps all data to `stocks.csv`
- 📊 Generates a bar graph `stocks_graph.png` with prices

---

## 🔍 Tech Stack (aka What Makes This Sexy)

| Purpose         | Package                         |
|----------------|----------------------------------|
| Scraping       | [`github.com/gocolly/colly`](https://github.com/gocolly/colly) |
| Graph plotting | [`gonum.org/v1/plot`](https://gonum.org/v1/plot) |
| CSV Writing    | Built-in `encoding/csv` |
| Concurrency    | `sync.Mutex` for safe Go-routines |
| Sanity & Logs  | `log`, `fmt`, and `os` |

---

## 📸 Screenshot

> When you actually understand the stock market through pixels.

![stocks_graph.png](stocks_graph.png)

---

## 🏗️ How to Run This Thing?

```bash
# Clone the beast
git clone https://github.com/anandbhalerao18/stock-scraper.git
cd stock-scraper

# Run the Go scraper
go run main.go
