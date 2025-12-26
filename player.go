package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/blang/mpv"
)

var (
	mpvClient *mpv.Client
	mpvCmd    *exec.Cmd
)

func initPlayer() error {
	socketPath := "/tmp/depthtui_mpv_socket"

	args := []string{
		"--idle",
		"--no-video",
		"--input-ipc-server=" + socketPath,
	}

	mpvCmd = exec.Command("mpv", args...)
	if err := mpvCmd.Start(); err != nil {
		return fmt.Errorf("failed to start mpv: %v", err)
	}

	time.Sleep(500 * time.Millisecond)

	client := mpv.NewClient(mpv.NewIPCClient(socketPath))
	mpvClient = client

	return nil
}

func playSong(songID string) error {
	if mpvClient == nil {
		return fmt.Errorf("player not initialized")
	}

	url := subsonicStream(songID)
	fmt.Println("Ordering MPV to play:", url)

	if err := mpvClient.Loadfile(url, mpv.LoadFileModeReplace); err != nil {
		return err
	}

	mpvClient.SetProperty("pause", false)

	return nil
}

func shutdownPlayer() {
	if mpvCmd != nil {
		mpvCmd.Process.Kill()
	}
}
