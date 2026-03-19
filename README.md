# Pokedex CLI

## Introduction

### About

This is a CLI tool imitating a "pokedex" using the [pokeapi.co](https://pokeapi.co) API.

It is a guided project from [boot.dev](https://boot.dev), made for their students to further practice writing HTTP clients in Go.

You can see my [TypeScript version here](https://github.com/bntrtm/ts-pokedex).

### Features

- page through a map of locations
- explore locations for pokemon
- throw pokeballs to catch pokemon
- list pokemon in your pokedex
- inspect pokemon within your pokedex
- in-memory cache

### Differences

As previously mentioned, this is the result of following a guided project. But apart from style, there are some differences/additions of my own built on top of it. Namely, this project builds consistent endpoint URLs in a hierarchical fashion. This avoids needless repetition of endpoint construction, allowing endpoints to be defined in one single place.

## Installation

With Go installed, you can use:

```bash
go install github.com/bntrtm/pokedex-cli@latest
```

## Usage

Run it with `ts-pokedex`.

You can then run `help` to see all available commands.
