# 🍪 Lazydot

Pacman like dot file manager

## 🎦 Image

Image

## 🍔 Installation

Install my-project with go

```bash
go get github.com/alice/lazydot
```

## 🍁 Usage

```bash
lazydot -S dotfile_name
```

- Download dotfile

```bash
lazydot -y
```

- Update the remote dotfiles metadata

```bash
lazydot -Q
```

- Query the local database of metadata

```bash
lazydot -R dotfile_name
```

- Remove dotfile

```bash
lazydot -s
```

- Search the dotfile on remote

## 🌸 Documentation

Main files :

- local.yml - for download dotfiles metadata
- sync.yml - for remote dotfiles metadata
- config.yml - for config
