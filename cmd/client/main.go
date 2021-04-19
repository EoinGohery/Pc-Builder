package client

import (
	"CS4227/pkg/factory"
	"CS4227/pkg/interceptor"
	"fmt"
	"log"
	"os"
)

var Tower factory.Tower

func partPicker(picker int, built bool) []factory.Component {
	var manager interceptor.InterceptorManager
	var logger interceptor.LoggingInterceptor
	var filterer interceptor.FilterInterceptor
	var part []factory.Component
	manager.SetFilter(filterer)
	manager.SetLogging(logger)

	if picker == 1 {
		part := manager.Execute("cpu", "")
		for i := range part {
			fmt.Print(part[i].PrintIDString())
		}
		return part

	} else if picker == 2 {
		//gpu
		part := manager.Execute("gpu", "")
		for i := range part {
			fmt.Print(part[i].PrintIDString())
		}
		return part

	} else if picker == 3 {
		//psu
		part := manager.Execute("psu", "")
		for i := range part {
			fmt.Print(part[i].PrintIDString())
		}
		return part

	} else if picker == 4 {
		//ram
		part := manager.Execute("ram", "")
		for i := range part {
			fmt.Print(part[i].PrintIDString())
		}
		return part
	} else if picker == 5 {
		//mobo
		part := manager.Execute("mbd", "")
		for i := range part {
			fmt.Print(part[i].PrintIDString())
		}
		return part

	} else if picker == 6 {
		//drive
		part := manager.Execute("drive", "")
		for i := range part {
			fmt.Print(part[i].PrintIDString())
		}
		return part

	} else if picker == 7 {
		fmt.Println("Build Finished/Cancelled, Exiting")
		os.Exit(0)
	}
	return part

}

func getpartbyid(picker int, id int, partlist []factory.Component) factory.Component {
	var part factory.Component

	if picker == 1 {
		//cpu
		partlist[id].Print()
		Tower.Add(partlist[id])
		return partlist[id]
	} else if picker == 2 {
		//gpu
		partlist[id].Print()
		Tower.Add(partlist[id])
		return partlist[id]
	} else if picker == 3 {
		//psu
		partlist[id].Print()
		Tower.Add(partlist[id])
		return partlist[id]
	} else if picker == 4 {
		//ram
		partlist[id].Print()
		Tower.Add(partlist[id])
		return partlist[id]
	} else if picker == 5 {
		//mobo
		partlist[id].Print()
		Tower.Add(partlist[id])
		return partlist[id]
	} else if picker == 6 {
		//drive
		partlist[id].Print()
		Tower.Add(partlist[id])

		return partlist[id]

	}
	return part
}

func getUserInput() bool {
	var picker, id int
	var built bool
	var part []factory.Component
	fmt.Println("")
	fmt.Println("PC Parts -  1: CPU, 2: GPU, 3: PSU, 4: Ram, 5: Motherboard, 6: Drive; 7: Finish ")
	fmt.Scanln(&picker)
	if picker == 7 {
		return false
	} else {
		part = partPicker(picker, built)
	}
	fmt.Println("\nPick One part!")
	fmt.Scanln(&id)
	getpartbyid(picker, id-1, part)
	return true
}

//Run client
func Run() {

	fmt.Println("Welcome to the PC Builder")
	// what part is required
	for {

		if !getUserInput() {
			break
		}

	}

	// Composite and Prototype in practise
	var manager interceptor.InterceptorManager
	var logger interceptor.LoggingInterceptor
	var filterer interceptor.FilterInterceptor
	manager.SetFilter(filterer)
	manager.SetLogging(logger)

	Tower.Print()
	cloneTower := Tower.Clone()
	cloneTower.Print()

	//write clone to file
	f, err := os.Create("TowerBuildClone.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(cloneTower.ToString())

	if err2 != nil {
		log.Fatal(err2)
	}

}
