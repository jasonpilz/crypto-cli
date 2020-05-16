package prompt

import (
	"log"
	"os"
	"path/filepath"

	"github.com/peterh/liner"
)

var histPath = tmpPath(".crypto.history")

func tmpPath(filename string) string {
	return filepath.Join(os.TempDir(), filename)
}

func LinerPrompt(label string) string {
	line := newLiner(histPath)
	defer line.Close()

	in, err := line.Prompt(label)

	if err == nil {
		line.AppendHistory(in)
	} else if err == liner.ErrPromptAborted {
		log.Print("Aborted reading line")
	} else {
		log.Print("Error reading line: ", err)
	}

	saveHistory(line, histPath)

	return in
}

func LinerPromptSuggest(label, suggest string) string {
	line := newLiner(histPath)
	defer line.Close()

	in, err := line.PromptWithSuggestion(label, suggest, -1)

	if err == liner.ErrPromptAborted {
		log.Print("Aborted reading line")
	}

	return in
}

func newLiner(historyPath string) *liner.State {
	line := liner.NewLiner()

	line.SetCtrlCAborts(true)

	if f, err := os.Open(historyPath); err == nil {
		line.ReadHistory(f)
		f.Close()
	}

	return line
}

func saveHistory(line *liner.State, historyPath string) {
	if f, err := os.Create(historyPath); err != nil {
		log.Print("Error writing history file: ", err)
	} else {
		line.WriteHistory(f)
		f.Close()
	}
}
