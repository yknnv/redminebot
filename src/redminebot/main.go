package main

import (
    //"fmt"
    "gopkg.in/telegram-bot-api.v4"
    "log"
    "os"
    "encoding/json"
    "net/http"
    "io/ioutil"
)

type Config struct {
    TelegramBotToken string
}

func getTasks(){
    client := &http.Client{}
    url := "http://redmine.aeroidea.ru/projects/270.json"
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Basic a29ub25vdjpIajEyZE9QTDlEMTI=")
    req.Header.Set("Content-Type", "application/json")
    res, _ := client.Do(req)
    defer res.Body.Close()

    data, _ := ioutil.ReadAll(res.Body)
    //var result myXMLstruct
    //xml.Unmarshal(xmlBytes, &result)
    log.Println(string(data))
    //return string(data)
}

func main() {
    file, _ := os.Open("config.json")
    decoder := json.NewDecoder(file)
    configuration := Config{}
    err := decoder.Decode(&configuration)
    if err != nil {
       log.Panic(err)
    }
    if configuration.TelegramBotToken == ""{
        panic("empty token")    
    }

    bot, err := tgbotapi.NewBotAPI(configuration.TelegramBotToken)

    if err != nil {
        log.Panic(err)
    }

    bot.Debug = false

    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates, err := bot.GetUpdatesChan(u)

    if err != nil {
        log.Panic(err)
    }
    //query := getTasks()
    // В канал updates будут приходить все новые сообщения.
    for update := range updates { 
        // Создав структуру - можно её отправить обратно боту
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
        msg.ReplyToMessageID = update.Message.MessageID
        bot.Send(msg)
        //bot.Send(query)
    }
}


