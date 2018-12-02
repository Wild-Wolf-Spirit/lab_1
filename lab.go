package main
import "fmt"
import "os"


func create_set() []string {
	var temp string
	var set []string
	fmt.Println("Введите множество из 5 чисел")
	for i:=0; i<5; i++{
		fmt.Scanln(&temp)
		set = append(set, temp)
	}
	return set
}

func intersection(set1 []string, set2 []string){
	var res_set []string
	
	i1 := 0
	j1 := 0
	for i1<5 {
		for j1<5 {
			if set1[i1] == set2[j1] {
				res_set = append(res_set, set1[i1])
				break 
				}	else {j1++}
			}
		i1++
		j1=0
		}
	
	fmt.Println("Результат пересечения:")
	fmt.Println(res_set)
}

func association(set1 []string, set2 []string){
	var res_set = set1
	var fl = 0

	i1 := 0
	j1 := 0
	for j1<5 {
		for i1<5 {
			if set2[j1] == set1[i1] {
				fl = 1
				break
				}	else {i1++}
			}
		if fl == 0 {
			res_set = append(res_set, set2[j1])
			}
		j1++
		i1 = 0
		fl = 0
		}
	
	fmt.Println("Результат объединения:")
	fmt.Println(res_set)
}

func division(set1 []string, set2 []string){
	var res_set []string
	var fl = 0

	i1 := 0
	j1 := 0
	for i1<5 {
		for j1<5 {
			if set1[i1] == set2[j1] {
				fl = 1
				break
				}	else {j1++}
			}		
		if fl == 0 {
			res_set = append(res_set, set1[i1])
			}
		fl = 0
		i1++
		j1=0
		}
	
	fmt.Println("Результат деления:")
	fmt.Println(res_set)
}
	
func main() {

	var set_container [][]string
	var key string

	s1:=0
	s2:=0
	lenth:=0

 	fmt.Println("1. Создать множество")
	fmt.Println("2. Вывести список множеств")
	fmt.Println("3. Пересечение")
	fmt.Println("4. Объединение")
	fmt.Println("5. Деление")
	fmt.Println("6. Выход")

	for true{
		fmt.Scanln(&key)
		switch(key) {
			case "1":
				{
				set_container = append(set_container, create_set())
				fmt.Println("Новое множество")
				fmt.Println(set_container[lenth])
				lenth++
				}
			case "2": 
				fmt.Println(set_container)
			case "3":
				{
				if lenth>0 {
					fmt.Println("Выберите два множества:")
					fmt.Scanln(&s1)
					fmt.Scanln(&s2)
					intersection(set_container[s1], set_container[s2])
					}	else {break}
				}
			case "4":
				{
				if lenth>0 {
					fmt.Println("Выберите два множества:")
					fmt.Scanln(&s1)
					fmt.Scanln(&s2)
					association(set_container[s1], set_container[s2])
					}	else {break}
				}
			case "5":
				{
				if lenth>0 {
					fmt.Println("Выберите два множества:")
					fmt.Scanln(&s1)
					fmt.Scanln(&s2)
					division(set_container[s1], set_container[s2])
					}	else {break}
				}
			case "6":
				os.Exit(3)
			}
	}
}
