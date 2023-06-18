package normal

import (
	"log"
	"testing"
)

func Test1(t *testing.T) {
	AddNum()
	log.Printf("\n\t[ Test1 ] num : %v\n", Num)
}

func Test2(t *testing.T) {
	AddNum()
	log.Printf("\n\t[ Test2 ] num : %v\n", Num)
}
