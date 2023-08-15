# Sangeet (A Discord Music Bot)

Sangeet is a discord music bot written in Golang. It searches the song title it youtube using youtubev3 API and gets the URL of the youtube video and all details like duration and full title etc. Then it's using yt-dlp binary to get the audio link and using ffmpeg to play audio from the link.

# How to Install and Run
- clone the repository onto the local machine
- Rename the .env example to .env
- Populate .env file with Youtube and Discord API Key and Discord Server ID
- Now you can build or run the main file directly. 

# Supported Commands

- !play <song name> or !p <song name> -- play music or add to queue if playing
- !skip -- skip the current playing song
- !queue or !q -- List the song in the queue
- !clear or !c -- clear queue
- !resume - After stopping resume from the next song in the playlist
- !stop -- stop music 
- !playnext <song name> -- play the song after finishing the current song

# Note - To Play Music User Needs to DM the Bot

# Comming up in next commits -

- Some better approach to get a Youtube link instead of depending on yt-dlp binary
- Instead of DM to Bot User will be able to send a command to a channel
- Multi Discord Server Support
- Documentation For Developer
  

