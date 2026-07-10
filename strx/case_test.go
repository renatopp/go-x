package strx_test

import (
	"testing"

	"github.com/renatopp/go-x/strx"
	"github.com/renatopp/go-x/testx"
)

func TestToLower(t *testing.T) {
	testx.Equal(t, "hello, world!", strx.ToLower("Hello, World!"))
	testx.Equal(t, "ohmyhello, world?!$&%", strx.ToLower("OHMyHello, World?!$&%"))
	testx.Equal(t, "helloworld99-", strx.ToLower("HELLOWorld99-"))
	testx.Equal(t, "123,32", strx.ToLower("123,32"))
	testx.Equal(t, "a1b2c3???d4", strx.ToLower("a1b2c3???d4"))
	testx.Equal(t, "d.m.c.", strx.ToLower("D.m.C."))
	testx.Equal(t, "samplesimplecase", strx.ToLower("sampleSimpleCase"))
	testx.Equal(t, "samplesimplecase", strx.ToLower("SampleSimpleCase"))
	testx.Equal(t, "sample-simple-case", strx.ToLower("sample-simple-case"))
	testx.Equal(t, "sample_simple_case", strx.ToLower("sample_simple_case"))
	testx.Equal(t, "sample:simple:case", strx.ToLower("sample:simple:case"))
	testx.Equal(t, "	", strx.ToLower("	"))
	testx.Equal(t, "", strx.ToLower(""))
}

func TestToUpper(t *testing.T) {
	testx.Equal(t, "HELLO, WORLD!", strx.ToUpper("Hello, World!"))
	testx.Equal(t, "OHMYHELLO, WORLD?!$&%", strx.ToUpper("OHMyHello, World?!$&%"))
	testx.Equal(t, "HELLOWORLD99-", strx.ToUpper("HELLOWorld99-"))
	testx.Equal(t, "123,32", strx.ToUpper("123,32"))
	testx.Equal(t, "A1B2C3???D4", strx.ToUpper("a1b2c3???d4"))
	testx.Equal(t, "D.M.C.", strx.ToUpper("D.m.C."))
	testx.Equal(t, "SAMPLESIMPLECASE", strx.ToUpper("sampleSimpleCase"))
	testx.Equal(t, "SAMPLESIMPLECASE", strx.ToUpper("SampleSimpleCase"))
	testx.Equal(t, "SAMPLE-SIMPLE-CASE", strx.ToUpper("sample-simple-case"))
	testx.Equal(t, "SAMPLE_SIMPLE_CASE", strx.ToUpper("sample_simple_case"))
	testx.Equal(t, "SAMPLE:SIMPLE:CASE", strx.ToUpper("sample:simple:case"))
	testx.Equal(t, "	", strx.ToUpper("	"))
	testx.Equal(t, "", strx.ToUpper(""))
}

func TestToTitle(t *testing.T) {
	testx.Equal(t, "Hello, World!", strx.ToTitle("Hello, World!"))
	testx.Equal(t, "Ohmyhello, World?!$&%", strx.ToTitle("OHMyHello, World?!$&%"))
	testx.Equal(t, "Helloworld99-", strx.ToTitle("HELLOWorld99-"))
	testx.Equal(t, "123,32", strx.ToTitle("123,32"))
	testx.Equal(t, "A1b2c3???D4", strx.ToTitle("a1b2c3???d4"))
	testx.Equal(t, "D.M.C.", strx.ToTitle("D.m.C."))
	testx.Equal(t, "Samplesimplecase", strx.ToTitle("sampleSimpleCase"))
	testx.Equal(t, "Samplesimplecase", strx.ToTitle("SampleSimpleCase"))
	testx.Equal(t, "Sample-Simple-Case", strx.ToTitle("sample-simple-case"))
	testx.Equal(t, "Sample_Simple_Case", strx.ToTitle("sample_simple_case"))
	testx.Equal(t, "Sample:Simple:Case", strx.ToTitle("sample:simple:case"))
	testx.Equal(t, "	", strx.ToTitle("	"))
	testx.Equal(t, "", strx.ToTitle(""))
	testx.Equal(t, "Emoji: 🧹, Unicode: 🂣", strx.ToTitle("emoji: 🧹, unicode: 🂣"))
}

