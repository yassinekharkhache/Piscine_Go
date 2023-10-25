package main
func IterativePower(nb int, power int) int {
	x := 1;
	if power == 0{
		return 1
	}
	if power < 0 || nb == 0{
		return 0;
	}
	for power > 0{
		x *= nb;
		power--;
	}
	return x;
}
