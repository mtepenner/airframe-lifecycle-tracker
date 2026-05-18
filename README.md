# ✈️ Airframe Lifecycle Tracker

## 📖 Description
Airframe Lifecycle Tracker is a high-performance analytics platform and visual dashboard designed to manage and track the immutable facts and historical lifecycle data of commercial airframes, including the MD-11, Boeing 737, and DC-8. Built with a robust Go backend and a modern React/TypeScript frontend, the system handles heavy multi-dimensional data, downtime tracking, and complex state changes over time.

## 📑 Table of Contents
- [Features](#-features)
- [Technologies Used](#-technologies-used)
- [Installation](#-installation)
- [Usage](#-usage)
- [Contributing](#-contributing)
- [License](#-license)

## 🚀 Features
* **Time-Travel Analytics**: Executes time-travel queries for historical livery and seating using SCD Type 2 configurations.
* **High-Performance Pipeline**: Spawns Goroutines in the Go analytics engine to concurrently fetch multi-dimensional data.
* **Advanced Visualizations**: Features a React dashboard with an interactive Timeline Scrubber for historical changes, an expandable Downtime Data Grid, and a Maintenance Heatmap for D-Check frequencies.
* **Heavy Data Pre-computation**: Utilizes PostgreSQL materialized views to pre-compute heavy CUBE matrices for fast analytical querying.
* **Complex Global Filtering**: Implements Zustand to manage complex global filters for Year, Model, and Operator data.

## 🛠️ Technologies Used
**Backend & Analytics Engine:**
* **Go**: High-performance data pipeline and minimal binary generation.
* **Chi / Gorilla Mux**: Router logic for API endpoints.
* **PostgreSQL & pgx**: Persistent data storage and database connection pooling.
* **golang-migrate**: Strict version control for database schemas and state changes.

**Frontend (Visual Presentation Layer):**
* **React & TypeScript**: Interactive user interface components.
* **Zustand**: Complex state and filter management.
* **React Query / SWR**: Aggressive data caching for matrix responses.
* **CSS-in-JS**: Minimalist blue/purple color palette theming.

**Infrastructure & DevOps:**
* **Docker & Docker Compose**: Multi-stage builds and local development orchestration.
* **Nginx**: Ingress routing and static asset serving.
* **GitHub Actions**: Automated CI/CD pipelines for testing Go APIs and migration verification.

## 💻 Installation
1. Clone the repository to your local machine.
2. Ensure Docker and Docker Compose are installed on your system.
3. Use the provided Makefile to spin up the local development environment:
   ```bash
   # Spins up Postgres, the Go Engine, and the React Dashboard
   docker-compose up

   ```


4. Run the database migrations to initialize the schema:

```bash
# Executes migrate up commands
make migrate-up

```



## 💡 Usage

Once the containers are running, the Nginx ingress will route traffic appropriately between the frontend and backend.

* **Frontend**: Access the interactive dashboard to view the timeline scrubbers and heatmaps.
* **API Endpoints**: The Go backend exposes endpoints such as `/api/v1/analytics/downtime` and `/api/v1/fleet` to deliver strict JSON payloads defining the analytical data.

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the issues page.

## 📄 License

This project is licensed under the MIT License - Copyright (c) 2026 Matthew Penner.
The software is provided "as is", without warranty of any kind, express or implied. In no event shall the authors or copyright holders be liable for any claim, damages, or other liability. See the `LICENSE` file for full details.
