package main

import "fmt"

func main(){
	/**
	 * 接口和结构体的关系
	 */

	//person1 := NewChinese()
	//person1.Sleep(6)
	//
	//person2 := NewJapanese()
	//person2.

	/**
	 * 总结：什么情况下适合使用接口
	 * 		关注共性时，我们可以采用接口来进行定义。共性的具体实现使用结构体进行实现
	 */
}

/**
 * 对象：模糊的标准
 */
type Person interface {
	Sleep (size int)//xx人每天睡眠多少小时
	Eat (food string)//xxx喜欢吃xxx食物
	Run (length int)//xx人跑步跑了xx公里
	Money () int  //该人有多少钱
	Height () int //身高多少
	Weight () int //体重
}

type Chinese struct {
	Color 	string  //皮肤
	Eye		string  //眼镜
	Name 	string  //中国人的名字
	Finance int		//财富金融
	High	int		//
	Wei		int
	Majiang bool
}

//上帝1
func NewChinese() *Chinese{
	chin := &Chinese{
		Color:   "黄皮肤",
		Eye:     "黑眼睛",
		Name:    "上海名媛",
		Finance: 3000000,
		High:    168,
		Wei:     98,
		Majiang: true,
	}
	return chin
}

//上帝
func NewJapanese() *Japanese{
	japan := &Japanese{
		Name:       "小仓优子",
		MoneSum:    30000000,
		HeightSize: 158,
		WeightNum:  200,
		Phps:       true,
	}
	return japan
}

func (c *Chinese) Sleep(size int){
	fmt.Println(c.Name+"每天睡眠：",size,"小时")

}

func (c *Chinese) Eat(food string){
	fmt.Println(c.Name+"喜欢吃：",food,)
}

func (c *Chinese) Run(length int){
	fmt.Println(c.Name+"每次跑步能跑：",length,"公里")
}

func (c *Chinese) Money() int{
	return c.Finance
}

func (c *Chinese) Height() int{
	return c.Height()
}
func (c *Chinese) Weight() int{
	return c.Wei
}
func (c *Chinese) PlayMajiang() bool{
	return c.Majiang
}

/**
 * 日本人
 */
type Japanese struct {
	Name string//姓名
	MoneSum int //钱
	HeightSize int
	WeightNum int
	Phps 	bool
}

func (j *Japanese) Sleep(size int){
	fmt.Println(j.Name+"每天睡：",size,"小时")
}

func (j *Japanese) Eat(food string){
	fmt.Println(j.Name+"喜欢吃",food)
}
func (j *Japanese) Run(length int){
	fmt.Println(j.Name+"跑步跑：",length,"公里")
}
func (j *Japanese) Money() int{
	return j.MoneSum
}
func ( j*Japanese) Height()int{
	return j.HeightSize
}
func ( j*Japanese) Weight()int{
	return j.WeightNum
}
func ( j*Japanese) Php()bool{
	return j.Phps
}