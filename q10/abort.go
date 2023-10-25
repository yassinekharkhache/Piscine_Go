package main

func Abort(a, b, c, d, e int) int {
	
	var arr  [5]int;
	arr[0] = a;
	arr[1] = b;
	arr[2] = c;
	arr[3] = d;
	arr[4] = e;
	i := 0;
	s := 0;
	size := 4;
	for(i <= size){
		for(i < size){
			if(arr[i] < arr[i + 1] ){
				s = arr[i];
				arr[i] = arr[i + 1];
				arr[i + 1] = s;
			}
			i++;
		}
		i = 0;
		size--;
	}
	return arr[2];
}
