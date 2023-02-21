# DailyEmoji

😄 DailyEmoji is a web app that displays a different emoji every day! 🍪

### How does it work?

#### Endpoints
The app has some endpoints:
* ``/``: The homepage with the emoji of the day!
* ``/about``: Page with some information about the project.
* ``/api/emoji``: Get the current emoji in JSON format.

#### Selecting emojis
* 🎲 The website selects every day a random emoji from the given emoji list (see emoji/emoji.go for this list).
Not all the emojis are in that list because some emojis didn't render that well in some browers, but the app has quite enough emojis to work with.
* When starting the app, the list of emojis will be shuffled based on a given seed (you can set a seed in emoji/emoji.go) 
and the current year. 
* An emoji will be taken out of the list based on the day of the year. Therefore, the year is added
to the seed to give a truly random emoji, every single year :)


Have a nice day! 👋

Used getemoji.com for the emoji list, thanks ♥️!