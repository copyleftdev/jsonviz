package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type nodeValue struct {
	Key   string
	Value interface{}
}

var (
	rootStyle     = ui.NewStyle(ui.ColorGreen)
	keyStyle      = ui.NewStyle(ui.ColorGreen)
	valueStyle    = ui.NewStyle(ui.ColorWhite)
	selectedStyle = ui.NewStyle(ui.ColorYellow, ui.ColorClear, ui.ModifierBold) // Correctly apply style with modifiers
)

func (nv nodeValue) String() string {
	if nv.Key == "root" {
		return fmt.Sprintf("[%s](fg:green)", nv.Key)
	}
	switch v := nv.Value.(type) {
	case string, float64, bool, int:
		return fmt.Sprintf("[%s: %v](fg:white)", nv.Key, v)
	default:
		return fmt.Sprintf("[%s](fg:green)", nv.Key)
	}
}

func parseJSONToTree(key string, data interface{}) *widgets.TreeNode {
	node := &widgets.TreeNode{
		Value: nodeValue{Key: key},
	}

	switch v := data.(type) {
	case map[string]interface{}:
		children := []*widgets.TreeNode{}
		for k, val := range v {
			child := parseJSONToTree(k, val)
			children = append(children, child)
		}
		node.Nodes = children
	case []interface{}:
		children := []*widgets.TreeNode{}
		for i, val := range v {
			child := parseJSONToTree(fmt.Sprintf("%s[%d]", key, i), val)
			children = append(children, child)
		}
		node.Nodes = children
	default:
		node.Value = nodeValue{Key: key, Value: v}
	}

	return node
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: jsonviz [filename]")
		os.Exit(1)
	}

	filename := os.Args[1]

	fmt.Println("Initializing UI... Press CTRL+C to exit.")

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open JSON file: %v", err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var result interface{}
	json.Unmarshal(byteValue, &result)

	root := parseJSONToTree("root", result)
	tree := widgets.NewTree()
	tree.Title = "JSON Visualizer (Press CTRL+C to exit)"
	tree.TextStyle = ui.NewStyle(ui.ColorWhite)
	tree.SelectedRowStyle = selectedStyle
	tree.WrapText = false
	tree.SetNodes([]*widgets.TreeNode{root})

	x, y := ui.TerminalDimensions()
	tree.SetRect(0, 0, x, y)

	ui.Render(tree)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			tree.ScrollDown()
		case "k", "<Up>":
			tree.ScrollUp()
		case "<Enter>":
			tree.ToggleExpand()
		case "<Resize>":
			x, y := ui.TerminalDimensions()
			tree.SetRect(0, 0, x, y)
			ui.Render(tree)
		}

		ui.Render(tree)
	}
}
