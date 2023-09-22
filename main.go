package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"sort"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// Game type constructor init
type Game struct {
}

//Game update loop
func (g *Game) Update() error {
	return nil
}

// Partition function for quicksort
func partition(arr []int, low int, high int) int {
	pivot := arr[high]

	i := low - 1

	for j := low; j <= high; j++ {
		time.Sleep(time.Millisecond * 10)
		redIndex = pivot
		if arr[j] < pivot {

			i++
			arr[i], arr[j] = arr[j], arr[i]

		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return (i + 1)
}

// Quicksort implementation
func quickSort(arr []int, low int, high int) {

	if low < high {
		partitionIndex := partition(arr, low, high)
		redIndex = partitionIndex
		time.Sleep(time.Millisecond * 50)
		quickSort(arr, low, partitionIndex-1)
		quickSort(arr, partitionIndex+1, high)

	}
}

// Bubble sort implementation
func bubbleSort(arr []int) {

	for i := 0; i < len(arr)-1; i++ {
		redIndex = i
		time.Sleep(time.Microsecond * 100)
		for j := 0; j < ((len(arr) - 1) - i); j++ {
			redIndex = j
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	if (len(arr) - 1) >= 1 {
		bubbleSort(arr[:len(arr)-1])
	}

}

// Insertion sort implementation
func insertionSort(arr []int) {

	for i := 1; i < len(arr); i++ {
		time.Sleep(time.Millisecond * 2)
		key := arr[i]

		j := i - 1
		redIndex = j
		for j >= 0 && arr[j] > key {
			time.Sleep(time.Millisecond * 2)
			arr[j+1] = arr[j]
			j = j - 1
			redIndex = j
		}
		arr[j+1] = key
	}

}

// Selection sort implementation

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		redIndex = i
		time.Sleep(time.Millisecond * 20)
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {

				redIndex = j
				time.Sleep(time.Millisecond * 20)
				min = j
			}
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]

		}

	}

}

// Merge sort implementation

func mergeSort(arr []int) {
	leng := 1
	for leng < len(arr) {

		i := 0
		for i < len(arr) {
			lone := i
			rone := i + leng - 1
			ltwo := i + leng
			rtwo := i + 2*leng - 1
			if ltwo >= len(arr) {
				break
			}
			if rtwo >= len(arr) {
				rtwo = len(arr) - 1
			}
			redIndex = i
			time.Sleep(time.Millisecond * 50)
			temp := merge(arr, lone, rone, ltwo, rtwo)

			for j := 0; j < (rtwo - lone + 1); j++ {
				arr[i+j] = temp[j]
			}
			i = i + 2*leng

		}
		leng = 2 * leng
	}
	redIndex = -1

}

// Merge function for merge sort
func merge(arr []int, lone int, rone int, ltwo int, rtwo int) []int {

	var temp = make([]int, len(arr))
	index := 0
	for lone <= rone && ltwo <= rtwo {
		if arr[lone] <= arr[ltwo] {
			temp[index] = arr[lone]
			index++
			lone++
			redIndex = lone
			time.Sleep(time.Millisecond * 5)

		} else {
			temp[index] = arr[ltwo]
			index++
			ltwo++
			redIndex = ltwo
			time.Sleep(time.Millisecond * 5)

		}

	}

	for lone <= rone {
		temp[index] = arr[lone]
		index++
		lone++
		redIndex = lone
		time.Sleep(time.Millisecond * 5)

	}

	for ltwo <= rtwo {
		temp[index] = arr[ltwo]
		index++
		ltwo++
		redIndex = ltwo
		time.Sleep(time.Millisecond * 5)

	}

	return temp
}

// Heap sort implementation
func heapSort(arr []int) {
	length := len(arr)
	i := int(length/2 - 1)
	k := length - 1

	for i >= 0 {

		heapify(arr, length, i)

		i--
		redIndex = i
		time.Sleep(time.Millisecond * 5)
	}

	for k >= 0 {
		arr[0], arr[k] = arr[k], arr[0]
		heapify(arr, k, 0)
		k--
		redIndex = k
		time.Sleep(time.Millisecond * 5)
	}

}

