FROM alpine

#RUN apk update && apk add curl && apk add python3 && apk add ffmpeg
RUN apk update && apk add curl && apk add python3

RUN curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -o /usr/local/bin/youtube-dl && \
    chmod a+rx /usr/local/bin/youtube-dl

ADD app .
ADD config.yaml .

RUN chmod +x ./app

EXPOSE 9990

ENTRYPOINT ["/app"]