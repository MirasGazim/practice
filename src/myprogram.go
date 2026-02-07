package main

import (
	"fmt"
)

type websiteinfo struct {
	name      string
	url       string
	ageinyear int
}

type Work interface {
	Canwork() bool
}

type miras struct {
	name string
}

type Person struct {
	Name string
	Age  int
}

type Speaker interface {
	Speak() string
}

// func (p Person) String() string {
// 	result := fmt.Sprintf("%s is %d years old", p.Name, p.Age)
// 	return result
// }

type User struct {
	name string
}

type Robot struct {
}

type Dog struct {
	name string
}

func (Dog) Speak() string { return "Woof" }

func MakeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func (web websiteinfo) basicInfo() {
	fmt.Println(web.name)
	fmt.Println(web.url)
	fmt.Println(web.ageinyear)
}

// func searchInsert(nums []int, target int) int {
// 	index_of_target := 0
// 	count := 0
// 	len_of_s := len(nums)
// 	var slice []int = nums
// 	for true {
// 		if len_of_s <= 1 {
// 			return index_of_target
// 		}
// 		if slice[len_of_s] == target {
// 			index_of_target = len_of_s
// 			break
// 		} else if slice[len_of_s] < target {
// 			slice = slice[(len_of_s + 1):]
// 		} else if slice[len_of_s] > target {
// 			slice = slice[:len_of_s]
// 		}
// 		count++
// 		len_of_s /= 2
// 	}
// 	return index_of_target
// }

// Node
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert
func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search
func (n *Node) Search(k int) bool {

	if n == nil {
		return false
	}
	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}
	return true
}

type youtube struct {
	websiteinfo
	firstvideo string
	lenofvideo int
}

func (yu youtube) Readstruct() {
	fmt.Println(yu.name)
	fmt.Println(yu.ageinyear)
	fmt.Println(yu.firstvideo)
}

type IINs struct {
	Person
	mir bool
}

func (i IINs) Canwork() bool {
	if i.mir == true {
		return true
	} else {
		return false
	}
}

type Builder interface {
	Build()
}

type Building struct {
	Builder
	Name string
}

type BrickBuilder struct {
	Person
}

type WoodBuilder struct {
	Person
}

type AZZZZ struct {
	Name string
	From string
	uni  string
}

func (w WoodBuilder) Build() {
	fmt.Println("i work with Wood")
}

func (b BrickBuilder) Build() {
	fmt.Println("i work with brick")
}

func (b AZZZZ) Build() {
	fmt.Println("i work with brick")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	var result []float64
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		sum := 0
		count := 0

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			count++

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, float64(sum)/float64(count))

	}
	return result
}

func Sum_of_val(key *Node) int {
	if key == nil {
		return 0
	}
	var sum int
	address := []*Node{key}

	for len(address) > 0 {
		n := address[0]
		address = address[1:]
		sum += n.Key

		if n.Left != nil {
			address = append(address, n.Left)
		}
		if n.Right != nil {
			address = append(address, n.Right)
		}
	}
	return sum
}

func exercise1(a int, b int, c int) int {
	return a * b * c
}

func exercise2(number int) int {
	if number > 1 {
		return number * exercise2(number-1)
	} else {
		return 1
	}
}

func exercise3(a int, b int) int {
	return exercise2(a) / exercise2(a-b)
}

func exercise4(programmers int, place int) int {
	return exercise2(programmers) / (exercise2(programmers-place) * exercise2(place))
}

func main() {
	fmt.Println(exercise4(12, 4))
}

// func (web websiteinfo) basicInfo(o) {
// 	fmt.Println("name of website: ", web.name)
// 	fmt.Println("Url of website: ", web.url)
// 	fmt.Println("How old is this website: ", web.ageinyear)
// }
