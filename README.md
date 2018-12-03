# Coffeeify

#Overview
The main aim behind our application is to recommend a coffee drink to the users according to the current weather in their city. Our application aims
to enhance the coffee experience users coffee experience. At the same time the user is able to view different coffee recepies that were posted by different users.

Services Offered:
. Insert a coffee drink
. View the recommended drink of the day according to the current Weather
. View the full recepie of the drink
. View all the existing recepies on the Website


Note: This applicaion is written in GoLang and the two backing services used are MYSQL and OpenWeather API

#Getting Started

#Requirements 
.Docker 17.09.0-ce
.Docker-Compose 2.1

#Installing

```git clone https://github.com/OmarShaker97/Coffeeify.git```

#Dependencies

Dependencies that were declared and isolated in our web-application was openweathermap, gorilla mux, securecookie, and mysql.

In our application, we are using govendor which is used for dependency declaration, and isolation which automatically vendors all of the dependencies delcared in our vendor file.

If you need change some source you can deploy it typing:

```govendor build .```

#Database
Run the following command first in order to setup the database.
``` ```

#Starting Services

```docker-compose up -d --build```

#Stopping Services

```docker-compose down```

#Include new changes

If you need change some source you can deploy it typing:

```docker-compose build```

