package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const baseURL = "https://storage.googleapis.com/chrome-for-testing-public"

type ChromeVersions struct {
	Stable struct {
		Version string `json:"version"`
	} `json:"Stable"`
}

type ChromeDownloads struct {
	Chromedriver []struct {
		Platform string `json:"platform"`
		URL      string `json:"url"`
	} `json:"chromedriver"`
}

func getSeleniumPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "AppData", "Local", "SeleniumBasic"), nil
}

func getLatestVersion() (string, error) {
	resp, err := http.Get("https://googlechromelabs.github.io/chrome-for-testing/last-known-good-versions-with-downloads.json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data struct {
		Versions ChromeVersions `json:"channels"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	return data.Versions.Stable.Version, nil
}

func getDownloadURL(_ string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://googlechromelabs.github.io/chrome-for-testing/last-known-good-versions-with-downloads.json"))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data struct {
		Channels struct {
			Stable struct {
				Downloads struct {
					Chromedriver []struct {
						Platform string `json:"platform"`
						URL      string `json:"url"`
					} `json:"chromedriver"`
				} `json:"downloads"`
			} `json:"Stable"`
		} `json:"channels"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	for _, download := range data.Channels.Stable.Downloads.Chromedriver {
		if download.Platform == "win64" {
			return download.URL, nil
		}
	}

	return "", fmt.Errorf("URL para Windows 64 bits não encontrada")
}

func downloadFile(url, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func unzip(src, dest string) error {
	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		cleanPath := strings.TrimPrefix(file.Name, "chromedriver-win64/")
		if cleanPath == "" {
			continue
		}

		path := filepath.Join(dest, cleanPath)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
			continue
		}

		outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()

		zipFile, err := file.Open()
		if err != nil {
			return err
		}
		defer zipFile.Close()

		_, err = io.Copy(outFile, zipFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func getInstalledVersion(seleniumPath string) (string, error) {
	cmd := exec.Command(filepath.Join(seleniumPath, "chromedriver.exe"), "--version")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	parts := strings.Split(string(output), " ")
	if len(parts) >= 2 {
		return parts[1], nil
	}
	return "", fmt.Errorf("não foi possível obter a versão instalada")
}

func clearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func showMenu(latestVersion, installedVersion string) int {
	clearTerminal()
	fmt.Println("Atualizador ChromeDriver")
	fmt.Println()
	fmt.Println("https://github.com/rRopelato")
	fmt.Println()
	fmt.Println("Versão disponível do ChromeDriver:", latestVersion)
	fmt.Println("Versão instalada do ChromeDriver:", installedVersion)
	fmt.Println()
	fmt.Println("1. Atualizar ChromeDriver")
	fmt.Println("2. Sair")
	fmt.Println()
	var choice int
	fmt.Scanln(&choice)
	return choice
}

func main() {

	cmd := exec.Command("title", "ChromeDriver Updater | @rRopelato")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}

	seleniumPath, err := getSeleniumPath()
	if err != nil {
		fmt.Println("Erro ao obter o caminho do SeleniumBasic:", err)
		return
	}

	installedVersion, err := getInstalledVersion(seleniumPath)
	if err != nil {
		fmt.Println("Erro ao verificar a versão instalada:", err)
		installedVersion = "Não instalada"
	}

	latestVersion, err := getLatestVersion()
	if err != nil {
		fmt.Println("Erro ao buscar a versão mais recente:", err)
		return
	}

	for {
		choice := showMenu(latestVersion, installedVersion)
		switch choice {
		case 1:
			clearTerminal()
			fmt.Println("Baixando ChromeDriver versão", latestVersion)
			url, err := getDownloadURL(latestVersion)
			if err != nil {
				fmt.Println("Erro ao obter URL de download:", err)
				continue
			}

			tempZip := "chromedriver.zip"
			fmt.Println("Baixando de:", url)
			if err := downloadFile(url, tempZip); err != nil {
				fmt.Println("Erro ao baixar o arquivo:", err)
				continue
			}

			tempFolder := "chromedriver_temp"
			os.MkdirAll(tempFolder, os.ModePerm)

			fmt.Println("Extraindo ChromeDriver...")
			if err := unzip(tempZip, tempFolder); err != nil {
				fmt.Println("Erro ao extrair o arquivo:", err)
				continue
			}

			fmt.Println("Atualizando ChromeDriver na pasta SeleniumBasic...")
			err = filepath.Walk(tempFolder, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() && strings.Contains(info.Name(), "chromedriver.exe") {
					destPath := filepath.Join(seleniumPath, "chromedriver.exe")
					os.Remove(destPath)
					return os.Rename(path, destPath)
				}
				return nil
			})

			if err != nil {
				fmt.Println("Erro ao atualizar ChromeDriver:", err)
			} else {
				fmt.Println("ChromeDriver atualizado com sucesso!")
			}

			os.Remove(tempZip)
			os.RemoveAll(tempFolder)

			installedVersion, err = getInstalledVersion(seleniumPath)
			if err != nil {
				fmt.Println("Erro ao verificar a versão instalada:", err)
				installedVersion = "Não instalada"
			}
		case 2:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida, tente novamente.")
		}
	}
}
