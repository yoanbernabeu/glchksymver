# GitLab Check Symfony Version (glchksymver)

`glchksymver` is a command-line tool designed to fetch and display Symfony version information for projects within a specified GitLab group. It utilizes the GitLab API to retrieve projects and examines each project's `composer.json` file to determine the Symfony version used. This tool is particularly useful for teams managing multiple Symfony projects within GitLab, allowing for a quick overview of Symfony versions across projects.

## Features

- Fetches all projects within a specified GitLab group.
- Identifies the Symfony version for each project by analyzing its `composer.json` file.
- Displays a summary table of projects with their corresponding Symfony version.
- Lists projects that do not specify a Symfony version or do not contain a `composer.json` file, indicating potential non-Symfony projects.
- Utilizes ASCII art and progress bars for a friendly user interface.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go installed on your machine (version 1.16 or later recommended).
- A GitLab personal access token with permissions to access projects and repositories within your GitLab group.

## Installation

To install `glchksymver`, follow these steps:

1. Download the latest release from the [releases page](

## Building from Source

1. Clone the repository:
   ```bash
   git clone https://your-repository-url.git
   cd glchksymver
   ```

2. Build the binary:
   ```bash
    go build -o glchksymver
    ```

3. Move the binary to a directory in your `$PATH`:
    ```bash
    mv glchksymver /usr/local/bin
    ```
