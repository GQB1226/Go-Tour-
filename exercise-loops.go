package main
/*
Go指南练习题：使用牛顿法实现开方函数
*/
import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64{
    z:=1.0
    for i:=0;i<10;i++{
        z=z-(z*z-x)/(2*z)
    }
    return z
}
func Sqrt1(x float64) float64{
    z:=1.0
    for{
        y:=z-(z*z-x)/(2*z)
        if math.Abs(y-z)<1e-10{
            return y
        }
        z=y
    }
}
func main(){
    fmt.Printf("使用标准数学库计算出%f的开方是%f\n",2.0,math.Sqrt(2.0))
    fmt.Printf("使用该函数计算出来的值为:%f\n",Sqrt(2))
    fmt.Printf("使用极限接近方法计算出来的值:%f\n",Sqrt1(2))


}
