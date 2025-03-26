# SimilarCLI 🌐

A powerful command-line tool for fetching website analytics from SimilarWeb. Get instant insights about any website's traffic, engagement, and more! 🚀

## Features ✨

- 📊 Basic website information (title, description, category)
- 📈 Monthly visits with human-readable numbers (K, M)
- 🎯 Traffic sources breakdown
- 🎨 Clean and modern table output
- ⚡ Fast and efficient

## Installation 🛠️

```bash
go install github.com/matisiekpl/similarcli@latest
```

## Usage 🎮

The simplest way to use SimilarCLI is to provide a domain name:

```bash
similarcli example.com
```

### Examples 📝

```bash
# Get analytics for GitHub
similarcli github.com

# Get analytics for Google
similarcli google.com

# Show help
similarcli
```

### Output Example 📊

```
Site Name    GitHub
Title        GitHub: Where the world builds software
Description  GitHub is where over 100 million developers shape the future of software, together.
Category     Technology

Date    December 2024    January 2025    February 2025
Visits  1.5M            2.3M            1.8M

Source    Direct    Search    Social    Referrals    Mail    Paid Referrals
Percentage    45.2%    25.3%    15.1%    8.4%        3.2%    2.8%
```

## Author 👨‍💻

Created by [Mateusz Woźniak](https://github.com/matisiekpl)

## License 📄

MIT License
