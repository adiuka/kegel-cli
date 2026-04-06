package ui

import (
	"github.com/manifoldco/promptui"
)

type MenuChoice int

const (
	MenuTrain			MenuChoice = 0
	MenuProgress	MenuChoice = 1
	MenuInfo			MenuChoice = 2
)

func ShowMenu() (MenuChoice, error) {
	prompt := promptui.Select{
		Label: "Options",
		Items: []string{
			"Train",
			"My Progress",
			"Information",
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		return 0, err
	}

	return MenuChoice(index), nil
}