func TestToDelimiter(t *testing.T) {
	testx.Equal(t, "hello#world", strx.ToDelimiter("Hello, World!", "#"))
	testx.Equal(t, "ohmy#hello#world", strx.ToDelimiter("OHMyHello, World?!$&%", "#"))
	testx.Equal(t, "helloworld#99", strx.ToDelimiter("HELLOWorld99-", "#"))
	testx.Equal(t, "123#32", strx.ToDelimiter("123,32", "#"))
	testx.Equal(t, "a#1#b#2#c#3#d#4", strx.ToDelimiter("a1b2c3???d4", "#"))
	testx.Equal(t, "d#m#c", strx.ToDelimiter("D.m.C.", "#"))
	testx.Equal(t, "sample#simple#case", strx.ToDelimiter("sampleSimpleCase", "#"))
	testx.Equal(t, "sample#simple#case", strx.ToDelimiter("SampleSimpleCase", "#"))
	testx.Equal(t, "sample#simple#case", strx.ToDelimiter("sample-simple-case", "#"))
	testx.Equal(t, "sample#simple#case", strx.ToDelimiter("sample_simple_case", "#"))
	testx.Equal(t, "sample#simple#case", strx.ToDelimiter("sample:simple:case", "#"))
	testx.Equal(t, "", strx.ToDelimiter("	", "#"))
	testx.Equal(t, "", strx.ToDelimiter("", "#"))
	testx.Equal(t, "emoji🌱unicode", strx.ToDelimiter("Emoji: 🧹, Unicode: 🂣", "🌱"))
}

func TestToSnake(t *testing.T) {
	testx.Equal(t, "hello_world", strx.ToSnake("Hello, World!"))
	testx.Equal(t, "ohmy_hello_world", strx.ToSnake("OHMyHello, World?!$&%"))
	testx.Equal(t, "helloworld_99", strx.ToSnake("HELLOWorld99-"))
	testx.Equal(t, "123_32", strx.ToSnake("123,32"))
	testx.Equal(t, "a_1_b_2_c_3_d_4", strx.ToSnake("a1b2c3???d4"))
	testx.Equal(t, "d_m_c", strx.ToSnake("D.m.C."))
	testx.Equal(t, "sample_simple_case", strx.ToSnake("sampleSimpleCase"))
	testx.Equal(t, "sample_simple_case", strx.ToSnake("SampleSimpleCase"))
	testx.Equal(t, "sample_simple_case", strx.ToSnake("sample-simple-case"))
	testx.Equal(t, "sample_simple_case", strx.ToSnake("sample_simple_case"))
	testx.Equal(t, "sample_simple_case", strx.ToSnake("sample:simple:case"))
	testx.Equal(t, "", strx.ToSnake("	"))
	testx.Equal(t, "", strx.ToSnake(""))
	testx.Equal(t, "emoji_unicode", strx.ToSnake("Emoji: 🧹, Unicode: 🂣"))
}

func TestToUpperSnake(t *testing.T) {
	testx.Equal(t, "HELLO_WORLD", strx.ToUpperSnake("Hello, World!"))
	testx.Equal(t, "OHMY_HELLO_WORLD", strx.ToUpperSnake("OHMyHello, World?!$&%"))
	testx.Equal(t, "HELLOWORLD_99", strx.ToUpperSnake("HELLOWorld99-"))
	testx.Equal(t, "123_32", strx.ToUpperSnake("123,32"))
	testx.Equal(t, "A_1_B_2_C_3_D_4", strx.ToUpperSnake("a1b2c3???d4"))
	testx.Equal(t, "D_M_C", strx.ToUpperSnake("D.m.C."))
	testx.Equal(t, "SAMPLE_SIMPLE_CASE", strx.ToUpperSnake("sampleSimpleCase"))
	testx.Equal(t, "SAMPLE_SIMPLE_CASE", strx.ToUpperSnake("SampleSimpleCase"))
	testx.Equal(t, "SAMPLE_SIMPLE_CASE", strx.ToUpperSnake("sample-simple-case"))
	testx.Equal(t, "SAMPLE_SIMPLE_CASE", strx.ToUpperSnake("sample_simple_case"))
	testx.Equal(t, "SAMPLE_SIMPLE_CASE", strx.ToUpperSnake("sample:simple:case"))
	testx.Equal(t, "", strx.ToUpperSnake("	"))
	testx.Equal(t, "", strx.ToUpperSnake(""))
	testx.Equal(t, "EMOJI_UNICODE", strx.ToUpperSnake("Emoji: 🧹, Unicode: 🂣"))
}

func TestToKebab(t *testing.T) {
	testx.Equal(t, "hello-world", strx.ToKebab("Hello, World!"))
	testx.Equal(t, "ohmy-hello-world", strx.ToKebab("OHMyHello, World?!$&%"))
	testx.Equal(t, "helloworld-99", strx.ToKebab("HELLOWorld99-"))
	testx.Equal(t, "123-32", strx.ToKebab("123,32"))
	testx.Equal(t, "a-1-b-2-c-3-d-4", strx.ToKebab("a1b2c3???d4"))
	testx.Equal(t, "d-m-c", strx.ToKebab("D.m.C."))
	testx.Equal(t, "sample-simple-case", strx.ToKebab("sampleSimpleCase"))
	testx.Equal(t, "sample-simple-case", strx.ToKebab("SampleSimpleCase"))
	testx.Equal(t, "sample-simple-case", strx.ToKebab("sample-simple-case"))
	testx.Equal(t, "sample-simple-case", strx.ToKebab("sample_simple_case"))
	testx.Equal(t, "sample-simple-case", strx.ToKebab("sample:simple:case"))
	testx.Equal(t, "", strx.ToKebab("	"))
	testx.Equal(t, "", strx.ToKebab(""))
	testx.Equal(t, "emoji-unicode", strx.ToKebab("Emoji: 🧹, Unicode: 🂣"))
}

