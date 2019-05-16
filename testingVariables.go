package main

var smallestArray = []int {2}
var smallerArray = []int {2, 2}
var smallArray = []int {2, 1, 3, 5, 3, 2}

var a = [][]int {
	{1, 2, 3} ,   /*  initializers for row indexed by 0 */
	{4, 7, 9} ,   /*  initializers for row indexed by 1 */
	{11, 13, 15},
}
var b = [][]int {
	{1},
}

var c = [][]int {
	{10,9,6,3,7},
	{6,10,2,9,7},
	{7,6,3,8,2},
	{8,9,7,9,9},
	{6,8,6,8,2},
}

 var d = [][]int {
	{63,68,22,93,86,23,96,73,81,65,23,39,4,63,4,64,43,99,53,71,66,26,31,54,70,91,57,37,60,8,68,68,16,47,37,88,62,20,48,76,69,7,40,90,34,30,9,4,77,83},
	{76,81,6,35,69,50,15,73,30,51,35,17,90,90,65,25,70,80,39,55,64,34,90,78,47,50,33,54,29,66,14,51,49,16,7,82,74,73,9,22,41,59,39,67,64,93,34,55,47,46},
	{79,19,5,46,33,33,69,98,16,71,70,25,19,92,82,24,95,97,33,24,14,94,93,63,4,86,40,70,22,25,36,23,28,76,46,60,12,29,75,14,36,37,39,71,18,27,71,78,38,93},
	{43,65,63,93,40,83,65,86,62,59,15,97,100,74,78,41,2,61,87,24,41,71,18,99,77,78,3,100,51,40,87,65,95,28,14,94,88,2,3,17,43,81,3,92,55,84,76,90,72,32},
	{3,42,85,55,74,51,59,17,75,75,87,3,79,90,74,29,71,7,13,2,18,40,98,4,49,11,64,7,40,8,29,34,96,59,65,37,47,50,31,42,83,17,4,73,44,16,78,91,50,15},
	{24,17,51,14,62,5,93,1,82,8,75,100,59,88,40,4,37,79,77,81,20,52,62,13,99,41,92,62,17,12,31,35,38,24,15,20,53,32,85,57,55,46,93,82,86,18,9,58,68,24},
	{92,45,85,25,69,11,95,53,18,11,88,37,67,35,18,15,71,13,47,14,32,25,64,58,4,8,32,48,99,85,74,1,88,6,52,12,56,83,72,81,27,72,94,25,82,10,2,33,24,83},
	{87,10,72,28,75,25,15,82,77,83,37,46,99,21,47,59,80,34,79,30,2,30,5,6,42,64,85,48,39,87,55,60,3,40,76,59,33,98,74,76,55,91,3,26,70,81,65,99,29,48},
	{15,47,10,37,1,86,32,63,89,68,92,85,28,56,57,15,40,36,13,51,28,23,56,77,72,99,39,37,46,20,14,23,69,7,63,11,11,65,29,71,93,4,77,15,33,7,46,2,3,92},
	{40,64,20,79,49,69,46,73,92,87,98,30,100,69,51,87,55,21,19,63,6,10,65,96,88,64,84,10,3,12,55,32,99,84,30,13,88,19,87,69,74,11,6,1,41,82,70,1,94,64},
	{14,97,35,60,98,9,30,89,64,48,21,48,95,37,19,15,93,12,2,14,63,67,95,90,62,20,86,92,50,32,12,67,38,6,29,21,83,92,84,20,7,33,7,3,63,89,24,12,29,12},
	{75,39,62,50,20,94,1,24,60,96,64,97,58,64,24,24,95,46,96,67,24,29,36,36,62,1,14,49,44,53,51,8,56,92,60,17,11,70,98,19,42,94,21,27,99,96,55,44,60,61},
	{44,61,16,69,26,40,79,80,63,71,25,68,13,54,100,80,50,89,83,94,29,26,34,24,79,6,20,39,20,50,96,43,1,34,26,9,82,92,46,3,68,68,57,72,73,87,38,6,68,19},
	{4,42,39,9,51,69,93,24,70,16,43,61,58,1,61,34,50,13,87,97,14,33,37,64,46,14,4,8,9,14,88,39,84,92,46,27,58,21,85,26,32,39,55,20,97,82,53,12,39,67},
	{61,70,41,52,46,88,13,1,90,2,40,85,75,31,52,94,84,36,28,8,27,61,68,83,87,80,4,68,63,76,52,68,64,64,32,30,53,12,77,4,61,47,92,92,41,41,16,87,52,87},
	{53,98,74,6,36,49,94,11,67,51,67,49,92,72,60,59,16,66,87,87,49,13,35,70,77,72,37,8,78,42,80,87,68,46,86,47,4,21,18,52,69,12,36,32,30,43,18,60,17,67},
	{13,78,66,62,84,74,87,87,96,4,61,60,48,74,14,78,89,54,59,28,32,63,35,47,90,37,23,16,92,97,87,65,19,50,98,40,39,20,41,81,9,23,85,85,86,15,11,3,54,91},
	{24,41,18,18,39,48,94,16,56,7,56,9,92,58,100,14,58,17,16,66,67,10,40,29,10,71,12,67,18,3,7,17,12,45,49,53,41,52,90,86,37,55,6,83,26,8,61,54,72,26},
	{44,73,100,95,28,69,6,77,95,2,94,44,56,97,98,19,56,78,53,20,23,25,11,52,43,32,79,57,2,53,49,32,73,89,48,53,55,87,61,66,87,46,7,21,73,9,60,65,94,57},
	{92,32,73,49,69,96,82,56,56,91,75,73,86,24,92,57,83,76,24,83,76,40,2,15,34,3,22,47,4,1,36,95,75,10,45,97,95,2,83,50,82,49,56,33,70,3,60,29,23,46},
	{27,23,9,60,26,45,74,38,72,72,29,9,94,34,22,26,19,98,68,58,70,17,100,88,4,6,87,70,2,44,97,19,35,16,20,20,93,79,60,79,38,40,85,31,14,34,75,25,28,38},
	{98,60,37,56,23,16,39,28,25,52,36,98,53,15,17,33,50,77,5,79,54,65,1,7,38,44,99,33,53,3,73,31,37,96,60,79,73,83,14,16,25,16,3,32,58,20,83,79,35,100},
	{26,17,3,27,48,81,63,37,53,80,98,8,37,61,82,49,34,87,23,27,99,76,100,55,55,91,94,89,20,36,88,94,28,91,3,87,11,26,25,73,95,36,86,99,95,36,77,95,40,51},
	{91,71,6,86,27,84,65,92,86,64,21,57,80,70,37,28,52,64,98,14,99,13,10,86,30,85,82,78,8,28,54,47,84,61,45,61,16,19,50,40,53,89,10,92,5,57,99,9,7,93},
	{62,91,15,97,55,62,73,71,57,56,62,31,39,54,31,2,12,7,67,97,59,86,98,82,8,30,62,61,15,87,46,46,12,79,46,31,18,79,20,22,54,72,62,72,39,66,27,96,16,15},
	{81,33,56,36,60,54,11,71,88,100,55,95,97,12,9,62,36,3,14,6,48,91,40,29,24,41,99,85,48,71,26,22,84,2,7,54,73,91,59,57,56,20,78,37,38,56,9,36,67,20},
	{13,93,98,77,8,31,86,72,6,87,67,49,45,9,77,15,55,79,82,26,77,77,4,23,53,72,48,76,54,58,35,12,97,90,7,25,61,54,61,83,28,64,44,91,18,80,78,28,19,28},
	{74,8,77,59,14,10,72,79,81,76,40,94,93,96,71,37,44,63,18,95,15,47,26,87,73,19,90,74,79,28,11,94,33,91,8,95,44,41,36,57,43,50,74,30,88,8,86,91,21,9},
	{45,36,31,93,79,47,84,73,54,55,5,33,62,22,67,16,38,31,20,95,49,68,29,16,14,49,36,97,2,15,75,40,54,6,80,60,15,93,36,83,95,71,99,57,6,11,52,51,95,92},
	{72,42,59,5,79,72,95,52,88,4,95,92,61,5,60,83,57,38,51,12,63,96,71,59,43,4,67,90,37,17,86,13,55,14,86,36,54,52,38,15,94,37,99,55,64,48,30,89,67,86},
	{36,77,74,33,67,42,76,69,100,94,60,100,97,50,95,2,68,100,42,87,79,54,65,54,71,14,95,54,73,70,36,10,29,78,3,64,1,47,66,12,47,39,10,86,71,24,29,65,58,99},
	{28,20,23,52,92,88,96,61,59,12,61,57,78,44,70,57,80,95,80,97,34,34,94,87,89,49,57,98,66,79,29,48,93,24,91,32,92,41,31,96,73,86,55,36,5,28,57,71,87,65},
	{46,56,57,13,74,60,54,95,76,89,53,81,81,80,85,3,61,27,43,20,79,43,79,18,27,95,28,39,44,76,59,70,34,70,51,48,68,9,68,84,84,61,66,50,66,80,43,4,32,61},
	{93,9,78,58,20,59,75,98,45,61,35,46,84,25,68,86,65,46,2,7,13,16,9,60,2,25,87,70,98,75,7,98,50,66,91,79,83,97,68,27,31,87,33,81,54,30,83,31,70,88},
	{76,24,99,60,31,83,65,52,31,64,18,29,100,81,6,45,95,86,22,32,93,42,71,10,16,25,77,92,50,21,55,34,17,34,20,39,1,38,68,98,49,51,62,82,37,92,18,50,24,22},
	{70,88,86,10,19,3,39,99,61,54,73,83,25,96,97,58,64,10,68,25,26,56,74,39,21,75,35,28,85,80,7,57,34,34,67,93,68,47,10,89,23,73,61,8,94,16,83,12,36,13},
	{100,40,31,23,86,11,6,85,47,44,84,66,27,33,81,45,3,60,7,38,90,47,51,35,12,9,92,7,22,18,93,33,57,14,58,75,45,33,8,23,90,4,58,58,10,94,53,98,71,7},
	{46,57,7,43,8,28,67,74,61,59,1,30,42,87,20,10,9,94,19,81,79,34,17,3,30,28,74,5,43,51,90,88,32,62,75,11,82,56,76,74,16,11,15,98,74,6,10,42,79,36},
	{16,6,8,4,64,18,81,84,55,72,82,75,34,80,76,2,4,74,17,68,22,34,7,23,97,6,57,43,100,92,55,15,66,49,56,98,91,21,67,87,61,25,45,70,15,32,5,78,24,55},
	{52,76,79,99,34,46,7,82,81,97,48,63,88,46,52,49,70,66,53,22,81,92,90,28,39,50,63,81,86,17,8,90,99,97,64,69,79,78,54,89,34,89,66,48,96,3,6,29,95,60},
	{12,99,86,70,31,97,83,45,79,32,19,85,19,51,45,82,57,44,67,67,20,5,12,67,53,91,76,81,33,64,75,48,72,51,52,29,83,99,12,82,58,89,30,98,68,37,78,67,25,2},
	{94,4,31,30,61,89,83,44,21,16,38,98,87,32,87,98,36,90,67,94,71,57,40,74,62,53,30,7,79,69,48,74,51,3,33,75,70,66,3,1,44,20,93,39,76,42,64,99,50,78},
	{65,64,48,38,22,59,13,42,84,24,97,45,72,51,77,96,18,70,11,29,27,35,89,6,43,88,39,85,29,100,44,8,11,50,35,37,17,5,70,86,40,25,56,42,64,67,12,24,62,61},
	{58,63,36,64,60,84,83,40,70,34,15,32,47,14,15,10,74,87,12,85,3,44,67,13,6,55,32,40,31,33,74,87,67,32,88,93,91,45,48,67,99,62,54,85,100,27,49,30,55,42},
	{58,78,62,36,22,52,33,72,49,13,75,43,25,2,91,23,40,96,21,41,57,24,7,6,80,82,45,29,2,43,42,47,19,6,70,95,89,80,81,4,95,33,3,90,32,25,38,79,27,42},
	{95,57,28,46,10,1,29,74,87,55,1,17,28,91,64,98,97,65,57,78,98,92,22,48,55,85,67,44,13,93,38,62,23,28,96,33,68,32,82,38,10,44,51,10,53,49,2,8,96,50},
	{98,8,40,17,88,61,7,47,65,16,65,40,2,53,1,42,51,45,34,77,32,58,59,43,69,9,39,36,72,64,86,87,79,99,66,14,12,52,81,15,34,100,61,29,53,46,9,76,84,72},
	{1,66,91,51,32,75,33,76,91,75,72,85,56,20,79,71,79,85,14,92,12,20,18,17,51,61,48,91,19,23,21,60,7,39,71,81,18,17,18,24,88,43,2,25,52,84,18,40,46,58},
	{33,96,93,47,55,89,22,76,46,23,87,59,72,60,35,1,12,27,21,81,69,22,38,36,37,68,90,39,51,100,52,69,48,43,21,8,16,51,9,64,26,72,94,90,45,96,45,81,88,16},
	{36,22,19,84,33,42,72,34,37,45,83,84,57,38,64,15,89,80,37,8,84,70,56,45,35,60,4,73,31,36,26,18,17,87,34,85,2,67,18,63,90,84,20,68,66,4,56,29,24,7},
}