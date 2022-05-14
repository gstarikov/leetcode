package ms_other

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberToWords(t *testing.T) {
	tests := []struct {
		num int
		out string
	}{
		{num: 123, out: "One Hundred Twenty Three"},                                                                                      //0
		{num: 12345, out: "Twelve Thousand Three Hundred Forty Five"},                                                                    //1
		{num: 1234567, out: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},                                     //2
		{num: 1234567891, out: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"}, //3
		{num: 0, out: "Zero"},              //4
		{num: 20, out: "Twenty"},           //5
		{num: 1000000, out: "One Million"}, //6
	}

	for i, test := range tests {
		res := numberToWords(test.num)
		assert.Equalf(t, test.out, res, "case(%d) %d != %s", i, test.num, res)
	}
}
