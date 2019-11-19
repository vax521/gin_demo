package main

import "fmt"

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0`
	return fmt.Sprintf(strFormat, de.dividee)
}

func divide(varDividee int, varDivider int) (result int, errMsg string) {
	if varDivider == 0 {
		varData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errMsg = varData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}
func main() {
	// 正常情况
	if result, errorMsg := divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当被除数为零的时候会返回错误信息
	if _, errorMsg := divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

}
