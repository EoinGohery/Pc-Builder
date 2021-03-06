package client

import (
	"CS4227/pkg/factory"
	"CS4227/pkg/interceptor"
	"fmt"
	"log"
	"os"
)

//build components
var Tower factory.Tower
var cpuPresent bool = false
var moboPresent bool = false
var gpuPresent bool = false
var psuPresent bool = false

var drivesPresent int = 0
var driveSlots int = 0
var ramPresent int = 0
var ramSlots int = 0

var filter string = ""
var skip bool = false

var motherboard factory.Component

//used when selecting the category of component to search
func partPicker(picker int, built bool) []factory.Component {
	var manager interceptor.InterceptorManager
	var logger interceptor.LoggingInterceptor
	var filterer interceptor.FilterInterceptor
	var part []factory.Component
	manager.SetFilter(filterer)
	manager.SetLogging(logger)

	if picker == 1 {
		part := manager.Execute("cpu", filter)
		//motherboard must be present
		if !cpuPresent && moboPresent {
			for i := range part {
				fmt.Print(part[i].PrintIDString())
			}
			//setting skip to false enable the getpartbyid method
			skip = false
			return part
		} else {
			fmt.Print("Cpu already added to build or motherboard has not been")
			//setting skip to true disables the getpartbyid method
			skip = true
		}

	} else if picker == 2 {
		//gpu
		part := manager.Execute("gpu", "")
		if !gpuPresent {
			for i := range part {
				fmt.Print(part[i].PrintIDString())
			}
			skip = false
			return part
		} else {
			fmt.Print("Gpu already added to build")
			skip = true
		}

	} else if picker == 3 {
		//psu
		part := manager.Execute("psu", "")
		if !psuPresent {
			for i := range part {
				fmt.Print(part[i].PrintIDString())
			}
			skip = false
			return part
		} else {
			fmt.Print("PSU already added to build")
			skip = true
		}

	} else if picker == 4 {
		//ram
		part := manager.Execute("ram", "")
		if ramPresent < ramSlots {
			for i := range part {
				fmt.Print(part[i].PrintIDString())
			}
			skip = false
			return part
		} else {
			if moboPresent {
				fmt.Println("Max ram cards have been reached.")
			} else {
				fmt.Println("Select your motherboard first")
			}
			skip = true
		}
	} else if picker == 5 {
		//mobo
		part := manager.Execute("mbd", "")
		if !moboPresent {
			for i := range part {
				fmt.Print(part[i].PrintIDString())
			}
			skip = false
			return part
		} else {
			fmt.Print("Motherboard already added to build")
			skip = true
		}

	} else if picker == 6 {
		//drive
		part := manager.Execute("drive", "")
		if drivesPresent < driveSlots {
			for i := range part {
				fmt.Print(part[i].PrintIDString())
			}
			skip = false
			return part
		} else {
			if moboPresent {
				fmt.Println("Max hard driveshave been reached.")
			} else {
				fmt.Println("Select your motherboard first")
			}
			skip = true
		}

	} else if picker == 7 {
		fmt.Println("Build Finished/Cancelled, Exiting")
		os.Exit(0)
	}
	return part

}

//get the selected item from the list printed
func getpartbyid(picker int, id int, partlist []factory.Component) factory.Component {
	if (id + 1) <= len(partlist) {
		var part factory.Component = partlist[id]
		if picker == 1 {
			//cpu
			part.Print()
			filter = part.GetFilter()
			//add cpu to motherboard and not tower
			motherboard.Add(part)
			cpuPresent = true
		} else if picker == 2 {
			//gpu
			part.Print()
			Tower.Add(part)
			gpuPresent = true
		} else if picker == 3 {
			//psu
			part.Print()
			Tower.Add(part)
			psuPresent = true
		} else if picker == 4 {
			//ram
			part.Print()
			motherboard.Add(part)
			ramPresent += 1
		} else if picker == 5 {
			//mobo
			part.Print()
			//get teh sovket of teh motherboard and set the filter
			filter = part.GetFilter()
			//set teh global moptherboard object to the selected motherboard
			motherboard = part
			moboPresent = true
			//set the max available slots for ram and drives
			ramSlots = part.GetRamSlots()
			driveSlots = part.GetDriveSlots()
		} else if picker == 6 {
			//drive
			part.Print()
			Tower.Add(part)
			drivesPresent += 1
			//invalid input
		} else if picker > 6 {
			part = nil
		}
		return part
	} else {
		return nil
	}
}

//funtion that loops to request the object categories and ids
func getUserInput() bool {
	var picker, id int
	var built bool
	var part []factory.Component
	fmt.Println("")
	fmt.Println("\nPC Parts -  1: CPU, 2: GPU, 3: PSU, 4: Ram, 5: Motherboard, 6: Drive; 7: Finish ")
	fmt.Scanln(&picker)
	if picker > 6 {
		return false
	} else {
		part = partPicker(picker, built)
		if part != nil && !skip {
			fmt.Println("\nPick One part!")
			fmt.Scanln(&id)
			getpartbyid(picker, id-1, part)
		}
	}
	return true
}

//Run client
func Run() {

	fmt.Println("Welcome to the PC Builder")
	//loops until the user selcts exit
	for {

		if !getUserInput() {
			break
		}

	}

	//add the global motherboard variable to the build
	Tower.Add(motherboard)
	fmt.Println(Tower.ToString())
	cloneTower := Tower.Clone()

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
