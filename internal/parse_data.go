package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Block struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	DisplayName string  `json:"displayName"`
	Hardness    float64 `json:"hardness"`
	StackSize   int     `json:"stackSize"`
	BoundingBox string  `json:"boundingBox"`

	IsCraftable bool `json:"isCraftable"`
}

type Entity struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	DisplayName string  `json:"displayName"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	Type        string  `json:"type"`
	Category    string  `json:"category"`
}

func loadJSON(path string, v interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func markCraftable(blocks map[int]*Block, recipes map[string]interface{}) {
	for itemID := range recipes {
		for _, block := range blocks {
			if block.Name == itemID || fmt.Sprintf("%d", block.ID) == itemID {
				block.IsCraftable = true
				break
			}
		}
	}
}

func getCraftableItems(version string) (map[string]bool, error) {
	basePath := filepath.Join("minecraft-data", "data", "pc", version)

	var recipes map[string]interface{}
	if err := loadJSON(filepath.Join(basePath, "recipes.json"), &recipes); err != nil {
		return nil, err
	}

	craftable := make(map[string]bool)
	for itemName := range recipes {
		craftable[itemName] = true
	}

	return craftable, nil
}

func ParseMinecraftData(version string) (map[int]*Block, error) {
	basePath := filepath.Join("data", "pc", version)

	var rawBlocks []Block
	err := loadJSON(filepath.Join(basePath, "blocks.json"), &rawBlocks)
	if err != nil {
		return nil, err
	}

	craftable, _ := getCraftableItems(version)

	blocks := make(map[int]*Block)
	for i := range rawBlocks {
		block := &rawBlocks[i]
		block.IsCraftable = craftable[block.Name]
		blocks[block.ID] = block
	}

	return blocks, nil
}
