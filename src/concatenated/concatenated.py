import os
import time
import pandas as pd
import psutil
import glob
from tqdm import tqdm

def concat_csv_files(directory):
    all_dfs = []
    filepaths = glob.glob(os.path.join(directory, '*.xlsx'))
    for filepath in tqdm(filepaths, desc="Concatenando arquivos"):
        df = pd.read_excel(filepath)
        all_dfs.append(df)
    return pd.concat(all_dfs, ignore_index=True)

def save_to_csv(directory, filename, df):
    filepath = os.path.join(directory, filename)
    df.to_csv(filepath, index=False)  # Save as CSV instead of Excel

def main():
    userprofile = os.path.expanduser('~')
    files_path = os.path.join(userprofile, "Documentos", "PythonVsGo_Dataset", "documentos_excel")
    log_path = os.path.join(userprofile, "Área de Trabalho", "Projetos", "PythonVsGo-FileManipulation", "PythonVsGo-FileManipulation", "logs", "concatenated_python_log.txt")

    start_time = time.time()
    cpu_usage_before = psutil.cpu_percent(interval=1)
    mem_usage_before = psutil.virtual_memory().percent

    concatenated_df = concat_csv_files(files_path)
    save_to_csv(files_path, "concatenated_results.csv", concatenated_df)  # Save as CSV

    cpu_usage_after = psutil.cpu_percent(interval=1)
    mem_usage_after = psutil.virtual_memory().percent
    end_time = time.time()
    elapsed_time = end_time - start_time

    log_data = f"Time elapsed: {elapsed_time} seconds\nCPU Usage Before: {cpu_usage_before}%, After: {cpu_usage_after}%\nMemory Usage Before: {mem_usage_before}%, After: {mem_usage_after}%\n"
    
    with open(log_path, 'w') as log_file:
        log_file.write(log_data)

    print("Process finished!", log_data)

if __name__ == "__main__":
    main()