# jsonviz

`jsonviz` is a command-line tool built in Go that visualizes JSON file structures in a navigable tree format using the `termui` library. This tool is designed to help developers and data scientists quickly understand and explore the structure of complex JSON data from the terminal.

## Features

- **Interactive Tree View**: Display JSON files as expandable and collapsible trees.
- **Color Coding**: Different colors for root, keys, and values to enhance readability.
- **Keyboard Navigation**: Use keyboard shortcuts to navigate through the JSON structure.
- **Resizable UI**: Adapt to terminal size changes dynamically.

## Installation

To install `jsonviz`, clone this repository and build the tool using Go.

```bash
git clone https://github.com/copyleftdev/jsonviz.git
cd jsonviz
go build
```

## Usage

After building the tool, you can run it by specifying the JSON file you want to visualize:

```bash
./jsonviz path/to/your/file.json
```

### Keyboard Shortcuts

- **Arrow Up (`k`)**: Move up in the tree.
- **Arrow Down (`j`)**: Move down in the tree.
- **Enter**: Expand or collapse the selected node.
- **CTRL+C**: Exit the application.

## Requirements

- Go 1.15 or higher
- `termui` library (automatically installed with go get)

## Contributing

Contributions are welcome! Feel free to submit pull requests, create issues for bugs you've found, or suggest new features.

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/fooBar`).
3. Commit your changes (`git commit -am 'Add some fooBar'`).
4. Push to the branch (`git push origin feature/fooBar`).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE) file for details.

## Contact

- GitHub: [@copyleftdev](https://github.com/copyleftdev)
