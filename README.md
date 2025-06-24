# 🍪 Lazydot

Pacman like dot file manager

## 🎦 Image

![lazydot_sq](https://res.cloudinary.com/dljmvvlte/image/upload/v1750706616/lazydot_sQ_lxflho.png)

## 🍔 Installation

```
git clone https://github.com/alicemarple/lazydot.git
cd lazydot
go build -o lazydot main.go
```

You can download it from releases and add it to the path for better use.

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
lazydot -Q package_name
```

- Query the local database for package_name metadata

```bash
lazydot -Q all
```

- Query the local database for all package metadata

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

- Manage local.yml and sync.yml for better operations.

- Commands: sync, update, query, search, remove.

- Follows XDG conventions:

  - ~/.config/lazydot/config.yml (user settings)

  - ~/.local/share/lazydot/ (state files)

  - ~/.cache/lazydot/pkg/ (temporary downloads)
