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
	userprofile, _ := os.UserHomeDir()
	filesPath := filepath.Join(userprofile, "Documentos", "PythonVsGo_Dataset", "documentos_excel")
	logPath := filepath.Join(userprofile, "Área de Trabalho", "Projetos", "PythonVsGo-FileManipulation", "PythonVsGo-FileManipulation", "logs")

	startTime := time.Now()
	cpuUsageBefore, _ := cpu.Percent(0, false)
	memUsageBefore, _ := mem.VirtualMemory()

	cumsumResults, err := calculateCumsum(filesPath)
	if err != nil {
		fmt.Println("Erro ao processar arquivos:", err)
		return
	}

	cpuUsageAfter, _ := cpu.Percent(0, false)
	memUsageAfter, _ := mem.VirtualMemory()
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	logFile, err := os.Create(filepath.Join(logPath, "cumsum_go_log.txt"))
	if err != nil {
		fmt.Println("Erro ao criar arquivo de log:", err)
		return
	}
	defer logFile.Close()

	fmt.Fprintf(logFile, "Time elapsed: %v\nCPU Usage Before: %v%%, After: %v%%\nMemory Usage Before: %v MB, After: %v MB\n",
		elapsedTime, cpuUsageBefore[0], cpuUsageAfter[0], memUsageBefore.Used/1024/1024, memUsageAfter.Used/1024/1024)
	for file, result := range cumsumResults {
		fmt.Fprintf(logFile, "%s: Cumsum = %v\n", file, result)
	}

	fmt.Println("Process finished!")
}

// calculateCumsum calcula a soma cumulativa da primeira coluna numérica de cada arquivo Excel
func calculateCumsum(dir string) (map[string]float64, error) {
	results := make(map[string]float64)
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

		cumsum, err := calculateFileCumsum(rows)
		if err != nil {
			return nil, err
		}

		results[filepath.Base(file)] = cumsum

		bar.Add(1)
	}

	return results, nil
}

// calculateFileCumsum calcula a soma cumulativa das linhas de um arquivo Excel
func calculateFileCumsum(rows [][]string) (float64, error) {
	var sum float64
	for _, row := range rows {
		if len(row) > 0 {
			var value float64
			_, err := fmt.Sscanf(row[0], "%f", &value)
			if err != nil {
				continue // Ignora células não numéricas
			}
			sum += value
		}
	}
	return sum, nil
}
