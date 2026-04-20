# Translate 64

**Translate 64** is a minimalist, native desktop utility for encoding and decoding strings across multiple Base64 standards. Built with a local-first philosophy, it provides a fast, secure alternative to web-based converters.

## Important: Base64 is encoding, not encryption
Base64 does not provide security. It is a reversible encoding format, not a cryptographic method.

## Features

* **Multiple RFC Modes** Supports Standard (RFC 4648), URL-safe, MIME (RFC 2045), and Unpadded variants.
* **Offline-First Security** All encoding/decoding happens locally in vanilla JavaScript. No network requests, no data transmission.
* **Developer-Centric UI** A simple dual-pane interface designed for quick workflows without framework overhead.
* **Native Desktop Performance** Runs as a lightweight Wails application for a clean, cross-platform experience.

## Code Structure

* **Logic:** Built with Wails, but all core logic lives in HTML, CSS, and Vanilla JavaScript.
* **Front-end Driven:** No backend system commands are used; the application is entirely front-end driven for maximum portability and speed.

## Platform Support

* **Linux:** Primary development and testing environment LMDE (Linux Mint Debian Edition).
* **macOS / Windows:** Compatible via the Wails compiler (formal testing pending).

## Current Status

The codebase is small, straightforward, and designed for internal utility with no external dependencies or known security concerns.

## Purpose

Translate 64 exists to support development workflows (such as the **Pharos** project) in situations where quick, local Base64 encoding or decoding is required without relying on external web tools.
