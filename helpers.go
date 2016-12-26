package main

import "math/big"

type IncomingMessage struct {
	Object  string   `json:"object,omitempty"`
	Entries *[]Entry `json:"entry,omitempty"`
}

type OutgoingMessage struct {
	Recipient Recipient            `json:"recipient,omitempty"`
	Message   *OutgoingMessageData `json:"message,omitempty"`

	SenderAction string `json:"sender_action,omitempty"`
}

type OutgoingMessageData struct {
	Text       string                  `json:"text,omitempty"`
	Attachment *OutgoingAttachmentData `json:"attachment,omitempty"`
}

type OutgoingAttachmentData struct {
	Type    string   `json:"type,omitempty"`
	Payload *Payload `json:"payload,omitempty"`
}

type Element struct {
	Title    string   `json:"title,omitempty"`
	Subtitle string   `json:"subtitle,omitempty"`
	ItemUrl  string   `json:"item_url,,omitempty"`
	ImageUrl string   `json:"image_url,omitempty"`
	Buttons  []Button `json:"buttons,omitempty"`
}

type Button struct {
	Type    string `json:"type,omitempty"`
	Url     string `json:"url,omitempty"`
	Title   string `json:"title,omitempty"`
	Payload string `json:"payload,omitempty"`
}

type Entry struct {
	ID        string    `json:"id,omitempty"`
	Time      big.Int   `json:"time,omitempty"`
	Messaging []Message `json:"messaging,omitempty"`
}

type Message struct {
	Sender      Sender       `json:"sender,omitempty"`
	Recipient   Recipient    `json:"recipient,omitempty"`
	Timestamp   big.Int      `json:"timestamp,omitempty"`
	MessageData MessageData  `json:"message,omitempty"`
	Read        ReadData     `json:"read,omitempty"`
	Delivery    DeliveryData `json:"delivery,omitempty"`
	Postback    *Postback    `json:"postback,omitempty"`
}

type Postback struct {
	Payload string `json:"payload,omitempty"`
}

type DeliveryData struct {
	MIDs      []string `json:"mids,omitempty"`
	Watermark big.Int  `json:"watermark,omitempty"`
	Sequence  int      `json:"seq,omitempty"`
}

type ReadData struct {
	Watermark big.Int `json:"watermark,omitempty"`
	Sequence  int     `json:"seq,omitempty"`
}

type Sender struct {
	ID string `json:"id,omitempty"`
}

type Recipient struct {
	ID string `json:"id,omitempty"`
}

type MessageData struct {
	IsEcho      bool          `json:"is_echo,omitempty"`
	AppID       big.Int       `json:"app_id,omitempty"`
	MID         string        `json:"mid,omitempty"`
	Sequence    int           `json:"seq,omitempty"`
	StickerID   *big.Int      `json:"sticker_id,omitempty"`
	Text        string        `json:"text,omitempty"`
	Attachments *[]Attachment `json:"attachments,omitempty"`
}

type Attachment struct {
	Type    string  `json:"type,omitempty"`
	Payload Payload `json:"payload,omitempty"`
}

type Payload struct {
	TemplateType string     `json:"template_type,omitempty"`
	Elements     *[]Element `json:"elements,omitempty"`
	Url          string     `json:"url,omitempty"`
	StickerID    *big.Int   `json:"sticker_id,omitempty"`
}