func TestToUpperKebab(t *testing.T) {
	testx.Equal(t, "HELLO-WORLD", strx.ToUpperKebab("Hello, World!"))
	testx.Equal(t, "OHMY-HELLO-WORLD", strx.ToUpperKebab("OHMyHello, World?!$&%"))
	testx.Equal(t, "HELLOWORLD-99", strx.ToUpperKebab("HELLOWorld99-"))
	testx.Equal(t, "123-32", strx.ToUpperKebab("123,32"))
	testx.Equal(t, "A-1-B-2-C-3-D-4", strx.ToUpperKebab("a1b2c3???d4"))
	testx.Equal(t, "D-M-C", strx.ToUpperKebab("D.m.C."))
	testx.Equal(t, "SAMPLE-SIMPLE-CASE", strx.ToUpperKebab("sampleSimpleCase"))
	testx.Equal(t, "SAMPLE-SIMPLE-CASE", strx.ToUpperKebab("SampleSimpleCase"))
	testx.Equal(t, "SAMPLE-SIMPLE-CASE", strx.ToUpperKebab("sample-simple-case"))
	testx.Equal(t, "SAMPLE-SIMPLE-CASE", strx.ToUpperKebab("sample_simple_case"))
	testx.Equal(t, "SAMPLE-SIMPLE-CASE", strx.ToUpperKebab("sample:simple:case"))
	testx.Equal(t, "", strx.ToUpperKebab("	"))
	testx.Equal(t, "", strx.ToUpperKebab(""))
	testx.Equal(t, "EMOJI-UNICODE", strx.ToUpperKebab("Emoji: 🧹, Unicode: 🂣"))
}

func TestToCamel(t *testing.T) {
	testx.Equal(t, "helloWorld", strx.ToCamel("Hello, World!"))
	testx.Equal(t, "ohmyHelloWorld", strx.ToCamel("OHMyHello, World?!$&%"))
	testx.Equal(t, "helloworld99", strx.ToCamel("HELLOWorld99-"))
	testx.Equal(t, "12332", strx.ToCamel("123,32"))
	testx.Equal(t, "a1B2C3D4", strx.ToCamel("a1b2c3???d4"))
	testx.Equal(t, "dMC", strx.ToCamel("D.m.C."))
	testx.Equal(t, "sampleSimpleCase", strx.ToCamel("sampleSimpleCase"))
	testx.Equal(t, "sampleSimpleCase", strx.ToCamel("SampleSimpleCase"))
	testx.Equal(t, "sampleSimpleCase", strx.ToCamel("sample-simple-case"))
	testx.Equal(t, "sampleSimpleCase", strx.ToCamel("sample_simple_case"))
	testx.Equal(t, "sampleSimpleCase", strx.ToCamel("sample:simple:case"))
	testx.Equal(t, "", strx.ToCamel("	"))
	testx.Equal(t, "", strx.ToCamel(""))
	testx.Equal(t, "emojiUnicode", strx.ToCamel("Emoji: 🧹, Unicode: 🂣"))
}

func TestToPascal(t *testing.T) {
	testx.Equal(t, "HelloWorld", strx.ToPascal("Hello, World!"))
	testx.Equal(t, "OhmyHelloWorld", strx.ToPascal("OHMyHello, World?!$&%"))
	testx.Equal(t, "Helloworld99", strx.ToPascal("HELLOWorld99-"))
	testx.Equal(t, "12332", strx.ToPascal("123,32"))
	testx.Equal(t, "A1B2C3D4", strx.ToPascal("a1b2c3???d4"))
	testx.Equal(t, "DMC", strx.ToPascal("D.m.C."))
	testx.Equal(t, "SampleSimpleCase", strx.ToPascal("sampleSimpleCase"))
	testx.Equal(t, "SampleSimpleCase", strx.ToPascal("SampleSimpleCase"))
	testx.Equal(t, "SampleSimpleCase", strx.ToPascal("sample-simple-case"))
	testx.Equal(t, "SampleSimpleCase", strx.ToPascal("sample_simple_case"))
	testx.Equal(t, "SampleSimpleCase", strx.ToPascal("sample:simple:case"))
	testx.Equal(t, "", strx.ToPascal("	"))
	testx.Equal(t, "", strx.ToPascal(""))
	testx.Equal(t, "EmojiUnicode", strx.ToPascal("Emoji: 🧹, Unicode: 🂣"))
}
