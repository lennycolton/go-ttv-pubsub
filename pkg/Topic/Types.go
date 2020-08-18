package Topic

import "strings"

type Type string

const (
	TypeBits                  = Type("channel-bits-events-v2")
	TypeBitsBadgeNotification = Type("channel-bits-badge-unlocks")
	TypeSubscriptions         = Type("channel-subscribe-events-v1")
	TypeChannelPoints         = Type("channel-points-channel-v1")
	TypeWhispers              = Type("whispers")
	TypeModerationAction      = Type("chat_moderator_actions")
	TypeInvalid               = Type("invalid")
	TypePong                  = Type("PONG")
	TypeReconnect             = Type("RECONNECT")
)

func GetType(topic string) Type {
	pieces := strings.Split(topic, ".")

	if len(pieces) < 2 {
		return TypeInvalid
	}

	switch Type(pieces[0]) {
	case TypeBits:
		return TypeBits
	case TypeBitsBadgeNotification:
		return TypeBitsBadgeNotification
	case TypeSubscriptions:
		return TypeSubscriptions
	case TypeChannelPoints:
		return TypeChannelPoints
	case TypeWhispers:
		return TypeWhispers
	case TypeModerationAction:
		return TypeModerationAction
	}

	return TypeInvalid
}
