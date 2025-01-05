# MiniDLNA Updater

Утилита для автоматического переиндексирования MiniDLNA при изменениях в медиа-директории.

## Установка

1. Скомпилируйте программу для Linux:
```bash
chmod +x build.sh
./build.sh
```

2. Скопируйте исполняемый файл в систему:
```bash
sudo cp minidlna_updater /usr/local/bin/
sudo chmod +x /usr/local/bin/minidlna_updater
```

3. Настройте systemd сервис:
```bash
# Отредактируйте путь к вашей медиа-директории в файле minidlna-updater.service
sudo cp minidlna-updater.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable minidlna-updater
sudo systemctl start minidlna-updater
```

## Использование

Программа автоматически отслеживает изменения в указанной директории и перезапускает службу minidlna при обнаружении изменений.

Для ручного запуска:
```bash
minidlna_updater -dir /path/to/your/media/directory
```

## Проверка статуса

```bash
sudo systemctl status minidlna-updater
```
