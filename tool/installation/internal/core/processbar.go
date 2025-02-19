package core

import (
	"fmt"
	"time"
)

type Bar struct {
	percent  int64  // progress percentage
	cur      int64  // current progress
	total    int64  // total value for progress
	rate     string // the actual progress bar to be printed
	graph    string // the fill value for progress bar
	taskname string //bartaskname
}

func (bar *Bar) NewOption(start, total int64, taskname string) {
	bar.cur = start
	bar.total = total
	bar.taskname = tasknameFormat(taskname)
	bar.rate = ""
	if bar.graph == "" {
		bar.graph = "="
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 2 {
		bar.rate += bar.graph // initial progress position
	}
}

func (bar *Bar) getPercent() int64 {
	return int64((float32(bar.cur) / float32(bar.total)) * 100)
}

func (bar *Bar) Add(cur int64) {
	if cur > 1 {
		var counter int64
		realcurl := cur
		for counter = bar.percent + 1; counter <= realcurl; counter++ {
			time.Sleep(100 * time.Millisecond)
			bar.cur = counter
			last := bar.percent
			bar.percent = bar.getPercent()
			if bar.percent != last && bar.percent%2 == 0 {
				bar.rate += bar.graph
			}
			fmt.Printf("\r[%-50s]%3d%% %8d/%d [ %s ]", bar.rate, bar.percent, bar.cur, bar.total, bar.taskname)
		}
	} else {
		bar.cur = cur
		last := bar.percent
		bar.percent = bar.getPercent()
		if bar.percent != last && bar.percent%2 == 0 {
			bar.rate += bar.graph
		}
		fmt.Printf("\r[%-50s]%3d%% %8d/%d [ %s ]", bar.rate, bar.percent, bar.cur, bar.total, bar.taskname)
	}

}

func (bar *Bar) Barflush() {
	fmt.Println("\033[2J")
	fmt.Println("\033[H")
}

func tasknameFormat(taskname string) string {
	var stringSize int = 40
	var pandingnum int
	var tempnum, remainder int
	var leftspace, rightspace, remainderspace, result string
	remainderspace = " "
	leftspace = " "
	rightspace = " "
	tempnum = 0
	remainder = 0
	tasknamenum := len(taskname)
	if tasknamenum <= stringSize {
		pandingnum = stringSize - tasknamenum
		remainder = pandingnum % 2
		tempnum = pandingnum / 2
	}

	if tempnum != 0 {
		for i := 1; i < tempnum; i++ {
			leftspace = leftspace + " "
			rightspace = rightspace + " "
		}

	}

	if remainder != 0 {
		for i := 1; i < remainder; i++ {
			remainderspace = remainderspace + " "
		}
	}

	if remainder > 0 {
		result = fmt.Sprintf("%s%s%s%s", leftspace, taskname, rightspace, remainderspace)
	} else {
		result = fmt.Sprintf("%s%s%s", leftspace, taskname, rightspace)
	}

	return result

}

func (bar *Bar) ChangeTaskName(taskname string) {
	bar.taskname = tasknameFormat(taskname)
}

func (bar *Bar) Finish() {
	bar.rate = ""
	fmt.Println()
}
