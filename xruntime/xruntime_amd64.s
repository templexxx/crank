#include "textflag.h"

TEXT ·pyield(SB),NOSPLIT,$0-0
	MOVL	cycles+0(FP), AX

loop:
	PAUSE
	SUBL	$1, AX
	JNZ	    loop
	RET