// Heapify function for heap sort
func heapify(arr []int, length int, i int) []int {
	largest := i
	left := i*2 + 1
	right := left + 1
	redIndex = largest

	if left < length && arr[left] > arr[largest] {
		largest = left
	}

	if right < length && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, length, largest)
	}
	redIndex = largest
	time.Sleep(time.Millisecond * 5)

	return arr

}

// Counting sort for radix sort
func countingSort(arr []int, size int, div int) {
	output := make([]int, size)
	count := [10]int{0}
	for i := 0; i < size; i++ {

		count[(arr[i]/div)%10]++
		redIndex = count[(arr[i]/div)%10]
		time.Sleep(time.Millisecond * 10)
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]

		redIndex = count[i]
		time.Sleep(time.Millisecond * 10)

	}

	for i := size - 1; i >= 0; i-- {
		output[count[(arr[i]/div)%10]-1] = arr[i]
		count[(arr[i]/div)%10]--
		redIndex = count[(arr[i]/div)%10]
		time.Sleep(time.Millisecond * 10)
	}
	for i := 0; i < size; i++ {
		arr[i] = output[i]
		redIndex = arr[i]
		time.Sleep(time.Millisecond * 10)
	}
}

// Radix sort implementation
func radixSort(arr []int, size int) {
	m := getMax(arr, size)
	for div := 1; m/div > 0; div *= 10 {

		countingSort(arr, size, div)
	}
}

func findMax(arr []int, n int) int {
	if n == 1 {
		return arr[0]
	}
	return int(math.Max(float64(arr[n-1]), float64(findMax(arr, n-1))))
}
func bucketSort(arr []int, n int) {
	max := findMax(arr, n)
	b := make([][]int, n)
	for i := 0; i < n; i++ {
		bi := n * arr[i] / (max + 1)
		b[bi] = append(b[bi], arr[i])
		redIndex = bi
		time.Sleep(time.Millisecond * 10)

	}

	for i := 0; i < n; i++ {
		insertionSort(b[i])
	}

	index := -1
	for i := 0; i < n; i++ {
		redIndex = i
		time.Sleep(time.Millisecond * 50)
		for j := 0; j < len(b[i]); j++ {
			index++
			arr[index] = b[i][j]
			redIndex = index
			time.Sleep(time.Millisecond * 50)

		}
	}
}

