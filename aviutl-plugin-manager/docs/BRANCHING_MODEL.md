# Branching Model

This project follows the **Git Flow** branching model to ensure a structured and predictable development process.

## Main Branches

The repository holds two main branches with an infinite lifetime:

-   `main`: This branch stores the official, stable release history. It should only be merged into from `develop` or `hotfix` branches.
-   `develop`: This is the main integration branch for new features. All feature branches are created from `develop` and merged back into it.

## Supporting Branches

Supporting branches are used to aid parallel development between team members, ease tracking of features, and assist in preparing for releases. Unlike the main branches, these branches always have a limited lifetime, since they will be removed eventually after the work is done.

The different types of branches we may use are:

### `feature/*`

-   **Purpose:** To develop new features.
-   **Branch from:** `develop`
-   **Merge back to:** `develop`
-   **Naming convention:** `feature/<short-description>` (e.g., `feature/add-plugin-sorting`)

### `fix/*`

-   **Purpose:** To fix non-critical bugs in the `develop` branch.
-   **Branch from:** `develop`
-   **Merge back to:** `develop`
-   **Naming convention:** `fix/<issue-description>` (e.g., `fix/profile-save-error`)

### `docs/*`

-   **Purpose:** For adding, updating, or correcting documentation.
-   **Branch from:** `develop`
-   **Merge back to:** `develop`
-   **Naming convention:** `docs/<document-name>` (e.g., `docs/update-readme`)

### `release/*`

-   **Purpose:** To prepare for a new production release. This branch allows for last-minute fixes and preparation.
-   **Branch from:** `develop`
-   **Merge back to:** `develop` and `main`
-   **Naming convention:** `release/vX.Y.Z` (e.g., `release/v0.2.0`)

### `hotfix/*`

-   **Purpose:** To patch a critical bug in a production version.
-   **Branch from:** `main`
-   **Merge back to:** `develop` and `main`
-   **Naming convention:** `hotfix/<issue-description>`
