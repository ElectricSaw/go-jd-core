package tests

import (
	"bitbucket.org/coontec/javaClass/class/model/message"
	"bitbucket.org/coontec/javaClass/class/service/deserializer"
	"bitbucket.org/coontec/javaClass/class/tests/testutils"
	"strings"
	"testing"
)

func TestLog4j(t *testing.T) {
	// 설정 클래스 선언.
	classLoader := testutils.NewDirectoryLoader("./test-features")
	printer := testutils.NewCounterPrinterEmpty()
	//statistics := make(map[string]int)
	configuration := make(map[string]interface{})

	message := message.NewMessage()
	message.Headers["loader"] = classLoader
	message.Headers["printer"] = printer
	message.Headers["configuration"] = configuration

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
