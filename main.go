package main

import (
	"fmt"
	"os"

	"github.com/MattiaPun/SubTUI/internal/api"
	"github.com/MattiaPun/SubTUI/internal/integration"
	"github.com/MattiaPun/SubTUI/internal/player"
	"github.com/MattiaPun/SubTUI/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	_ = api.LoadConfig()

	if api.AppConfig.Password != "" {
		if err := player.InitPlayer(); err != nil {
			fmt.Printf("Failed to start player: %v\n", err)
		}

	}
	defer player.ShutdownPlayer()

	p := tea.NewProgram(ui.InitialModel(), tea.WithAltScreen())

	instance := integration.Init(p)
	if instance != nil {
		defer instance.Close()
		go p.Send(ui.SetDBusMsg{Instance: instance})
	}

	if _, err := p.Run(); err != nil {
		fmt.Println("Error while running program:", err)
		os.Exit(1)
	}
}
