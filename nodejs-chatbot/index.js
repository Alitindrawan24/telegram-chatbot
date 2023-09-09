const TelegramBot = require('node-telegram-bot-api');
require("dotenv").config();
const token = process.env.TELEGRAM_CHATBOT_TOKEN

// Validate the token is not empty
if (token === undefined || token === null || token === "") {
    console.log("telegram bot token is required")
    process.exit(1)
}

// Initialize the TelegramBot
const bot = new TelegramBot(token, { polling: true });
bot.getMe().then(function(info) {
    console.log(`Authorized on account ${info.username}`);
});


// Listen for incoming messages
bot.on('message', (msg) => {
    const chatID = msg.chat.id;
    const text = msg.text;
    const messageID = msg.message_id;

    if (text !== undefined) {
        bot.sendMessage(chatID, text, {
            reply_to_message_id: messageID
        });
    }
});