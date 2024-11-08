package service

import (
	"bitbucket.org/coontec/javaClass/class/model/message"
	"bitbucket.org/coontec/javaClass/class/service/deserializer"
	"strings"
	"testing"
)

func TestLog4j(t *testing.T) {
	// 설정 클래스 선언.
	classLoader := NewClassLoader("./test-features")

	message := message.NewMessage()
	message.Headers["loader"] = classLoader
	message.Headers["printer"] = nil
	message.Headers["configuration"] = make(map[string]interface{})

	// 실행 테스트.
	deserializer := deserializer.NewDeserializeClassFileProcessor()

	for k, _ := range classLoader.Map {
		if strings.HasSuffix(k, ".class") && (strings.Index(k, "$") == -1) {
			internalTypeName := k[:len(k)-6]
			message.Headers["mainInternalTypeName"] = internalTypeName

			err := deserializer.Process(message)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}
