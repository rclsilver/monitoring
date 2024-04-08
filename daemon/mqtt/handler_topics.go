package mqtt

import (
	"fmt"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"golang.org/x/exp/maps"
)

type listTopicsOut struct {
	Topics []string `json:"topics"`
}

func (mqtt *MQTTComponent) handlerListTopics(c *gin.Context) (*listTopicsOut, error) {
	mqtt.topicsMut.Lock()
	defer mqtt.topicsMut.Unlock()

	topics := maps.Keys(mqtt.topics)
	slices.Sort(topics)

	return &listTopicsOut{
		Topics: topics,
	}, nil
}

type getTopicOut struct {
	Timestamp time.Time `json:"timestamp"`
	Topic     Topic     `json:"topic"`
}

func (mqtt *MQTTComponent) handlerGetTopic(c *gin.Context) (*getTopicOut, error) {
	topic, _ := c.Params.Get("topic")

	if len(topic) == 0 {
		return nil, errors.NewBadRequest(nil, "invalid topic name")
	}

	mqtt.topicsMut.Lock()
	defer mqtt.topicsMut.Unlock()

	v, ok := mqtt.topics[topic]
	if !ok {
		return nil, errors.NewNotFound(nil, fmt.Sprintf("topic %q not found", topic))
	}

	return &getTopicOut{
		Timestamp: time.Now(),
		Topic:     *v,
	}, nil
}
