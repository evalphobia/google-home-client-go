package googlehome

import (
	"context"
	"fmt"
	"net"

	"github.com/barnybug/go-cast"
	"github.com/barnybug/go-cast/controllers"
	"github.com/evalphobia/google-tts-go/googletts"
)

// Client is Google Home client.
type Client struct {
	ctx    context.Context
	lang   string
	accent string
	ip     net.IP
	port   int
}

// NewClient creates Client from environment values.
func NewClient() (*Client, error) {
	return NewClientWithConfig(Config{})
}

// NewClientWithConfig creates Client from given Config.
func NewClientWithConfig(conf Config) (*Client, error) {
	host, err := conf.GetIPv4()
	if err != nil {
		return nil, err
	}
	port := conf.GetPort()
	client := cast.NewClient(host, port)

	ctx := conf.GetOrCreateContext()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return &Client{
		ctx:    ctx,
		ip:     host,
		port:   port,
		lang:   conf.GetLang(),
		accent: conf.GetAccent(),
	}, nil
}

// SetLang sets lang.
func (c *Client) SetLang(lang string) {
	c.lang = lang
}

// SetAccent sets accent.
func (c *Client) SetAccent(accent string) {
	c.accent = accent
}

// GetIPv4 returns IPv4 address of Google Home.
func (c *Client) GetIPv4() string {
	return c.ip.String()
}

// Notify make Google Home say something interesting.
func (c *Client) Notify(text string, language ...string) error {
	lang := c.lang
	if len(language) != 0 {
		lang = language[0]
	}
	if c.accent != "" {
		lang = fmt.Sprintf("%s-%s", lang, c.accent)
	}

	url, err := googletts.GetTTSURL(text, lang)
	if err != nil {
		return err
	}
	return c.Play(url)
}

// Play make Google Home play music or sound.
func (c *Client) Play(url string) error {
	client := cast.NewClient(c.ip, c.port)
	defer client.Close()
	err := client.Connect(c.ctx)
	if err != nil {
		return err
	}
	client.Receiver().QuitApp(c.ctx)

	media, err := client.Media(c.ctx)
	if err != nil {
		return err
	}

	item := controllers.MediaItem{
		ContentId:   url,
		StreamType:  "BUFFERED",
		ContentType: "audio/mpeg",
	}
	_, err = media.LoadMedia(c.ctx, item, 0, true, map[string]interface{}{})
	return err
}

// GetVolume gets volume.
func (c *Client) GetVolume() (volume float64, err error) {
	client := cast.NewClient(c.ip, c.port)
	defer client.Close()
	err = client.Connect(c.ctx)
	if err != nil {
		return 0, err
	}

	vol, err := client.Receiver().GetVolume(c.ctx)
	if err != nil {
		return 0, err
	}

	return *vol.Level, nil
}

// SetVolume sets volume. volume must be 0.0 ~ 1.0.
func (c *Client) SetVolume(volume float64) error {
	client := cast.NewClient(c.ip, c.port)
	defer client.Close()
	err := client.Connect(c.ctx)
	if err != nil {
		return err
	}

	_, err = client.Receiver().SetVolume(c.ctx, &controllers.Volume{Level: &volume})
	return err
}

// QuitApp stops recveiver application.
func (c *Client) QuitApp() error {
	client := cast.NewClient(c.ip, c.port)
	defer client.Close()

	receiver := client.Receiver()
	_, err := receiver.QuitApp(c.ctx)
	return err
}

// StopMedia make Google Home stop music or sound.
func (c *Client) StopMedia() error {
	client := cast.NewClient(c.ip, c.port)
	defer client.Close()

	if !client.IsPlaying(c.ctx) {
		return nil
	}

	media, err := client.Media(c.ctx)
	if err != nil {
		return err
	}
	_, err = media.Stop(c.ctx)
	return err
}
