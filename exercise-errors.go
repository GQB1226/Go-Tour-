package main
import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
    //这里文中提出如果直接使用fmt.Sprint(e)的话会陷入死循环，为什么呢
    //因为这里的e本身是经过Error转换的，那么上面的表达式就变成了fmt.Sprint(r.Error())不断调用自身
    return "cannot Sqrt negative number:"+fmt.Sprint(float64(e))
}
func Sqrt(x float64)(float64,error){
    if(x<0){
        return 0,ErrNegativeSqrt(x)
    }
    return math.Sqrt(x),nil
}
func main(){
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
