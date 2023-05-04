package numberofvalidwordsinasentense

import "testing"

func TestCountValidWords(t *testing.T) {
	type test struct {
		name     string
		src      string
		expected int
	}
	tests := []test{
		{
			name:     "test 1",
			src:      "cat and  dog",
			expected: 3,
		},
		{
			name:     "test 2",
			src:      "!this  1-s b8d!",
			expected: 0,
		},
		{
			name:     "test 3",
			src:      "alice and  bob are playing stone-game10",
			expected: 5,
		},
		{
			name:     "test 4",
			src:      "0",
			expected: 0,
		},
		{
			name:     "test 5",
			src:      "a-b-c",
			expected: 0,
		},
		{
			name:     "test 6",
			src:      ". ! 7hk  al6 l! aon49esj35la k3 7u2tkh  7i9y5  !jyylhppd et v- h!ogsouv 5",
			expected: 4,
		},
		{
			name:     "test 7",
			src:      " 62   nvtk0wr4f  8 qt3r! w1ph 1l ,e0d 0n 2v 7c.  n06huu2n9 s9   ui4 nsr!d7olr  q-, vqdo!btpmtmui.bb83lf g .!v9-lg 2fyoykex uy5a 8v whvu8 .y sc5 -0n4 zo pfgju 5u 4 3x,3!wl  fv4   s  aig cf j1 a i  8m5o1  !u n!.1tz87d3 .9    n a3  .xb1p9f  b1i a j8s2 cugf l494cx1! hisceovf3 8d93 sg 4r.f1z9w   4- cb r97jo hln3s h2 o .  8dx08as7l!mcmc isa49afk i1 fk,s e !1 ln rt2vhu 4ks4zq c w  o- 6  5!.n8ten0 6mk 2k2y3e335,yj  h p3 5 -0  5g1c  tr49, ,qp9 -v p  7p4v110926wwr h x wklq u zo 16. !8  u63n0c l3 yckifu 1cgz t.i   lh w xa l,jt   hpi ng-gvtk8 9 j u9qfcd!2  kyu42v dmv.cst6i5fo rxhw4wvp2 1 okc8!  z aribcam0  cp-zp,!e x  agj-gb3 !om3934 k vnuo056h g7 t-6j! 8w8fncebuj-lq    inzqhw v39,  f e 9. 50 , ru3r  mbuab  6  wz dw79.av2xp . gbmy gc s6pi pra4fo9fwq k   j-ppy -3vpf   o k4hy3 -!..5s ,2 k5 j p38dtd   !i   b!fgj,nx qgif ",
			expected: 49,
		},
		{
			name:     "test 8",
			src:      "b-a-c f-d",
			expected: 1,
		},
		{
			name:     "test 9",
			src:      "q-o  x-p! g-l- q-n  f-o, m-u. m-i! y-k, i-j, d-p! e-t, h-u  j-j- d-z- v-w, r-a  i-h. d-a! z-o, v-l, ",
			expected: 17,
		},
		{
			name:     "test 10",
			src:      "kf6wo ! 6fjwh,cz48a jr f d !h qf qgh7s b7-0.,88tk 4h8y   bx 1 a0 90z6he6  ,gm,p.93-q 4e2pxyc o 46u ,o jm1r, - r1 b !,w,exml  4q6sok!y8.7fhs ihcy6ro3  u z e7 hg , -3b  jcz5vj 89v2dgrcm znq77u4us fvjrx qqs u u fb7pqz0s n6dfne375wz t  6sdsyz4wak-q73 1!j-uh 8p 6b.i , .u ,y-h tzmwam9bx qd c a  ! n3!-2m9c 0 uf6,6rqq, 7    wgl 6p3i v-br -e !o0gm  f3u3i7!e ,wef h fzpr xk ,0 vor ie2 2  4o jw z73 o  .7gpr 79.je5surs7z!ht uprf 8z9clojw0718yk  i  j a9 u 6iu  lsb8 pv   6c4xh4x q ! eo h r okll s 7 b dm9 4 n5q07 0uo4 5afzk vk mmp h  f ub  xb s51a o 8e vtg2ru 8k  p5f  btcio77 6,g  ac t s,jnj1 v6q!0mzvno8tv! gj   e s0q . o!gclzw!mh0  hpx fxp00 j!rb081r4 41szmd 78b1g2k  x  9 n v uo! -g sfs,  ta-  psw   grc  rkf h0ig yl .6v5t5u7v 8k5 dzej euyqy r28l15-zucss94tk89 bx3jl!ihpl t.9nlh.j6 5m0oqwy u 7xhe0 h- 6hz0  wwyqyw ifbljpkgz i,l0t zeq7jx839k15r n766n ,2 y2 dq  m y-yzq3os4  sethhm v0 i90le- m.ymc03 ua . c6  4 7ex3w  f4 n ruevbm ul 3c   nj.mv kvxcr.41vu5w6 .4p ma5f8lavofoh- i pj hq-ej-wtr5 1n4 1o",
			expected: 80,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countValidWords(tt.src)
			if actual != tt.expected {
				t.Errorf("Actual: %d; but Expected: %d", actual, tt.expected)
			}
		})
	}
}
