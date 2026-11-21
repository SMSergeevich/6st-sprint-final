package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// Index отдаёт HTML-форму из index.html.
func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// Upload парсит HTML-форму, читает загруженный файл,
// конвертирует его содержимое и сохраняет результат в локальный файл.
func Upload(w http.ResponseWriter, r *http.Request) {
	// 1. Парсим multipart-форму (ограничим память 10 МБ).
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. Получаем файл из формы.
	// Имя поля "file" должно совпадать с name="file" в index.html.
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// 3. Читаем данные из файла.
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Передаём данные в service для автоопределения и конвертации.
	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 5. Создаём локальный файл:
	//    имя — на основе текущего времени UTC,
	//    расширение — как у исходного загруженного файла.
	name := time.Now().UTC().String() + filepath.Ext(header.Filename)
	out, err := os.Create(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// 6. Записываем результат конвертации в локальный файл.
	if _, err := out.WriteString(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 7. Возвращаем результат конвертации клиенту.
	fmt.Fprint(w, result)
}
