package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// Определение флагов командной строки
	watchDir := flag.String("dir", "", "Directory to watch")
	flag.Parse()

	if *watchDir == "" {
		log.Fatal("Please specify directory to watch using -dir flag")
	}

	// Создание нового watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Создаем канал для дебаунсинга
	debounce := make(chan bool)
	var timer *time.Timer

	// Горутина для обработки событий
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// Игнорируем временные файлы
				if filepath.Ext(event.Name) == ".tmp" {
					continue
				}
				
				log.Printf("Event detected: %s", event.String())
				
				// Сбрасываем таймер если он существует
				if timer != nil {
					timer.Stop()
				}
				
				// Создаем новый таймер
				timer = time.NewTimer(5 * time.Second)
				go func() {
					<-timer.C
					debounce <- true
				}()

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Printf("Error: %s", err)
			case <-debounce:
				log.Println("Reindexing minidlna...")
				if err := reindexMinidlna(); err != nil {
					log.Printf("Error reindexing minidlna: %s", err)
				}
			}
		}
	}()

	// Рекурсивное добавление всех поддиректорий для отслеживания
	err = filepath.Walk(*watchDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			err = watcher.Add(path)
			if err != nil {
				log.Printf("Error watching path %s: %s", path, err)
				return err
			}
			log.Printf("Watching directory: %s", path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Started watching directory: %s", *watchDir)
	// Бесконечный цикл
	<-make(chan struct{})
}

func reindexMinidlna() error {
	// Перезапускаем службу minidlna
	cmd := exec.Command("minidlnad", "-r")
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Println("minidlna service restarted successfully")
	return nil
}
