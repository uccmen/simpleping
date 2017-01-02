package simpleping

import (
	"math/big"

	"github.com/uccmen/redisutil"
)

var Subcribers []string
var RedisInstance *redisutil.RedisInstance

// IncomingMessage represents incoming message sent by user on Messenger
type IncomingMessage struct {
	Object  string   `json:"object,omitempty"`
	Entries *[]Entry `json:"entry,omitempty"`
}

// OutgoingMessage represents outgoing message to be sent back to user on Messenger
type OutgoingMessage struct {
	Recipient Recipient            `json:"recipient,omitempty"`
	Message   *OutgoingMessageData `json:"message,omitempty"`

	SenderAction string `json:"sender_action,omitempty"`
}

// OutgoingMessageData represents outgoing message data to be sent back to user on Messenger
type OutgoingMessageData struct {
	Text       string                  `json:"text,omitempty"`
	Attachment *OutgoingAttachmentData `json:"attachment,omitempty"`
}

// OutgoingAttachmentData represents outgoing attachment data to be sent back to user on Messenger
type OutgoingAttachmentData struct {
	Type    string   `json:"type,omitempty"`
	Payload *Payload `json:"payload,omitempty"`
}

// Element represents element of a message on Messenger
type Element struct {
	Title    string   `json:"title,omitempty"`
	Subtitle string   `json:"subtitle,omitempty"`
	ItemURL  string   `json:"item_url,,omitempty"`
	ImageURL string   `json:"image_url,omitempty"`
	Buttons  []Button `json:"buttons,omitempty"`
}

// Button represents button of a message on Messenger
type Button struct {
	Type    string `json:"type,omitempty"`
	URL     string `json:"url,omitempty"`
	Title   string `json:"title,omitempty"`
	Payload string `json:"payload,omitempty"`
}

// Entry represents each message entry sent by user on Messenger
type Entry struct {
	ID        string    `json:"id,omitempty"`
	Time      big.Int   `json:"time,omitempty"`
	Messaging []Message `json:"messaging,omitempty"`
}

// Message represents message exchanged on Messenger
type Message struct {
	Sender      Sender       `json:"sender,omitempty"`
	Recipient   Recipient    `json:"recipient,omitempty"`
	Timestamp   big.Int      `json:"timestamp,omitempty"`
	MessageData MessageData  `json:"message,omitempty"`
	Read        ReadData     `json:"read,omitempty"`
	Delivery    DeliveryData `json:"delivery,omitempty"`
	Postback    *Postback    `json:"postback,omitempty"`
}

// Postback represents postback data sent by user on Messenger
type Postback struct {
	Payload string `json:"payload,omitempty"`
}

// DeliveryData represents delivery data of a message exchanged on Messenger
type DeliveryData struct {
	MIDs      []string `json:"mids,omitempty"`
	Watermark big.Int  `json:"watermark,omitempty"`
	Sequence  int      `json:"seq,omitempty"`
}

// ReadData represents read data of a message exchanged on Messenger
type ReadData struct {
	Watermark big.Int `json:"watermark,omitempty"`
	Sequence  int     `json:"seq,omitempty"`
}

// Sender represents sender data of a message exchanged on Messenger
type Sender struct {
	ID string `json:"id,omitempty"`
}

// Recipient represents recipient data of a message exchanged on Messenger
type Recipient struct {
	ID string `json:"id,omitempty"`
}

// MessageData represents message data of a message exchanged on Messenger
type MessageData struct {
	IsEcho      bool          `json:"is_echo,omitempty"`
	AppID       big.Int       `json:"app_id,omitempty"`
	MID         string        `json:"mid,omitempty"`
	Sequence    int           `json:"seq,omitempty"`
	StickerID   *big.Int      `json:"sticker_id,omitempty"`
	Text        string        `json:"text,omitempty"`
	Attachments *[]Attachment `json:"attachments,omitempty"`
}

// Attachment represents attachment data of a message exchanged on Messenger
type Attachment struct {
	Type    string  `json:"type,omitempty"`
	Payload Payload `json:"payload,omitempty"`
}

// Payload represents payload data of a message exchanged on Messenger
type Payload struct {
	TemplateType string     `json:"template_type,omitempty"`
	Elements     *[]Element `json:"elements,omitempty"`
	URL          string     `json:"url,omitempty"`
	StickerID    *big.Int   `json:"sticker_id,omitempty"`
}

// GetStartedTemplate Button initialization
type GetStartedTemplate struct {
	SettingType   string `json:"setting_type,omitempty"`
	ThreadState   string `json:"thread_state,omitempty"`
	CallToActions []CTA  `json:"call_to_actions,omitempty"`
}

// CTA represents call to action payload information of a message
type CTA struct {
	Payload string `json:"payload,omitempty"`
}
