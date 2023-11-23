package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/xuri/excelize/v2"
)

// Função para converter arquivos Excel em CSV
func convertExcelToCSV(sourceDir, targetDir string) error {
	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		return err
	}

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		os.Mkdir(targetDir, os.ModePerm)
	}

	bar := progressbar.Default(int64(len(files)), "Convertendo Excel para CSV")
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".xlsx") {
			excelFilePath := filepath.Join(sourceDir, file.Name())
			csvFilePath := filepath.Join(targetDir, strings.TrimSuffix(file.Name(), ".xlsx")+".csv")

			f, err := excelize.OpenFile(excelFilePath)
			if err != nil {
				return err
			}

			csvFile, err := os.Create(csvFilePath)
			if err != nil {
				return err
			}
			defer csvFile.Close()

			writer := bufio.NewWriter(csvFile)
			rows, err := f.GetRows(f.GetSheetName(1))
			if err != nil {
				return err
			}
			for _, row := range rows {
				writer.WriteString(strings.Join(row, ",") + "\n")
			}
			writer.Flush()
			bar.Add(1)
		}
	}
	return nil
}

// Função para concatenar arquivos CSV
func concatCSVFiles(directory, outputFile string) error {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return err
	}

	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	writer := bufio.NewWriter(output)
	bar := progressbar.Default(int64(len(files)), "Concatenando arquivos CSV")

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".csv") {
			filepath := filepath.Join(directory, file.Name())
			f, err := os.Open(filepath)
			if err != nil {
				return err
			}

			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				writer.WriteString(scanner.Text() + "\n")
			}
			f.Close()

			if err := scanner.Err(); err != nil {
				return err
			}
			bar.Add(1)
		}
	}
	return writer.Flush()
}

func main() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	sourceDir := filepath.Join(userHomeDir, "Documentos", "PythonVsGo_Dataset", "documentos_excel")
	targetDir := filepath.Join(userHomeDir, "Documentos", "PythonVsGo_Dataset", "documentos_csv")
	outputFile := filepath.Join(targetDir, "concatenated_results.csv")
	logPath := filepath.Join(userHomeDir, "Área de Trabalho", "Projetos", "PythonVsGo-FileManipulation", "PythonVsGo-FileManipulation", "logs", "concatenated_go_log.txt")

	startTime := time.Now()

	cpuUsageBefore, _ := cpu.Percent(time.Second, false)
	memUsageBefore, _ := mem.VirtualMemory()

	err = convertExcelToCSV(sourceDir, targetDir)
	if err != nil {
		panic(err)
	}

	err = concatCSVFiles(targetDir, outputFile)
	if err != nil {
		panic(err)
	}

	cpuUsageAfter, _ := cpu.Percent(time.Second, false)
	memUsageAfter, _ := mem.VirtualMemory()

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	logData := fmt.Sprintf("Tempo decorrido: %v segundos\nUso de CPU Antes: %.2f%%, Depois: %.2f%%\nUso de Memória Antes: %.2f%%, Depois: %.2f%%\n", elapsedTime.Seconds(), cpuUsageBefore[0], cpuUsageAfter[0], memUsageBefore.UsedPercent, memUsageAfter.UsedPercent)

	err = ioutil.WriteFile(logPath, []byte(logData), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Processo finalizado! Log salvo em:", logPath)
}
