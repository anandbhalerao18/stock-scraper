package main

import (
	"encoding/csv"         // For creating and writing CSV files
	"fmt"                  // For formatted I/O
	"log"                  // For logging errors and messages
	"os"                   // For file operations
	"sort"                 // For sorting operations
	"strconv"              // For converting strings to numeric types
	"strings"              // For string manipulation
	"sync"                 // For handling concurrent access

	"github.com/gocolly/colly"      // For web scraping
	"gonum.org/v1/plot"             // For creating plots
	"gonum.org/v1/plot/plotter"     // For creating various plot elements
	"gonum.org/v1/plot/plotutil"    // For utility functions for plots
	"gonum.org/v1/plot/vg"          // For specifying sizes and units for plots
)

// Stock represents a structure to store stock data
type Stock struct {
	company string // Name of the company
	price   string // Stock price
	change  string // Price change
}

func main() {
	// List of stock tickers to scrape data for
	tickers := []string{
		"MSFT:NASDAQ", "AAPL:NASDAQ", "GOOGL:NASDAQ", "TSLA:NASDAQ", "AMZN:NASDAQ",
		"NVDA:NASDAQ",
	}

	var stocks []Stock       // Slice to store stock data
	var mu sync.Mutex        // Mutex to handle concurrent writes to the stocks slice

	// Create a new Colly collector for web scraping
	c := colly.NewCollector()

	// Log the URL being visited (for debugging purposes)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})

	// Handle errors during web scraping
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error visiting %s: %v\n", r.Request.URL, err)
	})

	// Scrape stock data from the HTML body of the page
	c.OnHTML("body", func(e *colly.HTMLElement) {
		// Extract relevant stock data using CSS selectors
		stock := Stock{
			company: e.ChildText("div.zzDege"),                    // Company name
			price:   cleanPrice(e.ChildText("div.YMlKec.fxKbKc")), // Cleaned stock price
			change:  e.ChildText("div.P6K39c"),                    // Price change
		}

		// Validate if the extracted data is complete
		if stock.company != "" && stock.price != "" && stock.change != "" {
			mu.Lock() // Lock to safely append to the stocks slice
			stocks = append(stocks, stock)
			mu.Unlock()
			// Log fetched stock data
			fmt.Printf("Fetched - Company: %s, Price: %s, Change: %s\n", stock.company, stock.price, stock.change)
		} else {
			log.Println("Failed to fetch valid data for a stock. Check the selectors.")
		}
	})

	// Visit the URL for each ticker to scrape data
	for _, ticker := range tickers {
		c.Visit("https://www.google.com/finance/quote/" + ticker)
	}

	// Wait for all scraping processes to finish
	c.Wait()

	// Create a CSV file to store the scraped data
	file, err := os.Create("stocks.csv")
	if err != nil {
		log.Fatalf("Failed to create output CSV file: %v\n", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file) // CSV writer
	defer writer.Flush()

	// Write the CSV header row
	err = writer.Write([]string{"Company Name", "Price", "Change"})
	if err != nil {
		log.Fatalf("Failed to write CSV header: %v\n", err)
	}

	// Write each stock record to the CSV file
	for _, stock := range stocks {
		record := []string{stock.company, stock.price, stock.change}
		err = writer.Write(record)
		if err != nil {
			log.Printf("Error writing stock data to CSV: %v\n", err)
		}
	}

	fmt.Println("Data has been successfully written to stocks.csv")

	// Generate a graph of the stock prices if data is available
	if len(stocks) > 0 {
		generateGraph(stocks)
	}
}

// Cleans the stock price string and extracts a valid numeric value
func cleanPrice(price string) string {
	// Split the string by "$" and find the first valid numeric part
	parts := strings.Split(price, "$")
	for _, part := range parts {
		part = strings.TrimSpace(part) // Remove leading/trailing spaces
		if _, err := strconv.ParseFloat(part, 64); err == nil {
			return part
		}
	}
	return ""
}

// Generates a bar chart of stock prices and saves it as an image
func generateGraph(stocks []Stock) {
	// Sort the stocks by price in ascending order
	sort.Slice(stocks, func(i, j int) bool {
		price1, _ := strconv.ParseFloat(stocks[i].price, 64)
		price2, _ := strconv.ParseFloat(stocks[j].price, 64)
		return price1 < price2
	})

	p := plot.New() // Create a new plot
	if p == nil {
		log.Fatal("Failed to create plot")
	}

	p.Title.Text = "Stock Prices"       // Set the title of the plot
	p.X.Label.Text = "Companies"        // Set X-axis label
	p.Y.Label.Text = "Price (USD)"      // Set Y-axis label
	p.Add(plotter.NewGrid())            // Add gridlines to the plot

	// Convert stock prices to a format suitable for the bar chart
	values := make(plotter.Values, len(stocks))
	labels := make([]string, len(stocks)) // Company names for the X-axis
	points := make([]plotter.XY, len(stocks)) // Points for adding text labels

	for i, stock := range stocks {
		priceStr := stock.price
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Printf("Error parsing price for %s: %v\n", stock.company, err)
			continue
		}

		values[i] = price   // Add price value
		labels[i] = stock.company // Add company name
		points[i].X = float64(i)  // X-coordinate for label
		points[i].Y = price       // Y-coordinate for label
	}

	// Ensure valid data is present for plotting
	if len(values) == 0 {
		log.Fatal("No valid data for the graph")
	}

	// Create the bar chart
	barChart, err := plotter.NewBarChart(values, vg.Points(20))
	if err != nil {
		log.Fatalf("Failed to create bar chart: %v\n", err)
	}

	// Customize the appearance of the bar chart
	barChart.Color = plotutil.Color(2)         // Set bar color
	barChart.LineStyle.Width = vg.Points(0.5) // Set line width for bars
	p.Add(barChart)                           // Add bar chart to plot
	p.NominalX(labels...)                     // Set labels for X-axis

	// Add text labels above bars
	labelPlot, err := plotter.NewLabels(plotter.XYLabels{
		XYs:    points,
		Labels: formatPriceLabels(values), // Format prices as labels
	})
	if err != nil {
		log.Fatalf("Failed to create bar labels: %v\n", err)
	}
	p.Add(labelPlot)

	// Save the plot as a PNG image
	err = p.Save(12*vg.Inch, 6*vg.Inch, "stocks_graph.png")
	if err != nil {
		log.Fatalf("Failed to save graph: %v\n", err)
	}

	fmt.Println("Graph has been successfully saved as stocks_graph.png")
}

// Formats price values into a string format for bar chart labels
func formatPriceLabels(values plotter.Values) []string {
	labels := make([]string, len(values))
	for i, value := range values {
		labels[i] = fmt.Sprintf("$%.2f", value) // Format as $XX.XX
	}
	return labels
}
