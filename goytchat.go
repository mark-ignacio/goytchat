package goytchat

import (
	"bytes"
	"encoding/base64"
	"math/rand"
	"time"
)

func genHeader(videoID, channelID string) string {
	s1_3 := rs(1, []byte(videoID))
	s1_5 := append(rs(1, []byte(channelID)), rs(2, []byte(videoID))...)
	s1 := append(rs(3, s1_3), rs(5, s1_5)...)
	s3 := rs(48687757, rs(1, []byte(videoID)))
	h := append(rs(1, s1), rs(3, s3)...)
	h = append(h, nm(4, 1)...)
	return base64.URLEncoding.EncodeToString(h)
}

// GetArchivedParam gets the parameter string for archived videos' live chat.
func GetArchivedParam(videoID, channelID string, seekTime uint, topChatOnly bool) string {
	var (
		chatType  int
		timestamp = seekTime * 1000000
	)
	if topChatOnly {
		chatType = 4
	} else {
		chatType = 1
	}
	header := genHeader(videoID, channelID)
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

const liveTimeMult = 1000000

// GetLiveParam gets the parameter string for live videos' live chat.
func GetLiveParam(videoID, channelID string, pastSeconds uint, topChatOnly bool) string {
	var chatType int
	if topChatOnly {
		chatType = 4
	} else {
		chatType = 1
	}
	head := genHeader(videoID, channelID)
	// generate times
	// TODO: is math/rand uniformly distributed by default?
	now := float64(time.Now().Unix())
	ts1 := int((now - randFloat64Between(0, 3)) * liveTimeMult)
	ts2 := int((now - randFloat64Between(0.01, 0.99)) * liveTimeMult)
	ts3 := int((now - float64(pastSeconds) + rand.Float64()) * liveTimeMult)
	ts4 := int((now - randFloat64Between(600, 3600)) * liveTimeMult)
	ts5 := int((now - randFloat64Between(0.01, 0.99)) * liveTimeMult)

	// make the protobuf-lookin thing
	var bodyBuf = &bytes.Buffer{}
	bodyBuf.Write(nm(1, 0))
	bodyBuf.Write(nm(2, 0))
	bodyBuf.Write(nm(3, 0))
	bodyBuf.Write(nm(4, 0))
	bodyBuf.Write(rs(7, []byte{}))
	bodyBuf.Write(nm(8, 0))
	bodyBuf.Write(rs(9, []byte{}))
	bodyBuf.Write(nm(10, ts2))
	bodyBuf.Write(nm(11, 3))
	bodyBuf.Write(nm(15, 0))
	var entityBuf = &bytes.Buffer{}
	entityBuf.Write(rs(3, []byte(head)))
	entityBuf.Write(nm(5, ts1))
	entityBuf.Write(nm(6, 0))
	entityBuf.Write(nm(7, 0))
	entityBuf.Write(nm(8, 1))
	entityBuf.Write(rs(9, bodyBuf.Bytes()))
	entityBuf.Write(nm(10, ts3))
	entityBuf.Write(nm(11, ts4))
	entityBuf.Write(nm(13, chatType))
	entityBuf.Write(rs(16, nm(1, chatType)))
	entityBuf.Write(nm(17, 0))
	entityBuf.Write(rs(19, nm(1, 0)))
	entityBuf.Write(nm(20, ts5))
	continuation := rs(119693434, entityBuf.Bytes())
	return base64.URLEncoding.EncodeToString(continuation)
}

func randFloat64Between(floor, ceiling float64) float64 {
	diff := ceiling - floor
	return (rand.ExpFloat64() * diff) + floor
}
