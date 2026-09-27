# Neovim Commands Cheat Sheet

A practical reference for my Neovim + Go setup.

> **Tip:** When confused, press `Esc` to return to Normal mode.

---

## 1. Files

| Command | What it does |
|---|---|
| `:e main.go` | Open `main.go` |
| `:e path/to/file.go` | Open a file |
| `:e newfile.go` | Create/open a new file |
| `:w` | Save |
| `:w filename.go` | Save as another file |
| `:e!` | Reload file and discard unsaved changes |
| `:q` | Close current window |
| `:wq` | Save and close |
| `:q!` | Close without saving |
| `:qa` | Quit all windows |
| `:qa!` | Force quit everything |

---

## 2. Splits / Windows

A **window** is a visible pane. A **buffer** is the file being displayed.

| Command | What it does |
|---|---|
| `:split` | Horizontal split |
| `:vsplit` | Vertical split |
| `:split server.go` | Open `server.go` in a horizontal split |
| `:vsplit server.go` | Open `server.go` in a vertical split |
| `:only` | Keep only the current window |
| `:wincmd =` | Equalize window sizes |

### Move between windows

| Command | Direction |
|---|---|
| `:wincmd h` | Left |
| `:wincmd j` | Down |
| `:wincmd k` | Up |
| `:wincmd l` | Right |

### Built-in shortcut

You can also use:

```text
Ctrl-w h    left
Ctrl-w j    down
Ctrl-w k    up
Ctrl-w l    right
```

---

## 3. Terminals

### Terminal in current window

```vim
:terminal
```

This replaces the current window with a terminal.

### Terminal below

```vim
:botright split | terminal
```

This creates a terminal at the bottom.

### Terminal on the right

```vim
:botright vsplit | terminal
```

### Terminal with a specific height

```vim
:botright 15split | terminal
```

This creates a terminal about 15 lines high.

### Exit terminal mode

Press:

```text
Esc
```

Then you can use Neovim commands such as:

```vim
:q
```

---

## 4. Buffers / Open Files

A **buffer** is a file loaded into Neovim.

### List open buffers

```vim
:buffers
```

or:

```vim
:ls
```

Example:

```text
1  %a  "main.go"
2      "server.go"
3      "go.mod"
```

### Switch to a buffer

```vim
:buffer 2
```

or:

```vim
:buffer server.go
```

### Next buffer

```vim
:bnext
```

### Previous buffer

```vim
:bprevious
```

### Close/delete current buffer

```vim
:bdelete
```

---

## 5. NvimTree

Your file explorer.

| Command | What it does |
|---|---|
| `:NvimTreeOpen` | Open file tree |
| `:NvimTreeClose` | Close file tree |
| `:NvimTreeToggle` | Toggle file tree |
| `:NvimTreeFocus` | Move cursor to file tree |
| `:NvimTreeFindFile` | Find current file in tree |

Inside NvimTree:

```text
j / k       move down/up
Enter       open file or directory
q           close tree
```

---

## 6. Telescope

Fast file and text search.

```vim
:Telescope find_files
```

Find files.

```vim
:Telescope live_grep
```

Search text across the project.

```vim
:Telescope buffers
```

Find/switch open buffers.

```vim
:Telescope help_tags
```

Search Neovim help.

---

## 7. Search inside the current file

Search forward:

```vim
/pattern
```

Example:

```vim
/StartServer
```

Then:

```text
n       next match
N       previous match
```

Search backward:

```vim
?pattern
```

---

## 8. Editing in Normal Mode

These are keys, not `:` commands.

| Key | What it does |
|---|---|
| `i` | Insert before cursor |
| `I` | Insert at beginning of line |
| `a` | Insert after cursor |
| `A` | Insert at end of line |
| `o` | New line below |
| `O` | New line above |
| `Esc` | Return to Normal mode |

### Basic editing

