package main
import(
	"fmt"
	"strings"
)
type celsius float64
type fahrenheit float64
func F2C(f fahrenheit) celsius{
	c:=(f-32)/1.8
	return celsius(c)
}
func C2F(c celsius) fahrenheit{
	f:=(c*9.0/5.0)+32
	return fahrenheit(f)
}
func main(){
	var c1 celsius=100.0
	fmt.Printf("Celsius : %f, Fahrenheit : %f\n",c1,C2F(c1))
	var f1 fahrenheit=45.0
	fmt.Printf("Fahrenheit : %f, Celsius  : %f\n",f1,F2C(f1))
	s:="golang"
	for i:=0; i<len(s); i++{
		fmt.Printf("%c %d %T\n", s[i], s[i], s[i])
	}
	if strings.Contains(s,"la"){
		fmt.Printf("%s contains la\n", s)
	}else{
		fmt.Printf("%s does not contain la\n",s)
	}
}