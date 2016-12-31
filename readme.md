# yaml-manager
## A purpose-built tool to manage i18n yaml files that contain text portions.

### Problem
Your app has 20+ supported languages. You have various bits of text stored 
in a yaml file for each language: `en.yml`, `de.yml`, `he.yml` for english, dutch, and hebrew, for example.
You create a feature for your app that has some new buttons, and the buttons have new text on them.
Before you put the feature into production, you need to google translate the new text portions for each supported language - 
or maybe use english across the board since you're time constrained - and then update the yaml file of each language.
This will take you an hour or more, and you might make a mistake that slips past review and is encountered by your users.

It's possible that you've been in production for years, and you've been developing and deleting enough features that you
have unused cruft in your locale files. De-crufting your i18n yaml files will never be a priority, but it probably annoys 
you to some degree.

### Solution
This tool automates the process of de-crufting and updating your yaml files. Specify one yaml file, say `en.yml`, that is the
master version, and for every other `*.yml` file, it will add the new keys, erase keys not existing in `en.yml`, while not 
overwriting the localized bits for keys that do still exist in `en.yml`. In addition, you can add comments and whitespace 
to the master file, and those comments will be retained in the file you update.

### Live Demo: [coming soon](http://www.louisritchie.com)

### You really don't need a whole framework to do this.
Correct! This is a portfolio project to demonstrate my understanding of the Go language and the React frontend framework. 
While it is really excessive, it shows my competence with Go and React.
The task that this tool is built for can be done with a bash script and a python script. In fact, I have implemented this 
once before with shell and ruby.
