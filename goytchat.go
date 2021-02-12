package goytchat

import (
	"bytes"
	"encoding/base64"
)

// GetArchivedParam gets the parameter string for live chat.
func GetArchivedParam(videoID string, seekTime uint, topChatOnly bool, channelID string) string {
	var (
		chatType  int
		timestamp = seekTime * 1000000
	)
	if topChatOnly {
		chatType = 4
	} else {
		chatType = 1
	}
	// encode header
	s1_3 := rs(1, []byte(videoID))
	s1_5 := append(rs(1, []byte(channelID)), rs(2, []byte(videoID))...)
	s1 := append(rs(3, s1_3), rs(5, s1_5)...)
	s3 := rs(48687757, rs(1, []byte(videoID)))
	h := append(rs(1, s1), rs(3, s3)...)
	h = append(h, nm(4, 1)...)
	header := base64.URLEncoding.EncodeToString(h)
	// make the protobuf-lookin' thing
	var buf = &bytes.Buffer{}
	buf.Write(rs(3, []byte(header)))
	buf.Write(nm(5, int(timestamp)))
	buf.Write(nm(6, 0))
	buf.Write(nm(7, 0))
	buf.Write(nm(8, 0))
	buf.Write(nm(9, 4))
	buf.Write(rs(10, nm(4, 0)))
	buf.Write(rs(14, nm(chatType, 4)))
	buf.Write(nm(15, 0))
	continuation := rs(156074452, buf.Bytes())
	return base64.URLEncoding.EncodeToString(continuation)
}
