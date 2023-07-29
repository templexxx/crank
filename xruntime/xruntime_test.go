package xruntime

import (
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"github.com/klauspost/cpuid/v2"
	"github.com/templexxx/crank/xtest"
	"github.com/templexxx/tsc"
)

func TestProcYield(t *testing.T) {

	if !xtest.IsPropEnabled() {
		t.Skip("properties testing isn't enabled")
	}

	cs := []int{10, 20, 30, 60, 120, 240}

	for _, c := range cs {
		// TODO using TSC register to report cost (as property testing)
		tsc.ForbidOutOfOrder()
		start := tsc.UnixNano()
		for i := 0; i < 1000; i++ {
			ProcYield(uint32(c))
		}
		cost := (tsc.UnixNano() - start) / 1000
		t.Logf("%d cycles cost: %dns on %s", c, cost, getCPUBrand())
	}
}

func getCPUBrand() string {

	brand := cpuid.CPU.BrandName
	if brand != "" {
		return cpuid.CPU.BrandName
	}

	if runtime.GOOS == `darwin` {
		brand = getCPUBrandOnDarwin()
	}
	if brand == "" {
		return "unknown"
	}
	return brand
}

func getCPUBrandOnDarwin() string {
	grep := exec.Command("grep", "machdep.cpu.brand_string")
	sysctl := exec.Command("sysctl", "-a")

	pipe, err := sysctl.StdoutPipe()
	if err != nil {
		return ""
	}
	defer pipe.Close()

	grep.Stdin = pipe

	err = sysctl.Start()
	if err != nil {
		return ""
	}
	res, err := grep.Output()
	if err != nil {
		return ""
	}
	return strings.Split(string(res), ": ")[1]
}
