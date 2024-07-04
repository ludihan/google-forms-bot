Does your organization force you to fill tons of meaningless forms like a caveman?
I'm pretty sure they do!
Hit'em with some automation!
(this doesnt work with forms that send a copy of themselves to the user's email
 or ones that have Gmail authentication, probably because of
 [this](https://blog.talosintelligence.com/google-forms-quiz-spam/))

# How to compile this tool
1. Clone this repository
2. Install [Go](https://go.dev/)
3. Execute this inside the root directory of this repo:
```sh
go build bot.go
```

# How to use this tool

(this tutorial assumes you're using firefox, but you can use any browser you want)
1. Fill out the form
2. Open devtools(CTRL+SHIFT+I) and select the "Network" panel
3. Find a "POST" request containing a file called "formResponse"
4. Copy the url
5. Under the "Request" tab you'll find some keys like "entry.*", this is what you've send to the form
6. Write the keys and the data you'll want to send to a csv file ([example](example.csv))
7. Execute this (substitute the "\<url\>" with your actual target):
```sh
./bot file.csv <url>
```
VERY IMPORTANT: if you want fill out an email in the form put the key "emailAddress"
    in the csv
