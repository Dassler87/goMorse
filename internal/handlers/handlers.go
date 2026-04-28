package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

const outputDir = "./new_files/" // Указываем относительный путь к новой папке

// ServeIndex возвращает HTML‑страницу index.html
func ServeIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Определяем абсолютный путь к исполняемому файлу
	executablePath, err := os.Executable()
	if err != nil {
		log.Printf("Ошибка определения пути к исполняемому файлу: %v", err)
		http.Error(w, "Ошибка определения пути к файлу", http.StatusInternalServerError)
		return
	}
	// Строим путь к файлу index.html относительно корня проекта
	rootDirectory := filepath.Dir(executablePath)
	indexHTMLPath := filepath.Join(rootDirectory, "..", "index.html")

	// Подготавливаем заголовки ответа
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Отправляем файл клиенту
	http.ServeFile(w, r, indexHTMLPath)
}

// UploadHandler обрабатывает загрузку файла, конвертирует его содержимое и сохраняет результат
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Парсим форму с файлом (ограничение 10 МБ)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Printf("Ошибка парсинга формы: %v", err)
		http.Error(w, "Ошибка парсинга формы", http.StatusInternalServerError)
		return
	}

	// Получение файла из формы
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Printf("Ошибка получения файла: %v", err)
		http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Чтение содержимого файла
	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Ошибка чтения файла: %v", err)
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	input := string(data)

	// Конвертируем данные
	converted, err := service.Convert(input)
	if err != nil {
		log.Printf("Ошибка конвертации: %v", err)
		http.Error(w, "Ошибка конвертации", http.StatusInternalServerError)
		return
	}

	// Генерируем имя файла с текущим временем и расширением исходного файла
	ext := filepath.Ext(handler.Filename)
	timeFormat := time.Now().UTC().Format("20060102_150405")
	filename := timeFormat + ext

	// Полный путь к новому файлу
	fullPath := outputDir + filename

	// Создаем директорию, если она не существует
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Printf("Ошибка создания директории: %v", err)
		http.Error(w, "Ошибка создания директории", http.StatusInternalServerError)
		return
	}

	// Создаём и записываем файл с результатом
	outFile, err := os.Create(fullPath)
	if err != nil {
		log.Printf("Ошибка создания файла: %v", err)
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = outFile.WriteString(converted)
	if err != nil {
		log.Printf("Ошибка записи в файл: %v", err)
		http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат пользователю
	w.Write([]byte(converted))
}
