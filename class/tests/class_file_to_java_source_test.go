package tests

import (
	"github.com/ElectricSaw/go-jd-core/class/model/message"
	"github.com/ElectricSaw/go-jd-core/class/service/converter"
	"github.com/ElectricSaw/go-jd-core/class/service/deserializer"
	"github.com/ElectricSaw/go-jd-core/class/service/fragmenter"
	"github.com/ElectricSaw/go-jd-core/class/service/layouter"
	"github.com/ElectricSaw/go-jd-core/class/service/tokenizer"
	"github.com/ElectricSaw/go-jd-core/class/service/writer"
	"github.com/ElectricSaw/go-jd-core/class/tests/testutils"
	"strings"
	"testing"
)

func TestLog4j(t *testing.T) {
	// 설정 클래스 선언.
	classLoader := testutils.NewDirectoryLoader("./test-features")
	printer := testutils.NewCounterPrinterEmpty()
	//statistics := make(map[string]int)
	configuration := make(map[string]interface{})

	configuration["realignLineNumbers"] = true

	message := message.NewMessage()
	message.Headers["loader"] = classLoader
	message.Headers["printer"] = printer
	message.Headers["configuration"] = configuration

	// 실행 테스트.
	dsr := deserializer.NewDeserializeClassFileProcessor()
	cnv := converter.NewClassFileToJavaSyntaxProcessor()
	frg := fragmenter.NewJavaSyntaxToJavaFragmentProcessor()
	lyt := layouter.NewLayoutFragmentProcessor()
	tkn := tokenizer.NewJavaFragmentToTokenProcessor()
	wri := writer.NewWriteTokenProcessor()

	for k, _ := range classLoader.Map {
		if strings.HasSuffix(k, ".class") && (strings.Index(k, "$") == -1) {
			internalTypeName := k[:len(k)-6]
			message.Headers["mainInternalTypeName"] = internalTypeName

			err := dsr.Process(message)
			err = cnv.Process(message)
			err = frg.Process(message)
			err = lyt.Process(message)
			err = tkn.Process(message)
			err = wri.Process(message)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}
