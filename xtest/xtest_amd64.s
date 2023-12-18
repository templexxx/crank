#include "textflag.h"

TEXT ·spin(SB),NOSPLIT,$0-0
	MOVL	cycles+0(FP), AX

loop:
	PAUSE
	SUBL	$1, AX
	JNZ	    loop
	RET
