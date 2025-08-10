# Getting Started

This document outlines the key commands for setting up and working on this project.

## Environment Setup

The frontend dependencies must be installed via `npm`. The Wails build process for the `main` package depends on the frontend being built first.

1.  **Navigate to the frontend directory:**
    ```bash
    cd frontend
    ```

2.  **Install dependencies:**
    ```bash
    npm install
    ```
    This will install `vite` and other required packages into `node_modules`.

## Running in Development Mode

To run the application with hot-reloading for both the frontend and backend:

1.  **Navigate to the project root (`aviutl-plugin-manager`).**
2.  **Run the dev command:**
    ```bash
    wails dev
    ```

## Running Tests

The backend has a suite of unit tests. To run all tests for all packages:

1.  **Navigate to the project root (`aviutl-plugin-manager`).**
2.  **(Optional) Build the frontend first:** The test for the `main` package requires the `frontend/dist` directory to exist.
    ```bash
    cd frontend && npm run build && cd ..
    ```
3.  **Run the Go tests:**
    ```bash
    go test ./...
    ```
