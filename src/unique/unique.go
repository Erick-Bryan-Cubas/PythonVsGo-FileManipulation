package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/schollz/progressbar/v3"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/xuri/excelize/v2"
)

func main() {
	// Caminhos
	userprofile, _ := os.UserHomeDir()
	filesPath := filepath.Join(userprofile, "Documentos", "PythonVsGo_Dataset", "documentos_excel")
	logPath := filepath.Join(userprofile, "Área de Trabalho", "Projetos", "PythonVsGo-FileManipulation", "PythonVsGo-FileManipulation", "logs", "unique_go_log.txt")

	// Monitoramento de desempenho (início)
	startTime := time.Now()
	cpuUsageBefore, _ := cpu.Percent(0, false)
	memUsageBefore, _ := mem.VirtualMemory()

	// Processamento dos arquivos
	uniqueLinesCount, err := countUniqueLines(filesPath)
	if err != nil {
		fmt.Println("Erro ao processar arquivos:", err)
		return
	}

	// Monitoramento de desempenho (fim)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	cpuUsageAfter, _ := cpu.Percent(0, false)
	memUsageAfter, _ := mem.VirtualMemory()

	// Salvando em um arquivo de log
	logData := fmt.Sprintf("Time elapsed: %v\nCPU Usage Before: %v%%, After: %v%%\nMemory Usage Before: %v%%, After: %v%%\n",
		elapsedTime, cpuUsageBefore, cpuUsageAfter, memUsageBefore.UsedPercent, memUsageAfter.UsedPercent)

	for file, count := range uniqueLinesCount {
		logData += fmt.Sprintf("%s: %d linhas únicas\n", file, count)
	}

	if err := os.WriteFile(logPath, []byte(logData), 0644); err != nil {
		fmt.Println("Erro ao escrever no arquivo de log:", err)
		return
	}

	fmt.Println("Process finished!")
}

// countUniqueLines conta linhas únicas em arquivos Excel no diretório especificado
func countUniqueLines(dir string) (map[string]int, error) {
	uniqueLines := make(map[string]int)
	files, err := filepath.Glob(filepath.Join(dir, "*.xlsx"))
	if err != nil {
		return nil, err
	}

	bar := progressbar.Default(int64(len(files)), "Processando arquivos")

	for _, file := range files {
		f, err := excelize.OpenFile(file)
		if err != nil {
			return nil, err
		}

		rows, err := f.GetRows(f.GetSheetName(1))
		if err != nil {
			return nil, err
		}

		uniqueRowCount := countUniqueRows(rows)
		uniqueLines[filepath.Base(file)] = uniqueRowCount

		bar.Add(1)
	}

	return uniqueLines, nil
}

// countUniqueRows conta as linhas únicas em uma lista de linhas
func countUniqueRows(rows [][]string) int {
	uniqueRows := make(map[string]bool)
	for _, row := range rows {
		rowStr := fmt.Sprintf("%v", row)
		uniqueRows[rowStr] = true
	}
	return len(uniqueRows)
}
