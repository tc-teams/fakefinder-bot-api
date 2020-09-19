package bot

type webHookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

var emptyString = ""


var startCommand = "Olá, Para consultar a veracidade de uma notícia digite \n /consultar: descrição da notícia"
var consultCommand = "para consultar a veracidade de uma notícia digite\n/consultar: descrição da notícia"
var  defaltCommand = "Não entendi sua mensagem,\npara consultar a veracidade de uma notícia digite\n/consultar: descrição da notícia"