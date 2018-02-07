package googlehome

import (
	"context"
	"fmt"
	"net"
	"net/url"

	"github.com/barnybug/go-cast"
	"github.com/barnybug/go-cast/controllers"
)

const (
	defaultAPIURL = "https://translate.google.com/translate_tts"
)

// Client is Google Home client.
type Client struct {
	ctx    context.Context
	client *cast.Client
	lang   string
	accent string
	ip     net.IP
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
		client: client,
		ip:     host,
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

	params := &url.Values{}
	params.Set("ie", "UTF-8")
	params.Set("client", "google-home")
	params.Set("tl", lang)
	params.Set("q", text)

	url := fmt.Sprintf("%s?%s", defaultAPIURL, params.Encode())
	return c.Play(url)
}

// Play make Google Home play music or sound.
func (c *Client) Play(url string) error {
	media, err := c.client.Media(c.ctx)
	if err != nil {
		return err
	}
	contentType := "audio/mpeg"
	item := controllers.MediaItem{
		ContentId:   url,
		StreamType:  "BUFFERED",
		ContentType: contentType,
	}
	_, err = media.LoadMedia(c.ctx, item, 0, true, map[string]interface{}{})
	return err
}

// StopMedia make Google Home stop music or sound.
func (c *Client) StopMedia() error {
	if !c.client.IsPlaying(c.ctx) {
		return nil
	}
	media, err := c.client.Media(c.ctx)
	if err != nil {
		return err
	}
	_, err = media.Stop(c.ctx)
	return err
}

// Close closes Google Home client.
func (c *Client) Close() {
	c.client.Close()
}
