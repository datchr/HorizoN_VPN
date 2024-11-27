import os
import subprocess
import tkinter as tk
from tkinter import messagebox

# Пути к файлам
XRAY_EXECUTABLE = r"C:\HORIZON_TEST\Xray-Core\xray.exe"
CONFIG_FILE = r"C:\HORIZON_TEST\Configs\config.json"
LOG_FILE = r"C:\HORIZON_TEST\Logs\error.log"

def start_vpn():
    if not os.path.exists(XRAY_EXECUTABLE):
        messagebox.showerror("Ошибка", f"Файл {XRAY_EXECUTABLE} не найден.")
        return

    if not os.path.exists(CONFIG_FILE):
        messagebox.showerror("Ошибка", f"Конфиг {CONFIG_FILE} не найден.")
        return

    try:
        # Запуск Xray
        process = subprocess.Popen(
            [XRAY_EXECUTABLE, "-c", CONFIG_FILE],
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE
        )
        messagebox.showinfo("VPN", "VPN успешно запущен!")
    except Exception as e:
        messagebox.showerror("Ошибка", f"Не удалось запустить VPN: {e}")

def stop_vpn():
    # Завершение работы Xray (если потребуется, добавить механизм завершения)
    messagebox.showinfo("VPN", "Остановка VPN в разработке...")

# Интерфейс
root = tk.Tk()
root.title("HORIZON VPN")

frame = tk.Frame(root, padx=20, pady=20)
frame.pack()

start_button = tk.Button(frame, text="Запустить VPN", command=start_vpn, width=20, bg="green", fg="white")
start_button.grid(row=0, column=0, padx=10, pady=5)

stop_button = tk.Button(frame, text="Остановить VPN", command=stop_vpn, width=20, bg="red", fg="white")
stop_button.grid(row=1, column=0, padx=10, pady=5)

root.mainloop()
