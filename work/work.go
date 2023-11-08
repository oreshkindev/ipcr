package work

import (
	"log"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/oresdev/ipcr/converter"
)

type Worker interface {
	Run(id int)
}

func New(i, o string, queue chan string, quality int, converter converter.Converter) Worker {
	return &workerI{i, o, queue, quality, converter}
}

type workerI struct {
	i         string
	o         string
	queue     chan string
	quality   int
	converter converter.Converter
}

func (w *workerI) Run(id int) {
	log.Printf("run %d worker", id)

	// performing tasks from the task pool
	for q := range w.queue {
		time.Sleep(1 * time.Second)

		log.Printf("converting %s file", q)
		// prepare path: joins args input catalog and file
		i := path.Join(w.i, q)
		// replace file extension
		t := strings.ReplaceAll(q, strings.ToLower(filepath.Ext(q)), ".webp")
		// joins any number of path elements
		o := path.Join(w.o, t)
		// conver file
		if err := w.converter.Convert(i, o, w.quality); err != nil {
			log.Printf("failed to convert %s; %s", q, err)
		}
	}

	log.Printf("finish %d worker", id)
}
