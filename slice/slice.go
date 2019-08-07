package main

import "fmt"

func main(){

	months := []string{"","January","February","March","April","May","June","July","August","September","October","November","December"}
    fmt.Printf("months	cap: %d, len: %d\n", cap(months),len(months))
	Q2 := months[4:7]
	fmt.Printf("Q2	cap:%d, len: %d\n",cap(Q2),len(Q2))
	summer := months[6:9]
	fmt.Printf("summber	cap: %d, len: %d\n",cap(summer),len(summer))



	for _,s := range summer {
		for _,m := range months{
			if s == m{
				fmt.Printf("%s appears in both \n",s)
			}
		}
	}
}
