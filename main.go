package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Create the PowerShell script content
	psScript := `Write-Host "Hello, World!"`

	// Define the file path and name for the PowerShell script
	scriptPath := "C:\\Path\\To\\Script.ps1"

	// Create the script file
	file, err := os.Create(scriptPath)
	if err != nil {
		fmt.Println("Error creating script file:", err)
		return
	}
	defer file.Close()

	// Write the script content to the file
	_, err = file.WriteString(psScript)
	if err != nil {
		fmt.Println("Error writing to script file:", err)
		return
	}

	// Get the startup folder path for the administrator user
	startupFolder := getStartupFolder("Administrator")
	if startupFolder == "" {
		fmt.Println("Failed to get startup folder for the administrator user")
		return
	}

	// Create a shortcut to the script in the startup folder
	err = createShortcut(scriptPath, startupFolder)
	if err != nil {
		fmt.Println("Error creating shortcut:", err)
		return
	}

	fmt.Println("Shortcut created successfully.")
}

// getStartupFolder returns the path to the startup folder for the specified user.
func getStartupFolder(username string) string {
	userFolder, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return ""
	}

	startupFolder := filepath.Join(userFolder, "AppData", "Roaming", "Microsoft", "Windows", "Start Menu", "Programs", "Startup")
	return startupFolder
}

// createShortcut creates a shortcut to the specified file in the specified folder.
func createShortcut(filePath, folderPath string) error {
	shortcutPath := filepath.Join(folderPath, "Script.lnk")

	shortcut, err := os.Create(shortcutPath)
	if err != nil {
		return err
	}
	defer shortcut.Close()

	// Create a PowerShell command that runs the script
	psCommand := fmt.Sprintf("powershell.exe -ExecutionPolicy Bypass -NoLogo -NoProfile -WindowStyle Hidden -File \"%s\"", filePath)

	// Write the shortcut target information
	_, err = shortcut.WriteString("[InternetShortcut]\r\nURL=file:///" + strings.ReplaceAll(psCommand, "\\", "/"))
	if err != nil {
		return err
	}

	return nil
}
