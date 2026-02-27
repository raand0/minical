# Minical

Lightweight TUI calendar.

<img width="500" height="500" alt="screenshot-2026-02-27_18-58-51" src="https://github.com/user-attachments/assets/33ec7277-24a5-4b55-88fd-868288b83476" />

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GitHub Stars](https://img.shields.io/github/stars/yourusername/yourproject)](https://github.com/raand0/minical)

A fast and reliable TUI calendar. Styling is customizable via a config file, also if you know golang you can fully customize it.
This calendar only allows navigation do not expect to add events or reminders, ONLY for navigating between different times

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Customization](#customization)

## Installation

You can get it from yay AUR helper.

```bash
yay -S minical
```

## Usage

After installing simply type:
```bash
minical
```

Check available commands:
```bash
minical --help
```

Available commands:
```bash
minical --gen-config
    Generates default config file
minical --today
    Returns current date in YYYY-MM-DD format
```

## Customization

Even though there is not much to customize but basic color and symbol customization can be done in **~/.config/minical/config**
