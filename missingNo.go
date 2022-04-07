package main

import "fmt" 

func main(){
     a := []int { 1, 2, 3, 4, 5 }
     b := []int { 2, 3, 1, 0, 5 }
    var n int = len(a)
    var m int = len(b)
    findMissing(a, b, n, m);
}

func findMissing(a []int,b []int,n int ,m int ) {
	var i int				 
    for  i = 0; i < n; i++ {
        var j int
        for j = 0; j < m; j++ {
            if (a[i] == b[j]){
                break;
			}
 
        if (j == m) {
			fmt.Println(a[i], " ");
		}
    }
}
