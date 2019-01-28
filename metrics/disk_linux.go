package metrics

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/openether/ethcore/logger/glog"
)

func readDiskStats(stats *diskStats) {
	file := fmt.Sprintf("/proc/%d/io", os.Getpid())
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		glog.Errorf("%s: %s", file, err)
		return
	}

	for _, line := range strings.Split(string(bytes), "\n") {
		i := strings.Index(line, ": ")
		if i < 0 {
			continue
		}

		var p *int64
		switch line[:i] {
		case "syscr":
			p = &stats.ReadCount
		case "syscw":
			p = &stats.WriteCount
		case "rchar":
			p = &stats.ReadBytes
		case "wchar":
			p = &stats.WriteBytes
		default:
			continue
		}

		*p, err = strconv.ParseInt(line[i+2:], 10, 64)
		if err != nil {
			glog.Errorf("%s: line %q: %s", file, line, err)
		}
	}
}