| Key | What it does |
|---|---|
| `dd` | Delete line |
| `yy` | Copy/yank line |
| `p` | Paste |
| `u` | Undo |
| `Ctrl-r` | Redo |
| `dw` | Delete word |
| `D` | Delete to end of line |

---

## 9. Navigation

In Normal mode:

```text
h    left
j    down
k    up
l    right
```

Other useful navigation:

```text
gg    beginning of file
G     end of file
0     beginning of line
$     end of line
```

Go directly to a line:

```vim
:42
```

Go to line 100:

```vim
:100
```

---

## 10. Go Development

Format current Go file:

```vim
:GoFmt
```

Check LSP information:

```vim
:LspInfo
```

Open diagnostics:

```text
Space d
```

Go to definition:

```text
gd
```

Find references:

```text
gr
```

Hover documentation:

```text
K
```

Rename symbol:

```text
Space r n
```

Code action:

```text
Space c a
```

---

## 11. Your Custom Keybindings

Leader key is **Space**.

| Shortcut | Action |
|---|---|
| `Space w` | Save |
| `Space q` | Close current window |
| `Space e` | Toggle NvimTree |
| `Space o` | Focus NvimTree |
| `Space n f` | Find current file in NvimTree |
| `Space f f` | Find files |
| `Space f g` | Search project text |
| `Space f b` | Find open buffers |
| `Space f h` | Search help |
| `Space g f` | Go format |
| `Space t t` | Open terminal below |
| `Space t v` | Open terminal vertically |
| `Space r n` | Rename symbol |
| `Space c a` | Code action |
| `Space d` | Show diagnostic |

> If a custom shortcut behaves unexpectedly, use the `:` command instead. The `:` commands are the reliable fundamentals.

---

## 12. Lazy.nvim

Plugin manager.

```vim
:Lazy
```

Open plugin manager.

```vim
:Lazy update
```

Update plugins.

```vim
:Lazy install
```

Install missing plugins.

```vim
:Lazy clean
```

Remove unused plugins.

---

## 13. Configuration

Open your config:

```vim
:e ~/.config/nvim/init.lua
```

Reload config:

```vim
:source ~/.config/nvim/init.lua
```

Show config location:

```vim
:echo $MYVIMRC
```

---

## 14. Help

Neovim has built-in documentation.

```vim
:help
```

Specific help:

```vim
:help split
:help terminal
:help buffer
:help wincmd
```

Search help:

```vim
:helpgrep terminal
```

---

## 15. Health Checks

General health:

```vim
:checkhealth
```

LSP health:

```vim
:checkhealth vim.lsp
```

---

# Quick Reference

## Most important commands

```text
:e file.go                    Open file
:w                            Save
:q                            Close current window
:wq                           Save + close
:q!                           Close without saving

:split                        Horizontal split
:vsplit                       Vertical split
:only                         Keep only current window

:botright split | terminal    Terminal below
:botright vsplit | terminal   Terminal on right

:buffers                      List open files
:bnext                        Next buffer
:bprevious                    Previous buffer
:buffer NAME                  Switch to file

:NvimTreeToggle               Toggle file tree
:Telescope find_files         Find files
:Telescope live_grep          Search project

:LspInfo                      LSP status
:GoFmt                        Format Go file
:checkhealth                  Health check
:Lazy                         Plugin manager
```

## Most important Normal-mode keys

```text
Esc       Normal mode
i         Insert mode
h j k l   Move
dd        Delete line
yy        Copy line
p         Paste
u         Undo
Ctrl-r    Redo
gg        Top of file
G         Bottom of file

Ctrl-w h   Window left
Ctrl-w j   Window down
Ctrl-w k   Window up
Ctrl-w l   Window right
```

---

# Mental Model

```text
Neovim
│
├── Buffers = files loaded
│
├── Windows = visible panes
│   ├── main.go
│   ├── server.go
│   └── terminal
│
├── NvimTree = file browser
│
└── Telescope = search/navigation
```

When confused:

```text
1. Press Esc
2. Decide whether you want a file, window, or terminal
3. Use the appropriate : command
```
