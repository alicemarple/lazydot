# 🍪 Lazydot

Pacman like dot file manager

## 🎦 Image

![lazydot_sq](https://res.cloudinary.com/dljmvvlte/image/upload/v1750706616/lazydot_sQ_lxflho.png)

## 🍔 Installation

Install my-project with go

```
Download it from release
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
lazydot -s dotfile_name
```

- Search the dotfile on remote

```bash
lazydot -s all
```

- List all dotfiles from remote

## 🌸 Documentation

Main files :

- local.yml - for download dotfiles metadata
- sync.yml - for remote dotfiles metadata
- config.yml - for config
