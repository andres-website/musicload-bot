package bot

import (
	"github.com/andres-website/musicload-bot/proxy_config"
	"github.com/andres-website/musicload-bot/youtube_api"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func (t *TelegramBot) handleUpdates(update tgbotapi.Update) {

	if m := update.Message; m != nil {

		// Обработка команды /start
		if m.IsCommand() && m.Command() == "start" {
			t.send(m.Chat.ID, "Hi there! Send me a link to a video you want extract music from.")
			return
		}
		//

		// Обработка сообщения URL ютуба - если это URL ютуба
		if t.downloadService.IsValidURL(m.Text) {
			t.queue.Enqueue(m)
			return
		}
		//

		// Обработка текстового сообщения (и поиск в Ютуб апи подходящего видео)
		use_youtube_api := proxy_config.AppConfig.Use_youtube_api
		if use_youtube_api {
			apiKey := proxy_config.AppConfig.Youtube_api_key

			videoID, title, err := youtube_api.GetYoutubeVideoId_and_title(apiKey, m.Text)
			if err != nil {
				t.send(m.Chat.ID, "Не удалось получить ID ютуб видео по запросу: "+m.Text)
				return
			}

			url_video_by_recive_id := "https://www.youtube.com/watch?v=" + videoID
			m.Text = url_video_by_recive_id
			msg := title + " " + url_video_by_recive_id

			t.send(m.Chat.ID, msg)

			if t.downloadService.IsValidURL(url_video_by_recive_id) {

				t.queue.Enqueue(m)
				return
			}
		}
		//

		t.send(m.Chat.ID, "Invalid message text. I'm waiting for youtube link.")
	}
}