// Get max for radix sort
func getMax(arr []int, size int) int {
	max := arr[0]
	for i := 0; i < size; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// Shell sort implementation
func shellSort(arr []int) {

	n := len(arr)

	for gap := n / 2; gap > 0; gap /= 2 {

		redIndex = gap
		time.Sleep(time.Millisecond * 5)
		for i := gap; i < n; i += 1 {

			temp := arr[i]
			redIndex = i
			time.Sleep(time.Millisecond * 5)

			var j int
			for j = i; j >= gap && arr[j-gap] > temp; j -= gap {

				arr[j] = arr[j-gap]
				redIndex = j
				time.Sleep(time.Millisecond * 5)
			}
			redIndex = j
			time.Sleep(time.Millisecond * 5)
			arr[j] = temp
		}
	}

	redIndex = -1

}

// Is array sorted check
func isSorted(arr []int) bool {
	return sort.IntsAreSorted(arr)
}

// Is array empty check
func isEmpty(arr []int) bool {
	return len(arr) == 0
}

// Variable init
var (
	mPlusFont     font.Face
	mPlusPlusFont font.Face
	numarr        []int
	colsize       int
	colheight     int
	colsep        int
	gamew         int
	gameh         int
	numamt        int
	redIndex      int
	inMenu        bool
	isSorting     bool
	debounceMenu  int

	//sortAlg map[int]func()
)

// Game init function (set vars, fonts, etc)
func init() {

	// Init variables
	debounceMenu = 0
	gamew = 900
	gameh = 500
	inMenu = true
	redIndex = -1
	isSorting = false
	colsep = 2
	colsize = 16
	numamt = gamew / (colsep + colsize)
	colheight = gameh / numamt

	numarr = []int{}
	for i := 0; i <= numamt; i++ {
		numarr = append(numarr, rand.Intn(numamt))

	}

	// Parse font with opentype
	myfont, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	// Create fonts and handle errors
	const dpi = 72
	mPlusFont, err = opentype.NewFace(myfont, &opentype.FaceOptions{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal("Well shit")
	}
	mPlusPlusFont, err = opentype.NewFace(myfont, &opentype.FaceOptions{
		Size:    32,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal("Well shit")
	}

}

// Main game draw loop
func (g *Game) Draw(screen *ebiten.Image) {
	// Update isSorting value
	debounceMenu++
	if isSorted(numarr) {
		isSorting = false
	}

	// Draw menu top
	text.Draw(screen, "[1] Quicksort [2] Bubble Sort [3] Insertion Sort [4] Merge Sort [5] Selection Sort [6] Heap Sort [7] Radix sort [8] Bucket sort [9] Shell sort", mPlusFont, 20, 25, color.White)
	text.Draw(screen, "[C] Clear Array [A] Add data", mPlusFont, 20, 55, color.White)
	text.Draw(screen, "[W] Hide/Show visualisation", mPlusFont, 20, 85, color.White)
	// Draw depending on menu
	switch inMenu {
	case true:
		{
			// Draw menu
			text.Draw(screen, "Sorting Algorithms ([W] View Visualisation)", mPlusPlusFont, 100, 250, color.White)
		}

	case false:
		{
			// Draw visualisation
			for i, val := range numarr {

				posx := float64(((i * (colsize + colsep)) + 10) - colsize*2)
				posy := float64((gameh - 10) - val*colheight)

				if i == redIndex {

					ebitenutil.DrawRect(screen, posx, posy, float64(colsize), float64(val*colheight), color.RGBA{255, 0, 0, 255})

				} else {
					ebitenutil.DrawRect(screen, posx, posy, float64(colsize), float64(val*colheight), color.White)

				}

				//text.Draw(screen, strconv.Itoa(val), mPlusFont, int(posx), gameh-15, color.Black)
			}
		}
	}

	// Menu switch key

	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if debounceMenu >= 10 {
			debounceMenu = 0
			inMenu = !inMenu
		}

	}
	// Clear array key
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		if isSorted(numarr) {
			numarr = []int{}
			redIndex = -1
		}

	}

	// if inpututil.IsKeyJustPressed(ebiten.KeyB) {
	// 	if !isEmpty(numarr) && !isSorted(numarr) && !isSorting {
	// 		isSorting = true
	// 		go bogoSort(numarr)
	// 	}
	// }
	// Append array key
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if isEmpty(numarr) {
			for i := 0; i <= numamt; i++ {
				numarr = append(numarr, rand.Intn(numamt))
			}
			redIndex = -1
		}

	}
	// All sorts keys
	if inpututil.IsKeyJustPressed(ebiten.KeyDigit1) {
		if !isEmpty(numarr) && !isSorted(numarr) && !isSorting {
			isSorting = true
			go quickSort(numarr, 0, len(numarr)-1)

		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDigit2) {
		if !isEmpty(numarr) && !isSorted(numarr) && !isSorting {
			isSorting = true
			go bubbleSort(numarr)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDigit3) {
		if !isEmpty(numarr) && !isSorted(numarr) && !isSorting {
			isSorting = true
			go insertionSort(numarr)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDigit4) {
		if !isEmpty(numarr) && !isSorted(numarr) && !isSorting {
			isSorting = true
			go mergeSort(numarr)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDigit5) {
		if !isEmpty(numarr) && !isSorted(numarr) && !isSorting {
			isSorting = true
			go selectionSort(numarr)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDigit6) {
		if !isEmpty(numarr) && !isSorted(numarr) && !isSorting {
			isSorting = true
			go heapSort(numarr)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDigit7) {
		if !isEmpty(numarr) && !isSorted(numarr) && !isSorting {
			isSorting = true
			go radixSort(numarr, len(numarr))
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDigit8) {
		if !isEmpty(numarr) && !isSorted(numarr) && !isSorting {
			isSorting = true
			go bucketSort(numarr, len(numarr))
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDigit9) {
		if !isEmpty(numarr) && !isSorted(numarr) && !isSorting {
			isSorting = true
			go shellSort(numarr)
		}
	}

}

// Set game screen height/width
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gamew, gameh
}

// Actual game/window init
func main() {

	ebiten.SetWindowSize(gamew, gameh)
	ebiten.SetWindowTitle(fmt.Sprintf("sorting visualisation, sorting %v objects", numamt))
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
