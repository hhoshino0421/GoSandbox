package sandbox

import (
	"fmt"
	"math"
)

///*
// * 大文字は公開メソッド
// * 小文字は非公開メソッド
// *
// */
//func Test01() {
//	fmt.Printf("%s","Test01 called. \n")
//
//	test01()
//
//	fmt.Printf("%s","Test01 end. \n")
//
//}
//
//func test01() {
//	fmt.Printf("%s","test01 called. \n")
//}

const (
	width, height 	= 600, 300				//キャンバスの大きさ
	cells			= 100					//格子のマス目の数
	xyrange			= 30.0					//軸の範囲
	xyscale			= width / 2 /xyrange	//x単位およびy単位あたりの画素数
	zscale			= height * 0.4			//z単位あたりの画素数
	angle			= math.Pi / 6			//x,y軸の角度(=30度)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) 	//sin30,cos30

func Surface() {

	fmt.Printf("%s","Surface start. \n")

	message := "<svg xmlns='http://www.w3.org/2000svg' style='stroke: grey; fill: white; stroke-width: 0.7'"
	fmt.Printf("%s", message)
	fmt.Printf(" width='%d' height='%d'> \n", width, height)


	for i :=0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i + 1,j)
			bx, by := corner(i,j)
			cx, cy := corner(i, j + 1)
			dx, dy := corner(i + 1,j + 1)
			fmt.Printf("<polygon point='%g,%g,%g,%g %g,%g,%g,%g'/>\n",
			ax, ay, bx, by, cx, cy, dx, dy)

		}
 	}


	fmt.Printf("%s", "</svg> ")

	fmt.Printf("%s","Surface end. \n")


}

func corner(i, j int) (float64, float64){
	//マス目の角の点を見つける
	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)

	//面の高さを計算する
	z := f(x,y)

	//(x,y,z)を2-Dのキャンバスへ等角的に投影
	sx := width / 2 + (x - y) * cos30 * xyscale
	sy := height / 2 + (x + y) * sin30 * xyscale - z*zscale

	return sx, sy

}

func f(x,y float64) float64 {
	r := math.Hypot(x, y)		//(0,0)からの距離
	return math.Sin(r)
}