package service

//const offset = 10

/*type CommandData struct {
	Offset int
	Page int
}*/





/*func RenderProductListWithMarkup(currentPage int,products []models.Product)(string,tgbotapi.InlineKeyboardMarkup){
	maxPages:=len(products)/offset
	markUp:=theFirstPageMarkup
	text:=generateTextForBotMessage(products[currentPage*offset:currentPage*offset+offset])
	if currentPage>0&&currentPage<maxPages-1{
		markUp=middlePageMarkup
	}
	if currentPage==maxPages-1{
		markUp=lastPageMarkup
	}
	return text,markUp
}*/

/*func SendListMessagesIntoChat(products []models.Product,chatId int64,currentPage int,messageId *int)tgbotapi.Chattable{
	text,markUp:=RenderProductListWithMarkup(currentPage,products)
	var cfg tgbotapi.Chattable
	if messageId==nil{
		msg:=tgbotapi.NewMessage(chatId,text)
		msg.ReplyMarkup=markUp
		cfg=msg
	} else {
		msg:=tgbotapi.NewEditMessageText(chatId,*messageId,text)
		msg.ReplyMarkup=&markUp
		cfg=msg
	}
	return cfg
}

func HandleNavigationCallbackQuery(messageId int,data ...string){
	pagerType:=data[0]
	currentPage,_:=strconv.Atoi(data[1])

	if pagerType=="next"{
		nextPage:=currentPage+1
		if nextPage <{

		}
	}
}*/