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

func readFiles(directory string, bar *progressbar.ProgressBar) {
	files, _ := filepath.Glob(filepath.Join(directory, "*.xlsx"))
	for _, file := range files {
		f, err := excelize.OpenFile(file)
		if err != nil {
			fmt.Println("Erro ao ler o arquivo:", file, ", Erro:", err)
		} else {
			f.Close()
		}
		bar.Add(1)
	}
}

func main() {
	userprofile, _ := os.UserHomeDir()
	filesPath := filepath.Join(userprofile, "Documentos", "PythonVsGo_Dataset", "documentos_excel")
	logPath := filepath.Join(userprofile, "Área de Trabalho", "Projetos", "PythonVsGo-FileManipulation", "PythonVsGo-FileManipulation", "logs")

	startTime := time.Now()

	// Monitoramento de desempenho (início)
	cpuUsageBefore, _ := cpu.Percent(0, false)
	memUsageBefore, _ := mem.VirtualMemory()

	files, _ := filepath.Glob(filepath.Join(filesPath, "*.xlsx"))
	bar := progressbar.Default(int64(len(files)))

	readFiles(filesPath, bar)

	// Monitoramento de desempenho (fim)
	cpuUsageAfter, _ := cpu.Percent(0, false)
	memUsageAfter, _ := mem.VirtualMemory()
	endTime := time.Now()

	logFile, _ := os.Create(filepath.Join(logPath, "read_go_log.txt"))
	defer logFile.Close()

	elapsedTime := endTime.Sub(startTime)
	logFile.WriteString(fmt.Sprintf("Time elapsed: %v\n", elapsedTime))
	logFile.WriteString(fmt.Sprintf("CPU Usage Before: %v%%, After: %v%%\n", cpuUsageBefore[0], cpuUsageAfter[0]))
	logFile.WriteString(fmt.Sprintf("Memory Usage Before: %v MB, After: %v MB\n", memUsageBefore.Used/1024/1024, memUsageAfter.Used/1024/1024))

	fmt.Println("Process finished!")
}
