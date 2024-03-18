package ui

import (
	"bufio"
	"github.com/charmbracelet/log"
	"os"
	"strconv"
)

type SelectOption struct {
	Key   string
	Value string
}

func Select(options []SelectOption) string {
	ClearScreen()

	reader := bufio.NewReader(os.Stdin)

	log.Info("Select an option:")
	for id, option := range options {
		println(id, " - ", option.Value)
	}

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	index, _ := strconv.Atoi(text[:len(text)-1])

	ClearScreen()

	return options[index].Key
}
