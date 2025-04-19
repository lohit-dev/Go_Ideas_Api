package storage

import (
	"fmt"
	"io"
	"os"
	"test_project/test/internal/model"
	utils "test_project/test/pkg"
)

func (js *JsonStore) ReadFile() utils.Result[[]model.Idea] {
	file, err := os.Open(js.filepath)
	if err != nil {
		return utils.Result[[]model.Idea]{Err: fmt.Errorf("failed to open file: %v", err)}
	}

	defer file.Close()

	// File Contents
	fc, err := io.ReadAll(file)
	if err != nil {
		return utils.Result[[]model.Idea]{Err: fmt.Errorf("failed to read file: %v", err)}
	}

	// Unmarshal JSON
	var ideas []model.Idea
	if err := utils.UnmarshalJson(fc, &ideas); err != nil {
		return utils.Result[[]model.Idea]{Err: fmt.Errorf("failed to unmarshal json: %v", err)}
	}

	return utils.Result[[]model.Idea]{Data: ideas}
}

func (js *JsonStore) WriteJson(filepath string, data []byte) utils.Result[string] {
	file, err := os.Create(filepath)
	if err != nil {
		return utils.Result[string]{Err: fmt.Errorf("could not open file for writing: %v", err)}
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return utils.Result[string]{Err: fmt.Errorf("could not write data to file: %v", err)}
	}

	return utils.Result[string]{Data: fmt.Sprintf("Successfully saved data to %s", filepath)}
}
