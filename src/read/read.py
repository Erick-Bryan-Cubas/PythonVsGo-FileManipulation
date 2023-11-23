import os
import sys
import time
from tqdm import tqdm
import psutil
from openpyxl import load_workbook

def read_files(directory, bar):
    files = os.listdir(directory)
    for file in files:
        if file.endswith(".xlsx"):
            file_path = os.path.join(directory, file)
            try:
                workbook = load_workbook(filename=file_path, read_only=True)
                workbook.close()
            except Exception as e:
                print(f"Erro ao ler o arquivo: {file}, Erro: {e}")
            bar.update(1)

def main():
    userprofile = os.environ['USERPROFILE']
    files_path = os.path.join(userprofile, "Documentos", "PythonVsGo_Dataset", "documentos_excel")
    log_path = os.path.join(userprofile, "Área de Trabalho", "Projetos", "PythonVsGo-FileManipulation", "PythonVsGo-FileManipulation", "logs")

    start_time = time.time()
    cpu_usage_before = psutil.cpu_percent()
    memory_usage_before = psutil.virtual_memory().percent

    files = os.listdir(files_path)
    with tqdm(total=len(files), desc="Processando arquivos") as pbar:
        read_files(files_path, pbar)

    cpu_usage_after = psutil.cpu_percent()
    memory_usage_after = psutil.virtual_memory().percent
    end_time = time.time()
    elapsed_time = end_time - start_time

    with open(os.path.join(log_path, "read_python_log.txt"), "w") as log_file:
        log_file.write(f"Time elapsed: {elapsed_time} seconds\n")
        log_file.write(f"CPU Usage Before: {cpu_usage_before}%, After: {cpu_usage_after}%\n")
        log_file.write(f"Memory Usage Before: {memory_usage_before}%, After: {memory_usage_after}%\n")

    print("Process finished!")

if __name__ == "__main__":
    main()
