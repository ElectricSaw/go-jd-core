package tests

import (
	"bitbucket.org/coontec/go-jd-core/class/model/message"
	"bitbucket.org/coontec/go-jd-core/class/service/converter"
	"bitbucket.org/coontec/go-jd-core/class/service/deserializer"
	"bitbucket.org/coontec/go-jd-core/class/tests/testutils"
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
	dsr := deserializer.NewDeserializeClassFileProcessor()
	cnv := converter.NewClassFileToJavaSyntaxProcessor()

	for k, _ := range classLoader.Map {
		if strings.HasSuffix(k, ".class") && (strings.Index(k, "$") == -1) {
			internalTypeName := k[:len(k)-6]
			message.Headers["mainInternalTypeName"] = internalTypeName

			err := dsr.Process(message)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}
