# PortScanGo

CLI-инструмент для TCP и UDP сканирования портов, написанный на Go.

## Установка

```bash
git clone https://github.com/Andrei9460/portscango
cd portscango
go mod init portscango
go build -o portscan main.go
```

## Примеры использования

```
./portscan -host 127.0.0.1 -range 20-80 -tcp
./portscan -host 127.0.0.1 -range 53-53 -udp
./portscan -host 127.0.0.1 -range 1-9999 -tcp -udp
```
<img width="796" height="119" alt="image" src="https://github.com/user-attachments/assets/1da8b7eb-fc98-43ff-bb6f-1bbe137879c9" />



| Аргумент | Описание                            |
| -------- | ----------------------------------- |
| `-host`  | IP-адрес или домен для сканирования |
| `-range` | Диапазон портов, например `1-1000`  |
| `-tcp`   | Включить TCP-сканирование           |
| `-udp`   | Включить UDP-сканирование           |
