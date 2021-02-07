package wingPi

import (
	"bufio"
	"github.com/shopspring/decimal"
	"log"
	"os"
)

func ReadTemp() float64 {
	p, _ := decimal.NewFromString("0.001")

	file, err := os.Open("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		log.Fatalf("open file failed: %s \n", err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	//i := 0
	for scanner.Scan() {
		temp, _ := decimal.NewFromString(scanner.Text())
		tempReal, _ := temp.Mul(p).Float64()
		return tempReal
	}
	return 100.00
}
