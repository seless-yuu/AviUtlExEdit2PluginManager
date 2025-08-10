# Coding Agent Document

## Project Overview

This repository contains a plugin manager for AviUtl, built using Wails (Go + Svelte).

## Core Philosophy

This document is for coding agents working on this repository. Please adhere to the following guidelines. If you gain new knowledge or establish new procedures during your work, please update the documentation in the `docs/` directory to help future agents.

### Coding Guidelines

-   **UNIX Philosophy**: Strive to follow the UNIX philosophy. Keep functions and packages small, focused on a single responsibility, and reusable.
-   **Code Separation**: Treat backend code as a library as much as possible, cleanly separating it from the application layer (the Wails `app.go` file).
-   **Clarity**: Comments should explain the "why" behind the code, not the "what". The code itself should be as self-evident as possible.

### Testing Guidelines

-   **TDD Approach**: When implementing new features or fixing bugs, please implement tests to facilitate iterative development and prevent regressions. The goal is to build a robust and reliable application.

## Detailed Documentation

For more specific guidelines and project information, please refer to the documents in the `docs/` directory.

-   **[Getting Started](./docs/GETTING_STARTED.md)**: For environment setup, key development commands, and testing procedures.
-   **[Branching Model](./docs/BRANCHING_MODEL.md)**: For the `git-flow` based branching strategy and naming conventions.
