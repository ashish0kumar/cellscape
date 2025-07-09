<h1 align="center">cellscape</h1>

<p align="center">
Terminal-based cellular automata playground
</p>

<table align="center">
  <tr>
    <td><img src="https://github.com/user-attachments/assets/e1550df3-c425-48a0-ada3-e231e4cac5b3" width="200"></td>
    <td><img src="https://github.com/user-attachments/assets/2870d84d-523a-46b8-b289-532ec4aac05d" width="200"></td>
    <td><img src="https://github.com/user-attachments/assets/92bdf10b-3bae-4b0f-85e4-0b16f9677bdf" width="200"></td>
    <td><img src="https://github.com/user-attachments/assets/f3b38933-45b5-4166-9d43-91be7683a44e" width="200"></td>
  </tr>
  <tr>
    <td><img src="https://github.com/user-attachments/assets/671f3449-bfc1-4789-bbd3-f39894bb100d" width="200"></td>
    <td><img src="https://github.com/user-attachments/assets/31e2b66d-beac-4858-a73d-81ad1517e45d" width="200"></td>
    <td><img src="https://github.com/user-attachments/assets/445d9a5e-da8d-4358-9e89-1ea0f7194ec9" width="200"></td>
    <td><img src="https://github.com/user-attachments/assets/b7e8e77f-da08-43d4-80c3-961e0d00bd9b" width="200"></td>
  </tr>
</table>

---

## Features

- Includes **Eight classic** cellular automata
- Beautiful TUI **navigation menu**
- **Real-time speed control** with up/down arrow keys
- **Focus mode** for distraction-free fullscreen viewing
- **Monochrome toggle** for terminal theme compatibility
- **Generation-based coloring** creating trippy wave effects

### Supported Automata

| Automaton             | Description                                                                | Alias    |
| --------------------- | -------------------------------------------------------------------------- | -------- |
| Conway's Game of Life | The classic cellular automaton with simple rules creating complex patterns | life     |
| Brian's Brain         | 3-state automaton with beautiful trailing patterns                         | brain    |
| Langton's Ant         | Simple ant following two rules creating emergent complexity                | ant      |
| Larger than Life      | Generalized Game of Life with extended neighborhoods                       | ltl      |
| Belousov-Zhabotinsky  | Chemical reaction simulation with spiral waves                             | belousov |
| Faders                | Multi-state cells that fade through colors before dying                    | faders   |
| Forest Fire           | Trees grow, catch fire, and burn in natural cycles                         | forest   |
| Wildfire              | Stochastic fire spread through vegetation                                  | wildfire |

---

## Installation

### Via `go install`

```bash
go install github.com/ashish0kumar/cellscape@latest
```

### Build from Source

```bash
git clone https://github.com/ashish0kumar/cellscape.git
cd cellscape
go mod tidy
go build
sudo mv cellscape /usr/local/bin/
cellscape --help
```

---

## Usage

### Interactive Menu

```bash
cellscape
```

### Direct Automaton Launch

```bash
cellscape run life
cellscape run brain -m -f
cellscape run belousov --focus --monochrome
```

**Available keywords:** life, brain, ant, ltl, belousov, faders, forest, wildfire

### Key Bindings

| Key         | Action                    |
| ----------- | ------------------------- |
| `Space`     | Pause / resume simulation |
| `↑` / `k`   | Increase speed (max 10)   |
| `↓` / `j`   | Decrease speed (min 1)    |
| `s`         | Single-step while paused  |
| `f`         | Toggle focus mode         |
| `c`         | Toggle monochrome / color |
| `r`         | Reset current automaton   |
| `q` / `Esc` | Return to menu (or exit)  |

---

## Configuration

### Command Line Options

- `--monochrome`, `-m` - Start in monochrome mode
- `--focus`, `-f` - Start in focus mode (fullscreen, no UI)

### Visual Modes

- **Colorful Mode:** Generation-based rainbow coloring with wave effects
- **Monochrome Mode:** Clean grayscale using ANSI terminal colors
- **Focus Mode:** Fullscreen grid with no status bars or help text

---

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling and layout
- [Cobra](https://github.com/spf13/cobra) - CLI framework

## Contributing

Contributions are welcome! Please feel free to open an issue or submit a pull request.

<br>

<p align="center">
	<img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/footers/gray0_ctp_on_line.svg?sanitize=true" />
</p>

<p align="center">
        <i><code>&copy 2025-present <a href="https://github.com/ashish0kumar">Ashish Kumar</a></code></i>
</p>

<div align="center">
<a href="https://github.com/ashish0kumar/cellscape/blob/main/LICENSE"><img src="https://img.shields.io/github/license/ashish0kumar/cellscape?style=for-the-badge&color=CBA6F7&logoColor=cdd6f4&labelColor=302D41" alt="LICENSE"></a>&nbsp;&nbsp;
</div>
