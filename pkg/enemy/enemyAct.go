package enemy

import "github.com/sh-miyoshi/ryuzinroku/pkg/common"

//移動パターン0
//下がってきて停滞して上がっていく
func act0(obj *enemy) {
	if obj.count == 0 {
		obj.vy = 3
	}

	if obj.count == 40 {
		obj.vy = 0
	}

	if obj.count == 40+obj.Wait {
		obj.vy = -3
	}
}

//移動パターン1
//下がってきて停滞して左下に行く
func act1(obj *enemy) {
	if obj.count == 0 {
		obj.vy = 3
	}

	if obj.count == 40 {
		obj.vy = 0
	}

	if obj.count == 40+obj.Wait {
		obj.vx = -1
		obj.vy = 2
		obj.direct = common.DirectLeft
	}
}

// //移動パターン2
// //下がってきて停滞して右下に行く
func act2(obj *enemy) {
	if obj.count == 0 {
		obj.vy = 3
	}

	if obj.count == 40 {
		obj.vy = 0
	}

	if obj.count == 40+obj.Wait {
		obj.vx = 1
		obj.vy = 2
		obj.direct = common.DirectRight
	}
}

// //行動パターン3
// //すばやく降りてきて左へ
func act3(obj *enemy) {
	if obj.count == 0 {
		obj.vy = 5
	}

	if obj.count == 30 {
		obj.direct = common.DirectLeft
	}

	if obj.count < 100 {
		obj.vx -= 5.0 / 100.0
		obj.vy -= 5.0 / 100.0
	}
}

// //行動パターン4
// //すばやく降りてきて右へ
func act4(obj *enemy) {
	if obj.count == 0 {
		obj.vy = 5
	}

	if obj.count == 30 {
		obj.direct = common.DirectRight
	}

	if obj.count < 100 {
		obj.vx += 5.0 / 100.0
		obj.vy -= 5.0 / 100.0
	}
}

// //行動パターン5
// //斜め左下へ
func act5(obj *enemy) {
	if obj.count == 0 {
		obj.vx = -1
		obj.vy = 2
		obj.direct = common.DirectLeft
	}
}

// //行動パターン6
// //斜め右下へ
func act6(obj *enemy) {
	if obj.count == 0 {
		obj.vx = 1
		obj.vy = 2
		obj.direct = common.DirectRight
	}
}

// //移動パターン7
// //停滞してそのまま左上に
func act7(obj *enemy) {
	if obj.count == obj.Wait {
		obj.vx = -0.7
		obj.vy = -0.3
		obj.direct = common.DirectLeft
	}
}

// //移動パターン8
// //停滞してそのまま右上に
func act8(obj *enemy) {
	if obj.count == obj.Wait {
		obj.vx = 0.7
		obj.vy = -0.3
		obj.direct = common.DirectRight
	}
}

// //移動パターン9
// //停滞してそのまま上に
func act9(obj *enemy) {
	if obj.count == obj.Wait {
		obj.vy = -1
	}
}

// TODO: enemy move by angle and speed
// //移動パターン10
// //下がってきてウロウロして上がっていく
// void enemy_pattern10(int i){
//     int t=enemy[i].cnt;
//     if(t==0) enemy[i].vy=4;//下がってくる
//     if(t==40)enemy[i].vy=0;//止まる
//     if(t>=40){
//         if(t%60==0){
//             int r=cos(enemy[i].ang)< 0 ? 0 : 1;
//             enemy[i].sp=6+rang(2);
//             enemy[i].ang=rang(PI/4)+PI*r;
//             enemy[i].muki=2-2*r;
//         }
//         enemy[i].sp*=0.95;
//     }
//     if(t>=40+enemy[i].wait){
//         enemy[i].vy-=0.05;
//     }
// }
