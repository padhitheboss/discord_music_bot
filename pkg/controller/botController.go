package controller

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"github.com/padhitheboss/sangeet/pkg/config"
	"github.com/padhitheboss/sangeet/pkg/model"
)

// var PlayFromQueueRunning = false

func getAudioUrl(videoUrl string) string {
	path, _ := os.Getwd()
	cmd := exec.Command(path+"/yt-dlp_linux", "ytsearch:bestaudio", "--get-url", videoUrl)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	audioUrls := strings.Split(string(output), "\n")
	return audioUrls[len(audioUrls)-2]
}

func playFromQueue(s *discordgo.Session, voiceConnection *discordgo.VoiceConnection) {
	model.PlayFromQueueRunning = true
	defer func() {
		// err := voiceConnection.Disconnect()
		// if err != nil {
		// 	log.Println("already disconnected")
		// }
		model.PlayFromQueueRunning = false
		fmt.Println("Exited Function")
	}()
	for {
		if model.P.Playing && len(model.Q.Playlist) > 0 {
			audioURL := getAudioUrl(model.Q.Playlist[0].Url)
			TitlePlaying := model.Q.Playlist[0].Title
			if len(model.Q.Playlist) > 0 {
				model.Q.Lock.Lock()
				model.Q.Playlist = model.Q.Playlist[1:]
				model.Q.Lock.Unlock()
			}
			fmt.Println("Playing:", TitlePlaying)
			dgvoice.PlayAudioFile(voiceConnection, audioURL, model.P.RunChan)
		} else {
			// voiceConnection.Disconnect()
			model.PlayFromQueueRunning = false
			break
		}
	}
}

func getUserVoiceChannelID(s *discordgo.Session, userID string) (string, error) {
	user, err := s.User(userID)
	if err != nil {
		return "", err
	}

	for _, guild := range s.State.Guilds {
		for _, vs := range guild.VoiceStates {
			if vs.UserID == user.ID {
				return vs.ChannelID, nil
			}
		}
	}

	return "", nil
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Content == "" {
		return
	}
	var b model.BotCommand
	args := strings.Fields(m.Content)
	b.Command = args[0]
	args = args[1:]
	b.Query = strings.Join(args, " ")
	b.User = m.Author.Username
	response := b.Do()
	s.ChannelMessageSend(m.ChannelID, response)
	voiceChannelID, err := getUserVoiceChannelID(s, m.Author.ID)
	if err != nil {
		fmt.Println("Failed to get the voice channel ID:", err)
		return
	}
	if model.P.Playing {
		voiceConnection, err := s.ChannelVoiceJoin(config.DiscordServerID, voiceChannelID, false, false)
		if err != nil {
			fmt.Println("Failed to join the voice channel:", err)
			return
		}
		if !model.PlayFromQueueRunning {
			go playFromQueue(s, voiceConnection)
		}
	}
}
