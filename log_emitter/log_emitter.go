package le

import (
	"bufio"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

type LogItem struct {
	LOG       []string
	DEST      string
	CATEGORY  string
	TS_RE     string
	TS_FORMAT string
	INTERVAL  int
}

type LogEmitter struct {
	LOGS []LogItem
}

func (le *LogEmitter) LoadConf(path string) {
	f, err := ioutil.ReadFile(path)
	check(err)

	err = yaml.Unmarshal(f, le)
	check(err)
}

func (le *LogEmitter) appendLog(path string, s string) {
	dir, err := filepath.Abs(filepath.Dir(path))
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Printf("create new dir <%s>", dir)
		err = os.MkdirAll(dir, 0777)
		check(err)
		log.Printf("create new dir <%s> successfully", dir)
	}

	s += "\n"

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer f.Close()
	check(err)

	w := bufio.NewWriter(f)
	w.WriteString(s)
	w.Flush()
}

func (le *LogEmitter) updateTimestamp(raw string, ts_re string, ts_format string, t time.Time) string {
	r, err := regexp.Compile(ts_re)
	check(err)

	time_s := t.Format(ts_format)
	new_s := r.ReplaceAllString(raw, time_s)

	return new_s
}

func (le *LogEmitter) generateLog(item LogItem) {
	t := time.Now()

	logs := item.LOG
	ts_re := item.TS_RE
	ts_format := item.TS_FORMAT
	dest := item.DEST

	for _, s := range logs {
		new_s := le.updateTimestamp(s, ts_re, ts_format, t)
		le.appendLog(dest, new_s)
	}
}

func (le *LogEmitter) Execute() {
	quit := make(chan string)

	for _, v := range le.LOGS {
		ticker := time.NewTicker(time.Second * time.Duration(v.INTERVAL))
		go func(item LogItem) {
			for {
				select {
				case <-ticker.C:
					le.generateLog(item)
				case <-quit:
					ticker.Stop()
					return
				}
			}
		}(v)
	}

	<-quit
}
