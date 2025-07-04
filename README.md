<h1 align="center">cellscape</h1>

<p align="center">
Terminal-based cellular automata playground
</p>

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

## License

[MIT License](LICENSE)
