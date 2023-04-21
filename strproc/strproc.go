package strproc

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go-reloaded/utils"
)

func Process(inputFile, outputFile string) error {
	inputExt, outputExt := filepath.Ext(inputFile), filepath.Ext(outputFile)
	if inputExt != ".txt" || outputExt != ".txt" {
		return fmt.Errorf("input and output file must have .txt extension")
	}

	if inputFile == outputFile {
		return fmt.Errorf("output file should have different name")
	}
	bs, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("error reading input file: %v", err)
	}
	data := strings.Split(string(bs), "\n")
	outputLines := []string{}
	for i := range data {
		s := strings.Fields(data[i])
		sliceAfterModifier := []string{}
		for i := 0; i < len(s); i++ {
			if utils.CheckModifier(&s, &sliceAfterModifier, s[i], i) {
				i++
			}
		}
		sliceAfterPunctuation := []string{}
		for i := 0; i < len(sliceAfterModifier); i++ {
			utils.CheckPunctuation(&sliceAfterModifier, &sliceAfterPunctuation, i)
		}
		sliceAfterSpecialCases := utils.CheckSpecialCase(sliceAfterPunctuation)
		sliceAfterSpecialQuotes := utils.AppendQuotes(&sliceAfterSpecialCases)
		sliceAfterQuotes := utils.CheckQuotes(&sliceAfterSpecialQuotes)
		outputLines = append(outputLines, strings.TrimRight(strings.Join(sliceAfterQuotes, " "), "\n"))
		if i != len(data)-1 {
			outputLines = append(outputLines, "\n")
		}
	}
	err = os.WriteFile(outputFile, []byte(strings.Join(outputLines, "")), 0o666)
	if err != nil {
		return fmt.Errorf("error writing output file: %v", err)
	}
	return nil
}
