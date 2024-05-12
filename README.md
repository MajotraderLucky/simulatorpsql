# psql-gRPC Simulator
  

## Overview

The psql-gRPC Simulator is a software tool designed for synchronizing data from CSV files into a PostgreSQL database and providing an interface to interact with the data via gRPC. The application handles the conversion of `.csv` data into entries in a PostgreSQL database, which are then accessible through a gRPC interface.

  

## Architecture

The system architecture includes several key components:

- **CSV Data Input**: The source data in `.csv` format.

- **Data Converter**: Converts data from `.csv` files into PostgreSQL format.

- **PostgreSQL Database (PgDB)**: Stores the converted data.

- **gRPC Interface**: Facilitates data retrieval and interaction through gRPC services.

- **Sync Service**: Ensures consistent and up-to-date data across the system.

  

## Features

- **Data Conversion**: Automated conversion from `.csv` to PostgreSQL.

- **Data Retrieval**: Facilitates data access through a gRPC interface.

- **Data Synchronization**: Maintains data consistency across components.

  

## Setup and Installation

Detailed instructions on setting up and running the simulator in your local environment.

  

### Prerequisites

- PostgreSQL

- Go 1.21.4

  

### Installation

1. Clone the repository:

```bash
git clone https://github.com/MajotraderLucky/simulatorpsql.git

```


[Readme на русском языке (Russian)]([docs/ru/README.md](https://github.com/MajotraderLucky/simulatorpsql/blob/main/docs/ru/README.md))
