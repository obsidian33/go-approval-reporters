package reporters

import (
	"os"
	"os/exec"
	"strconv"

	. "github.com/approvals/go-approval-tests/reporters"
	"github.com/approvals/go-approval-tests/utils"
)

type deltaDiff struct{ width int }

func NewDeltaDiffReporter() Reporter {
	return &deltaDiff{width: 200}
}

func NewDeltaDiffReporterWithWidth(width int) Reporter {
	return &deltaDiff{width: width}
}

func (d *deltaDiff) Report(approved, received string) bool {
	utils.EnsureExists(approved)

	cmd := exec.Command("delta", "--side-by-side", "--width="+strconv.Itoa(d.width), approved, received)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	NewClipboardReporter().Report(approved, received)
	return true
}
