package main

import "fmt"

//func canFinish(numCourses int, prerequisites [][]int) bool {
//	// Made a map of the course that will be taken
//	courses := map[int][]int{}
//	courseTaken := 0
//
//	// Loop through the prerequisites list
//	for _, prerequisite := range prerequisites {
//		course := prerequisite[0]
//		requiredCourse := prerequisite[1]
//
//		// Check if the course taken is already existed or not
//		if _, ok := courses[course]; !ok {
//			// If it hasn't, add to the map
//			courses[course] = append(courses[course], requiredCourse)
//		} else {
//			for _, c := range courses[course] {
//				if c == requiredCourse {
//					return false
//				}
//			}
//			courses[course] = append(courses[course], requiredCourse)
//		}
//		courseTaken++
//		if courseTaken > numCourses {
//			return false
//		}
//	}
//	return true
//}

func canFinish(numCourses int, prerequisites [][]int) bool {
	// Made a map of the course that will be taken
	courses := map[int][]int{}
	courseTaken := 0

	// Loop through the prerequisites list
	for _, prerequisite := range prerequisites {
		course := prerequisite[0]
		requiredCourse := prerequisite[1]

		// Check if the course taken is already existed or not
		if _, ok := courses[course]; !ok {
			// If it hasn't, add to the map
			courses[course] = append(courses[course], requiredCourse)
		} else {
			for _, c := range courses[course] {
				if c == requiredCourse {
					continue
				}
			}
			courses[course] = append(courses[course], requiredCourse)
		}
	}

	var checkViableCourse func(c int) bool
	checkViableCourse = func(c int) bool {
		if len(courses[c]) != 0 {
			courseTaken++
			for _, course := range courses[c] {
				if courseTaken > numCourses || !checkViableCourse(course) {
					return false
				}
				if len(courses[c]) != 0 {
					courses[c] = courses[c][1:]
				}
			}
		}
		return true
	}

	for _, prereq := range courses {
		for _, course := range prereq {
			if !checkViableCourse(course) {
				return false
			}
		}
	}

	return true
}

func main() {
	numCourses := 7
	//prerequisites := [][]int{{1, 0}, {0, 3}, {0, 2}, {3, 2}, {2, 5}, {4, 5}, {5, 6}, {2, 4}}
	prerequisites := [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {3, 4}}
	fmt.Println(canFinish(numCourses, prerequisites))
}
