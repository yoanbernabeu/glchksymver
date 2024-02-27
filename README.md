# GitLab Check Symfony Version (glchksymver)

`glchksymver` is a command-line tool designed to fetch and display Symfony version information for projects within a specified GitLab group. It utilizes the GitLab API to retrieve projects and examines each project's `composer.json` file to determine the Symfony version used. This tool is particularly useful for teams managing multiple Symfony projects within GitLab, allowing for a quick overview of Symfony versions across projects.

## Features

- Fetches all projects within a specified GitLab group.
- Identifies the Symfony version for each project by analyzing its `composer.json` file (Check the version of `symfony/framework-bundle`).
- Displays a summary table of projects with their corresponding Symfony version.
- Lists projects that do not specify a Symfony version or do not contain a `composer.json` file, indicating potential non-Symfony projects.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- A GitLab personal access token with permissions to access projects and repositories within your GitLab group.
- A GitLab group ID, which can be found in the `Settings` > `General` section of your GitLab group.

## How to use

### Usage

`glchksymver` is **VERY easy to use**. Just run the following command:

```bash
glchksymver -gitlab_url="https://gitlab.com" -group_id=123456789 -token="glpat-xxxxxxxxxxxxxx"
```

### Installation

#### From binary

* Linux/Darwin

Just run the following command:

```bash
curl -sL https://raw.githubusercontent.com/yoanbernabeu/glchksymver/main/install.sh | bash
```

* Other Operating Systems

Please download the binary from the [release page](https://github.com/yoanbernabeu/glchksymver/releases) and move it to your PATH.

#### From source

glchksymver is written in Go, so you need to install it first.

```bash
git clone git@github.com:yoanbernabeu/glchksymver.git
cd glchksymver
go build -o glchksymver
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

Yoan Bernabeu - [yoanbernabeu](https://github.com/yoanbernabeu)