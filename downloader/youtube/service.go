package youtube

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/andres-website/musicload-bot/proxy_config"
	"github.com/pkg/errors"

	"os/exec"

	"log"

	"net/url"
)

const (
	expression = "^(http(s)?:\\/\\/)?((w){3}.)?(music\\.)?youtu(be|.be)?(\\.com)?\\/.+"
)

type Downloader struct {
	maxVideoDuration time.Duration

	r *regexp.Regexp
}

func NewDownloader(maxVideoDuration int64) (*Downloader, error) {
	r, err := regexp.Compile(expression)
	if err != nil {
		return nil, err
	}

	return &Downloader{
		maxVideoDuration: time.Minute * time.Duration(maxVideoDuration),
		r:                r,
	}, nil
}

func (d *Downloader) Download(ctx context.Context, url string) (string, error) {

	// log.Println("url: " + url)
	// info, err := ytdl.GetVideoInfo(ctx, url)
	// if err != nil {
	// return "", errors.Wrap(err, "error getting video info")
	//}

	// if info.Duration > d.maxVideoDuration {
	// return "", downloader.ErrDurationLimitExceeded
	// }

	// filename := info.Title
	// filename := "title"
	// Вызов функции для получения значения параметра w
	id_youtube, err := Get_id_youtube(url)
	if err != nil {
		fmt.Println(err)
		return "", errors.New(fmt.Sprintf("Ошибка парсинага url Для полученя id видео: %s", string(url)))
	}
	filename := id_youtube
	log.Printf("id_youtube: " + id_youtube)
	// strings.Replace(filename, " ", "\\ ", 0)

	// this command downloads video and extracts webm
	// youtube-dl -x --audio-format webm https://www.youtube.com/watch?v=dQw4w9WgXcQ -o "Rick Astley - Never Gonna Give You Up.webm
	// cmd := exec.Command("youtube-dl", "-а, "--audio-format", "webm", "-o", outputDir+"/%(title)s.%(ext)s", url, "--user-agent", "Mozilla/5.0 (Android 14; Mobile; rv:128.0) Gecko/128.0 Firefox/128.0")

	log.Printf("Start_2")

	// cmd := exec.CommandContext(ctx, "youtube-dl", "-x", "--audio-format", "webm", url, "--throttled-rate", "500K", "-o", filename+".%(ext)s", "--user-agent", "Mozilla/5.0 (Android 14; Mobile; rv:128.0) Gecko/128.0 Firefox/128.0", "--proxy", "http://user:pass@5.35.103.157:8888")
	// cmd := exec.CommandContext(ctx, "youtube-dl", "-f", "bestaudio[ext=webm]", url, "--throttled-rate", "500K", "-o", filename+".%(ext)s", "--user-agent", "Mozilla/5.0 (Android 14; Mobile; rv:128.0) Gecko/128.0 Firefox/128.0", "--proxy", "http://username:p1a2s3s4@5.35.103.157:1742")
	// used on unsused proxy
	cmdArgs := []string{
		"-f", "bestaudio[ext=webm]",
		url,
		"--throttled-rate", "500K",
		"-o", filename + ".%(ext)s",
		"--user-agent", "Mozilla/5.0 (Android 14; Mobile; rv:128.0) Gecko/128.0 Firefox/128.0",
		// "--proxy", "http://user:pass@0.0.0.0:8888",
	}

	// from Global varible: proxy_config.AppConfig
	use_proxy := proxy_config.AppConfig.UseProxy
	// Добавляем параметр прокси, если use_proxy установлен
	if use_proxy {
		proxy := proxy_config.AppConfig.Proxy
		cmdArgs = append(cmdArgs, "--proxy", proxy)
	}

	// Создаём команду с общими и дополнительными аргументами
	cmd := exec.CommandContext(ctx, "youtube-dl", cmdArgs...)

	data, err := cmd.CombinedOutput()
	if err != nil {
		os.Remove(filename)
		return "", errors.Wrap(err, fmt.Sprintf("error from CombinedOutput, data: %s", string(data)))
	}

	if strings.Contains(string(data), "ERROR") {
		os.Remove(filename)
		return "", errors.New(fmt.Sprintf("error downloading video with youtube-dl, output: %s", string(data)))
	}

	// Переименование файла с .webm на .mp3
	newFilename := filename + ".mp3"
	oldFilename := filename + ".webm"
	if err := os.Rename(oldFilename, newFilename); err != nil {
		return "", errors.Wrap(err, "не удалось переименовать файл")
	}

	return filename + ".mp3", nil
}

func (d *Downloader) IsValidURL(url string) bool {
	return d.r.MatchString(url)
}

// GetWValue извлекает значение параметра "w" из переданного URL.
func Get_id_youtube(urlStr string) (string, error) {

	// Парсинг URL
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("ошибка при парсинге URL: %v", err)
	}

	// Получение параметров запроса
	queryParams := parsedUrl.Query()

	// Извлечение значения параметра w
	wValue := queryParams.Get("v")
	return wValue, nil
}
