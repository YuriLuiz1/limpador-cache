package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"time"
)

func main() {
	version := 1.0

	fmt.Println("Version this program is:", "V.", version)
	fmt.Println("Operating System:", runtime.GOOS)
	fmt.Println("Hello", getNamePc(), "...")
	fmt.Println("")
	fmt.Println("Program developed by: Yuri Luiz.")
	fmt.Println("Linkedin: https://www.linkedin.com/in/yuri-luiz/")
	fmt.Println("")

	alertUser()

	switch choiceUser() {
	case 1:
		searchDirectory()
	case 2:
		fmt.Println("Exit Program...")
		os.Exit(0)
	default:
		fmt.Println("Invalid option. Exiting...")
		os.Exit(1)
	}
}

func getNamePc() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Don't possible get hostname")
		return "Unknown"
	}
	return hostname
}

func getCurrentUser() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return currentUser.Username, nil
}

func getChromeDefaultCachePath() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("erro ao obter usuário atual: %v", err)
	}

	var cachePath string

	switch runtime.GOOS {
	case "windows":
		cachePath = filepath.Join(currentUser.HomeDir, "AppData", "Local", "Google", "Chrome", "User Data", "Default", "Cache", "Cache_Data")
	case "darwin": // macOS
		cachePath = filepath.Join(currentUser.HomeDir, "Library", "Caches", "Google", "Chrome", "Default", "Cache Data")
	case "linux":
		cachePath = filepath.Join(currentUser.HomeDir, ".cache", "google-chrome", "Default", "Cache")
	default:
		return "", fmt.Errorf("sistema operacional não suportado: %s", runtime.GOOS)
	}

	return cachePath, nil
}

func getAllChromeCachePaths() []string {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("Erro ao obter usuário: %v\n", err)
		return []string{}
	}

	var paths []string

	switch runtime.GOOS {
	case "windows":
		basePath := filepath.Join(currentUser.HomeDir, "AppData", "Local", "Google", "Chrome", "User Data")
		
		// Caminhos comuns do Chrome no Windows
		commonPaths := []string{
			filepath.Join(basePath, "Default", "Cache", "Cache_Data"),
			filepath.Join(basePath, "Default", "Code Cache"),
			filepath.Join(basePath, "Default", "GPUCache"),
			filepath.Join(basePath, "Profile 1", "Cache", "Cache_Data"),
			filepath.Join(basePath, "Profile 2", "Cache", "Cache_Data"),
		}
		paths = commonPaths

	case "darwin": // macOS
		basePath := filepath.Join(currentUser.HomeDir, "Library", "Caches", "Google", "Chrome")
		
		commonPaths := []string{
			filepath.Join(basePath, "Default", "Cache Data"),
			filepath.Join(basePath, "Default", "Code Cache"),
			filepath.Join(basePath, "Default", "GPUCache"),
			filepath.Join(basePath, "Profile 1", "Cache Data"),
			filepath.Join(basePath, "Profile 2", "Cache Data"),
		}
		paths = commonPaths

	case "linux":
		basePath := filepath.Join(currentUser.HomeDir, ".cache", "google-chrome")
		
		commonPaths := []string{
			filepath.Join(basePath, "Default", "Cache"),
			filepath.Join(basePath, "Default", "Code Cache"),
			filepath.Join(basePath, "Default", "GPUCache"),
			filepath.Join(basePath, "Profile 1", "Cache"),
			filepath.Join(basePath, "Profile 2", "Cache"),
		}
		paths = commonPaths
	}

	return paths
}

func alertUser() {
	fmt.Println("===============================================================")
	fmt.Println("ALERT FOR YOU USER: ")
	fmt.Println("These actions are irreversible, Make sure you want to proceed.")
	fmt.Println("This will clear Chrome browser cache from all profiles.")
	fmt.Println("================================================================")
}

func choiceUser() int {
	fmt.Println("Please digit your choice:")
	fmt.Println("1 - Clear Chrome Cache")
	fmt.Println("2 - Exit Program...")

	var digitUser int
	fmt.Scan(&digitUser)
	return digitUser
}

func searchDirectory() {
	userName, err := getCurrentUser()
	if err != nil {
		fmt.Printf("Error getting current user: %v\n", err)
		return
	}

	fmt.Printf("Current user: %s\n", userName)
	fmt.Printf("Operating System: %s\n", runtime.GOOS)

	// Obtém todos os possíveis caminhos de cache do Chrome
	cachePaths := getAllChromeCachePaths()
	
	var foundPaths []string
	var totalFiles int

	// Verifica quais caminhos existem
	fmt.Println("\nSearching for Chrome cache directories...")
	for _, path := range cachePaths {
		if _, err := os.Stat(path); err == nil {
			fmt.Printf("✓ Found: %s\n", path)
			foundPaths = append(foundPaths, path)
			
			// Conta arquivos no diretório
			if files, err := os.ReadDir(path); err == nil {
				fileCount := 0
				for _, file := range files {
					if !file.IsDir() {
						fileCount++
					}
				}
				totalFiles += fileCount
				fmt.Printf("  └─ Contains %d files\n", fileCount)
			}
		}
	}

	if len(foundPaths) == 0 {
		fmt.Println("❌ No Chrome cache directories found!")
		fmt.Println("Make sure Chrome is installed and has been used at least once.")
		return
	}

	fmt.Printf("\nFound %d cache directories with approximately %d files total.\n", len(foundPaths), totalFiles)
	
	fmt.Println("\n⚠️  WARNING: Clear ALL Chrome Cache? *THIS ACTION IS IRREVERSIBLE*")
	fmt.Println("This will delete all cached data from Chrome browser.")
	fmt.Println("Digit 3 to CONFIRM or any other number to cancel:")
	
	var digitUser int
	fmt.Scan(&digitUser)

	if digitUser == 3 {
		fmt.Println("\n🧹 Starting cache cleanup...")
		
		totalDeleted := 0
		totalErrors := 0

		for _, path := range foundPaths {
			fmt.Printf("\nCleaning: %s\n", path)
			
			directory, err := os.ReadDir(path)
			if err != nil {
				fmt.Printf("❌ Error reading directory %s: %v\n", path, err)
				totalErrors++
				continue
			}

			for _, content := range directory {
				pathCompleted := filepath.Join(path, content.Name())
				
				err := os.RemoveAll(pathCompleted)
				if err != nil {
					fmt.Printf("❌ Error deleting: %s - %v\n", pathCompleted, err)
					totalErrors++
				} else {
					fmt.Printf("✓ Deleted: %s\n", content.Name())
					totalDeleted++
				}
			}
		}

		fmt.Printf("\n🎉 Cleanup completed!\n")
		fmt.Printf("✅ Successfully deleted: %d items\n", totalDeleted)
		if totalErrors > 0 {
			fmt.Printf("❌ Errors encountered: %d items\n", totalErrors)
		}
		fmt.Println("💡 Tip: Restart Chrome to see the effects.")
		time.Sleep(20 * time.Second)
		
	} else {
		fmt.Println("❌ Operation cancelled by user.")
		time.Sleep(3 * time.Second)
	}
}

// Função auxiliar para verificar se um caminho existe
func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Função para obter informações detalhadas sobre o sistema
func getSystemInfo() {
	fmt.Printf("System Information:\n")
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	
	if user, err := user.Current(); err == nil {
		fmt.Printf("Username: %s\n", user.Username)
		fmt.Printf("Home Directory: %s\n", user.HomeDir)
	}
	
	if hostname, err := os.Hostname(); err == nil {
		fmt.Printf("Hostname: %s\n", hostname)
	}
}