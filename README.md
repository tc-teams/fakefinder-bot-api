# fakefinder-bot-api
This API aims the interaction of our Telegram bot `@fake_finder_bot` with the other API's of the project.

## Installation

```bash

```

## Usage

```docker

docker build -t gcr.io/model-framing-272522/fakefinder-bot-api:1.0.0 .
docker run -e TELEGRAM_BOT_KEY=$TELEGRAM_BOT_KEY -d --name bot -t  gcr.io/model-framing-272522/fakefinder-bot-api:1.0.0 
docker logs bot --follow

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## References
```
openssl req -newkey rsa:2048 -sha256 -nodes -keyout YOURPRIVATE.key -x509 -days 365 -out YOURPUBLIC.pem -subj "/C=US/ST=New York/L=Brooklyn/O=Example Brooklyn Company/CN={{hostname}}"

curl -sL -F "url=https://{{hostname}}/telegram/bot/api" -F "certificate=@YOURCERTIFICATE.pem" https://api.telegram.org/bot1228162506:AAFQr_ipJ3dsaOieJFupA5Pw4BlhuRFoOyE/setWebhook
```

### Go Telegram bot API
- wiki: https://github.com/go-telegram-bot-api/telegram-bot-api/wiki
- Documentation: https://go-telegram-bot-api.dev/

## License
