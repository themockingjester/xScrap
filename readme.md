# xScrap 🕷️

xScrap is a high-performance scraping service built with Golang. It is designed to efficiently extract data while dynamically optimizing system resources to prevent crashes or hangs.

### 🚀 Features

Adaptive Resource Handling – Automatically adjusts scraping windows based on available RAM, ensuring smooth operation.

Scrape by ID – Extract elements directly using their HTML id.

Scrape by CSS Selector – Retrieve data with flexible and powerful CSS selectors.

Scrape by XPath – Perform advanced scraping with precise XPath queries.

Fast & Efficient – Optimized for speed and concurrency with Go’s goroutines.

Robust Handling – Better manages edge cases to keep scraping stable.

xScrap can be tuned as per your needs by updating values in the configuration file. You can adjust settings such as concurrency, scraping windows, retries, and other runtime behaviors to optimize performance and resource usage.

### 📦 Getting Started

Follow these steps to run the app locally:

1. Clone the Repository
git clone https://github.com/themockingjester/xScrap.git
then go to project using (cd xScrap)

2. Install Dependencies
(using go get command ).

4. Run the Application
go run ./cmd/server/server.go

### 🛠️ Tech Stack

Language: Golang

Core: net/http, chromedp

### Demo
Check out a demo where xScrap is shown handling both normal and heavy load scenarios (simulated/mimicked): ![xScrap Demo](https://drive.google.com/file/d/1Z_4di1q2wjE4MkkcHnFLPNms000ce1xa/view?usp=sharing)
