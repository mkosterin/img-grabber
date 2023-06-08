package logic

import (
	"fmt"
	"img-grubber/internal/counter"
	"img-grubber/internal/reader"
)

func ImageCounter(arr *reader.Data) {
	img := *counter.NewImg()
	for i := 0; i < len(arr.Text); i++ {
		err, key, value := counter.Parse(arr.Text[i])
		if err == nil {
			img[key] = value
		}
	}

	for key := range img {
		fmt.Printf("%v\n", key)
	}
}
