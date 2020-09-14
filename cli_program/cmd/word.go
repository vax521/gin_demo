package cmd

import (
	"github.com/spf13/cobra"
	"strings"
	"unicode"
)

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  "支持多种单词格式转换",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {

}

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func toLower(s string) string {
	return strings.ToLower(s)
}

//下划线单词转驼峰
func underscoreTOUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

func underscoreToLowerCamelCase(s string) string {
	s = underscoreTOUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func camelToUnderscore(s string) string {
	var output []rune
	for i, char := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(char))
			continue
		}
		if unicode.IsUpper(char) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(char))
	}
	return string(output)
}
