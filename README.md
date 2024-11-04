# MusicLoadBot

Simple Telegram Bot written in Go that downloads mp3 from Youtube videos.

![bot-screenshot](https://user-images.githubusercontent.com/24574014/54880761-5a8cb880-4e51-11e9-8800-699156243b8c.png)

# Stack

1. `Golang`
2. `Docker`

# Installation 


0. Rename config.yaml.example -> config.yaml
0.1 If you from Russua: use Proxy (example: tinyproxy or other)
0.2 Change use_youtube_api to true (and insert your youtube_api_key) in config.yaml to enable the ability to search Youtube videos by free request.
1. Set your bot's token, username and other settings in config.yaml, than execute:
2. `make build`
3. `make run`
