import subprocess
import tkinter as tk
from tkinter import messagebox
import psutil

# Пути к программам
GOODBYE_DPI_PATH = "C:\\PROJECTS\\HorizoN_VPN\\GOODBYEDPI\\x86_64\\goodbyedpi.exe"  # Укажите путь к GoodbyeDPI
INVISIBLE_MAN_PATH = "C:\\PROJECTS\\HorizoN_VPN\\VPN\\InvisibleManXRay.exe"  # Укажите путь к Invisible Man
VPN_LINK_FILE = "C:\\PROJECTS\HorizoN_VPN\\vpn_link.txt"  # Ссылка на VPN (Xray)

# Флаг для отслеживания состояния GoodbyeDPI
gdpi_running = False

# Функция для запуска VPN через Invisible Man
def start_vpn():
    try:
        with open(VPN_LINK_FILE, "r") as f:
            vpn_link = f.read().strip()
        subprocess.Popen([INVISIBLE_MAN_PATH, vpn_link], shell=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
        messagebox.showinfo("VPN", "VPN успешно подключен!")
    except Exception as e:
        messagebox.showerror("Ошибка", f"Не удалось подключиться к VPN: {e}")

# Функция для полной остановки VPN (убивает все процессы)
def stop_vpn():
    stopped = kill_process_by_name("InvisibleManXRay.exe")
    if stopped:
        messagebox.showinfo("VPN", "VPN отключен.")
    else:
        messagebox.showinfo("VPN", "VPN не был запущен.")

# Функция для управления GoodbyeDPI (вкл/выкл)
def toggle_goodbye_dpi():
    global gdpi_running
    if not gdpi_running:
        try:
            subprocess.Popen([GOODBYE_DPI_PATH, "-1", "-f", "2", "-k", "1"], 
                             shell=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
            gdpi_running = True
            messagebox.showinfo("GoodbyeDPI", "GoodbyeDPI запущен!")
            btn_toggle_gdpi.config(text="Остановить GoodbyeDPI")
        except Exception as e:
            messagebox.showerror("Ошибка", f"Не удалось запустить GoodbyeDPI: {e}")
    else:
        stop_goodbye_dpi()

# Функция для полной остановки GoodbyeDPI (убивает все процессы)
def stop_goodbye_dpi():
    global gdpi_running
    stopped = kill_process_by_name("GoodbyeDPI.exe")
    if stopped:
        gdpi_running = False
        messagebox.showinfo("GoodbyeDPI", "GoodbyeDPI остановлен.")
        btn_toggle_gdpi.config(text="Запустить GoodbyeDPI")
    else:
        messagebox.showinfo("GoodbyeDPI", "GoodbyeDPI не был запущен.")

# Утилита для убийства процесса по его имени
def kill_process_by_name(process_name):
    killed = False
    for proc in psutil.process_iter(['pid', 'name']):
        if proc.info['name'].lower() == process_name.lower():
            proc.kill()
            killed = True
    return killed

# Создание GUI
def create_gui():
    root = tk.Tk()
    root.title("VPN и GoodbyeDPI")

    # Кнопки для управления VPN
    btn_start_vpn = tk.Button(root, text="Подключить VPN", command=start_vpn, width=25)
    btn_start_vpn.pack(pady=5)

    btn_stop_vpn = tk.Button(root, text="Отключить VPN", command=stop_vpn, width=25)
    btn_stop_vpn.pack(pady=5)

    # Кнопка для управления GoodbyeDPI (вкл/выкл)
    global btn_toggle_gdpi
    btn_toggle_gdpi = tk.Button(root, text="Запустить GoodbyeDPI", command=toggle_goodbye_dpi, width=25)
    btn_toggle_gdpi.pack(pady=5)

    root.geometry("300x200")
    root.mainloop()

if __name__ == "__main__":
    create_gui()
